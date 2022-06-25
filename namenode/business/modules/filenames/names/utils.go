// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

//

package names

import (
	"errors"
	"namenode/ifilestorage"
)

// CheckNodeNotNull 节点必填校验
func CheckNodeNotNull(node ifilestorage.TNode, isRoot bool) error {
	if len(node.Id) == 0 {
		return errors.New("node.Id is empty")
	}
	if !isRoot && len(node.Pid) == 0 {
		return errors.New("node.Pid is empty")
	}
	if len(node.Name) == 0 {
		return errors.New("node.Name is empty, id: " + node.Id)
	}
	if node.Flag == FLAG_FILE && len(node.Addr) == 0 {
		return errors.New("node.Addr is empty, id: " + node.Id)
	}
	return nil
}
