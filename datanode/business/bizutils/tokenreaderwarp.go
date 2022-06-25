// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

//  可以在持续读取数据的情况下保持token有效

package bizutils

import (
	"datanode/ifilestorage"
	"io"
	"time"
)

// NewTokenReaderWarp 可以在持续读取数据的情况下保持token有效
func NewTokenReaderWarp(t *ifilestorage.StreamToken, refreshms int64, r io.Reader, tr TokenRefreshI) io.Reader {
	return &tokenReaderWarp{maxCount: 1000, refreshMS: refreshms, token: t, r: r, tri: tr}
}

// tokenReaderWarp 可以在持续读取数据的情况下保持token有效
type tokenReaderWarp struct {
	maxCount     int
	currentCount int
	refreshMS    int64
	token        *ifilestorage.StreamToken
	tri          TokenRefreshI
	r            io.Reader
}

func (trw *tokenReaderWarp) Read(p []byte) (n int, err error) {
	if trw.token != nil {
		if trw.currentCount <= 0 {
			trw.currentCount = trw.maxCount
			if time.Now().UnixMilli()-trw.token.MTime >= trw.refreshMS {
				trw.token, _ = trw.tri.DoRefreshToken(trw.token.Token)
			}
		} else {
			trw.currentCount--
		}
	}
	return trw.r.Read(p)
}

type TokenRefreshI interface {
	DoRefreshToken(token string) (st *ifilestorage.StreamToken, err error)
}
