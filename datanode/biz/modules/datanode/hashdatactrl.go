// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package datanode

import (
	"crypto/sha256"
	"datanode/biz/bizutils"
	"datanode/ifilestorage"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"pakku/utils/utypes"
	"sort"
	"strconv"
	"sync"
)

// NewHashDataCtrl  hash文件写入、读取、删除控制
func NewHashDataCtrl() *HashDataCtrl {
	return &HashDataCtrl{
		delhashlocker: new(sync.RWMutex),
		dellock:       utypes.NewSafeMap(),
		wtokens:       (&bizutils.TokenManager{}).Init(),
	}
}

// HashDataCtrl hash文件写入、读取、删除控制
type HashDataCtrl struct {
	delhashlocker *sync.RWMutex
	dellock       *utypes.SafeMap
	wtokens       *bizutils.TokenManager
}

// GetWriteReader 获取写入用的reader
func (hdc *HashDataCtrl) GetWriteReader(opts *ifilestorage.StreamWriteOpts, reader io.Reader) (*WriteReader, error) {
	wr := &WriteReader{
		reader: reader,
		wopts:  opts,
		hash:   sha256.New(),
	}
	if err := hdc.putWriteReader(opts, wr); nil != err {
		return nil, err
	}
	return wr, nil
}

// GetPiecesAndSucceed 获取写的pieces信息, 前提是全部写入成功
func (hdc *HashDataCtrl) GetPiecesAndSucceed(token string) (wopts []ifilestorage.StreamWriteOpts, err error) {
	if val, ok := hdc.wtokens.GetTokenBody(token); ok {
		hdc.wtokens.RefreshToken(token)
		if err = val.(*utypes.SafeMap).DoRange(func(key, val interface{}) error {
			wr := val.(*WriteReader)
			if !wr.IsClosed() {
				return errors.New("the file is being written, the operation failed, token: " + strconv.Itoa(key.(int)))
			}
			// 校验hash
			if wr.wopts.Sha256 != wr.GetHash() {
				return errors.New("hash verification failed, token: " + strconv.Itoa(key.(int)))
			}
			wopts = append(wopts, *wr.wopts)
			return nil
		}); nil == err {
			sort.Sort(&StreamWriteOptsSort{
				Asc:  true,
				arry: wopts,
			})
		}
	} else {
		err = ErrInvalidToken
	}
	return wopts, err
}

// GetReader 读取数据用的reader
func (hdc *HashDataCtrl) GetReader(opts ifilestorage.StreamReadOpts) (*Reader, error) {
	r := &Reader{
		ropts:  opts,
		pieces: opts.Pieces,
		lock:   new(sync.RWMutex),
	}
	return r, nil
}

// CanQuoteHash 是否可以使用这个hash连接, 要求不在删除队列中
func (hdc *HashDataCtrl) CanQuoteHash(hash string) bool {
	hdc.delhashlocker.RLock()
	defer hdc.delhashlocker.RUnlock()
	return !hdc.dellock.ContainsKey(hash)
}

// LockHashOnArchive 等待可以归档hash文件的时机, 如果在删除队列中, 则需要等待
func (hdc *HashDataCtrl) LockHashOnArchive(hash string) {
	hdc.delhashlocker.RLock()
	if val, ok := hdc.dellock.Get(hash); ok {
		hdc.delhashlocker.RUnlock()
		val.(*sync.Mutex).Lock()
		if hdc.dellock.ContainsKey(hash) {
			// 如果获得锁后, hash依然存在, 则说明在删除执行前获得了锁, 则可以先取消这次删除
			hdc.dellock.Delete(hash)
			val.(*sync.Mutex).Unlock()
		}
	} else {
		hdc.delhashlocker.RUnlock()
	}
}

// LockHashOnDelete 锁定这个即将删除的hash, 如果锁定成功,则在同hash文件归档时需要等待解锁, 反之这个hash不能删除
func (hdc *HashDataCtrl) LockHashOnDelete(hash string) bool {
	hdc.delhashlocker.RLock()
	if val, ok := hdc.dellock.Get(hash); ok {
		hdc.delhashlocker.RUnlock()
		val.(*sync.Mutex).Lock()
		return hdc.dellock.ContainsKey(hash)
	}
	return false
}

// ReleaseDeleteLock 释放hash删除锁
func (hdc *HashDataCtrl) ReleaseDeleteLock(hash string) {
	hdc.delhashlocker.Lock()
	defer hdc.delhashlocker.Unlock()
	if val, ok := hdc.dellock.Get(hash); ok {
		hdc.dellock.Delete(hash)
		val.(*sync.Mutex).Unlock()
	}
}

//RemoveHashDeleteMark 删除hash删除标记(取消删除)
func (hdc *HashDataCtrl) RemoveHashDeleteMark(hash string) {
	hdc.delhashlocker.Lock()
	defer hdc.delhashlocker.Unlock()
	hdc.dellock.Delete(hash)
}

// MarkHashOnDelete 标记这个hash正在被删除中, 返回标记否成功的(正在使用中的无法锁定)
func (hdc *HashDataCtrl) MarkHashOnDelete(hashs []string) []string {
	// 这个执行时间内不允许新建
	hdc.delhashlocker.Lock()
	defer hdc.delhashlocker.Unlock()
	hashsmap := make(map[string]uint8)
	for i := 0; i < len(hashs); i++ {
		if _, ok := hashsmap[hashs[i]]; !ok {
			hashsmap[hashs[i]] = 0
		}
	}
	if tokens := hdc.wtokens.ListTokens(); len(tokens) > 0 {
		for i := 0; i < len(tokens); i++ {
			// 获取一个token的所有信息
			if val, ok := hdc.wtokens.GetTokenBody(tokens[i]); ok {
				// 如果token的某一片里有和待锁定相同的hash, 说明这些hash正在使用
				(val.(*utypes.SafeMap)).DoRange(func(key, val interface{}) error {
					wr := val.(*WriteReader)
					if len(wr.wopts.Sha256) > 0 {
						delete(hashsmap, wr.wopts.Sha256)
					}
					return nil
				})
			}
		}
	}
	res := make([]string, 0)
	if len(hashsmap) > 0 {
		for key := range hashsmap {
			if nil == hdc.dellock.PutX(key, new(sync.Mutex)) {
				res = append(res, key)
			}
		}
	}
	return res
}

// putWriteReader 记录每个token的每一片的写入情况
func (hdc *HashDataCtrl) putWriteReader(opts *ifilestorage.StreamWriteOpts, reader *WriteReader) error {
	hdc.delhashlocker.RLock()
	defer hdc.delhashlocker.RUnlock()
	if val, ok := hdc.wtokens.GetTokenBody(opts.Token); ok {
		hdc.wtokens.RefreshToken(opts.Token)
		m := val.(*utypes.SafeMap)
		if rval, ok := m.Get(opts.PieceNumber); ok {
			if !rval.(*WriteReader).IsClosed() {
				return errors.New("cannot write to the same file at the same time")
			}
		}
		m.Put(opts.PieceNumber, reader)
	} else {
		m := utypes.NewSafeMap()
		m.Put(opts.PieceNumber, reader)
		hdc.wtokens.PutTokenBody(opts.Token, m, tokenexp)
	}
	return nil
}

// Reader Reader
type Reader struct {
	reader   io.Reader
	piecenum int
	pieces   []string
	readed   int64
	lock     *sync.RWMutex
	ropts    ifilestorage.StreamReadOpts
	doGetPR  func(hash string, offset int64) (io.Reader, error)
	doGetPS  func(hash string) (int64, error)
}

// OnPiecesReader 数据流获取时调用
func (r *Reader) OnPiecesReader(fuc func(hash string, offset int64) (io.Reader, error)) {
	r.doGetPR = fuc
}

// OnPiecesSize 数据大小获取时调用
func (r *Reader) OnPiecesSize(fuc func(hash string) (int64, error)) {
	r.doGetPS = fuc
}

// getNextPieces 获取下一片reader
func (r *Reader) getNextPieces() error {
	if len(r.pieces) <= r.piecenum+1 {
		return io.EOF
	}
	if val, err := r.doGetPR(r.pieces[r.piecenum+1], 0); nil == err {
		r.piecenum++
		r.reader = val
	} else {
		return err
	}
	return nil
}

// SetOffset 计算偏移
func (r *Reader) SetOffset(offset int64) (err error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	// 往已经读过的位置读取
	if offset < r.readed {
		return errors.New("unsupported, offset < readed size")
	} else if offset > r.readed {
		for {
			// 往没有读过的位置读取
			var size int64
			if size, err = r.doGetPS(r.pieces[r.piecenum]); nil != err {
				return err
			} else {
				// 当前这一片够用
				if r.readed+size > offset {
					poffset := offset - r.readed
					// 读取这一片的reader
					if val, err := r.doGetPR(r.pieces[r.piecenum], poffset); nil == err {
						r.reader = val
					}
					break
				}
				// 尝试下一片
				r.readed += size
				r.piecenum++
			}
		}
	}
	return err
}

// Read Read
func (r *Reader) Read(p []byte) (n int, err error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	if nil == r.reader {
		if r.reader, err = r.doGetPR(r.pieces[r.piecenum], 0); nil != err {
			return 0, err
		}
	}
	n, err = r.reader.Read(p)
	if err == io.ErrUnexpectedEOF || err == io.EOF {
		if n < 1 {
			// 一片数据读取完成
			var size int64
			if size, err = r.doGetPS(r.pieces[r.piecenum]); nil == err {
				r.readed += size
				// 下一片数据读取开始
				if err = r.getNextPieces(); nil == err {
					n, err = r.reader.Read(p)
				}
			}
		} /*else {
			err = nil
		}*/
	}
	return n, err
}

// WriteReader WriteReader
type WriteReader struct {
	wopts   *ifilestorage.StreamWriteOpts
	hash    hash.Hash
	reader  io.Reader
	hashStr string
	closed  bool
}

// Read Read & 计算hash
func (wrw *WriteReader) Read(p []byte) (n int, err error) {
	n, err = wrw.reader.Read(p)
	if err == nil || err == io.ErrUnexpectedEOF || err == io.EOF {
		if n > 0 {
			newbyte := p[0:n]
			if _, err := wrw.hash.Write(newbyte); err != nil {
				return -1, err
			}
		}
	}
	return n, err
}

// GetHash 获取片段Hash
func (wrw *WriteReader) GetHash() string {
	if len(wrw.hashStr) > 0 {
		return wrw.hashStr
	}
	return hex.EncodeToString(wrw.hash.Sum(nil))
}

// SetHash 设置片段Hash[]
func (wrw *WriteReader) SetHash(hash string) {
	wrw.hashStr = hash
}

// IsQuote 是否是引用hash, 没有传递实际文件
func (wrw *WriteReader) IsQuote() bool {
	return len(wrw.hashStr) > 0
}

// Close 是否关闭
func (wrw *WriteReader) IsClosed() bool {
	return wrw.closed
}

// Close 关闭reader
func (wrw *WriteReader) Close() {
	wrw.closed = true
	wrw.reader = nil
	wrw.SetHash(wrw.GetHash())
}

// StreamWriteOptsSort 文件排序
type StreamWriteOptsSort struct {
	arry []ifilestorage.StreamWriteOpts
	Asc  bool
}

// 实现sort.Interface接口取元素数量方法
func (sort *StreamWriteOptsSort) Len() int {
	return len(sort.arry)
}

// 实现sort.Interface接口比较元素方法
func (sort *StreamWriteOptsSort) Less(i, j int) bool {
	less := sort.arry[i].PieceNumber < sort.arry[j].PieceNumber
	return less
}

// 实现sort.Interface接口交换元素方法
func (sort *StreamWriteOptsSort) Swap(i, j int) {
	sort.arry[i], sort.arry[j] = sort.arry[j], sort.arry[i]
}
