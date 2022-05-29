// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

//  随机上传示例

package example

import (
	"fmt"
	"io"
	"opensdk"
	"opensdk/utils"
	"os"
	"strings"
)

// UploadFile 上传文件
func UploadFile(dest, local string, override bool, o opensdk.IOpenApi) error {
	uploadPath := utils.Parse2UnixPath(dest + "/" + utils.GetPathName(local))
	if !override {
		if ok, err := o.IsExist(uploadPath); nil == err && ok {
			return fmt.Errorf("skip existing: %s", uploadPath)
		} else if nil != err {
			return err
		}
	}
	if token, err := o.DoAskWriteToken(uploadPath); nil != err {
		return err
	} else {
		if file, err := os.Open(local); nil == err {
			defer file.Close()
			if err = opensdk.TokenWriter(file, token, 128*1024*1024, o.DoWriteToken); nil != err {
				return err
			}
			if node, err := o.DoSubmitWriteToken(token.Token, override); nil != err {
				return err
			} else {
				fmt.Println("Uploaded successfully", uploadPath, node)
			}
		} else {
			return err
		}
	}
	return nil
}

// DownloadFile 下载文件
func DownloadFile(path, dest string, o opensdk.IOpenApi) error {
	if token, err := o.DoAskReadToken(path); nil != err {
		return err
	} else {
		if reader, err := o.DoReadToken(token.NodeNo, token.Token, token.EndPoint, 0); nil != err {
			return err
		} else {
			parentPath := utils.GetPathParent(utils.Parse2UnixPath(dest))
			if !utils.IsDir(parentPath) {
				if err := utils.MkdirAll(parentPath); nil != err {
					return err
				}
			}
			if file, err := utils.GetWriter(dest); nil != err {
				return err
			} else {
				if _, err := io.Copy(file, reader); nil != err {
					return err
				}
			}
		}
	}
	return nil
}

// uploadRandom 随机上传文件
func uploadRandom(destDir string, o opensdk.IOpenApi) string {
	random := utils.GetRandom(16)
	fmt.Println(random)
	uploadPath := utils.Parse2UnixPath(destDir + "/" + random[:3] + "/" + random[3:6] + "/" + random[6:9] + "/" + random[9:])
	token, err := o.DoAskWriteToken(uploadPath)
	checkError(err)
	err = streamUploader(random, token, o.DoWriteToken)
	checkError(err)
	node, err := o.DoSubmitWriteToken(token.Token, true)
	checkError(err)
	// err = o.DoDelete(uploadPath)
	// checkError(err)
	fmt.Println("Uploaded successfully", uploadPath, node)
	return uploadPath
}

type doWriteTokenFunc func(nodeNo, token, endpoint string, pieceNumber int, sha256 string, reader io.Reader) (err error)

// streamUploader streamUploader
func streamUploader(random string, token *opensdk.StreamToken, doWriteToken doWriteTokenFunc) (err error) {
	var sha256 string
	reader := strings.NewReader(random)
	// if sha256, err = utils.GetFileSHA256(reader); nil == err {
	// reader.Seek(0, io.SeekStart)
	err = doWriteToken(token.NodeNo, token.Token, token.EndPoint, 1, sha256, reader)
	// }
	return err
}

func checkError(err error) {
	if nil != err {
		panic(err)
	}
}
