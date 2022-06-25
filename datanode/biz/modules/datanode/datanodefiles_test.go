// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 数据存储节点-文件归档
package datanode

import (
	"datanode/biz/bizutils"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/wup364/pakku/utils/utypes"
)

func TestGetArchivedPath4Hash(t *testing.T) {
	// e/0d/afb6109ade/198327e54c/04b9e92ba9/25f29292f3/16210f4a98/e0dafb6109ade198327e54c04b9e92ba925f29292f316210f4a988c0851ea9b8
	fmt.Println(getArchivedPath4Hash("e0dafb6109ade198327e54c04b9e92ba925f29292f316210f4a988c0851ea9b8"))
}
func TestMapLoop(t *testing.T) {
	tk := (&bizutils.TokenManager{}).Init()
	timeStart := time.Now()
	for i := 0; i < 10000; i++ {
		tb := utypes.NewSafeMap()
		for j := 0; j < 100; j++ {
			tb.Put(i, j)
		}
		tk.AskToken(tb, 3000)
	}
	fmt.Println((time.Now().Nanosecond() - timeStart.Nanosecond()) / int(time.Microsecond))
	timeStart = time.Now()
	keys := tk.ListTokens()
	for i := 0; i < len(keys); i++ {
		if val, ok := tk.GetTokenBody(keys[i]); ok {
			err := (val.(*utypes.SafeMap)).DoRange(func(key, val interface{}) error {
				if key.(int)%4 == 0 {
					return errors.New("break")
				}
				return nil
			})
			if nil != err && err.Error() == "break" {
				// fmt.Println(i, val)
				continue
			}
		}
	}
	fmt.Println((time.Now().Nanosecond() - timeStart.Nanosecond()) / int(time.Microsecond))
}
