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
	"encoding/json"
	"errors"
)

// ErrInvalidToken 无效的token
var ErrInvalidToken = errors.New("invalid token")

// PropJSON4DataNode PropJSON4DataNode
type PropJSON4DataNode struct {
	NodeNo      string
	NodePid     string
	RPCAddress  string
	HTTPAddress string
}

// ToJSON ToJSON
func (prop PropJSON4DataNode) ToJSON() string {
	if bt, err := json.Marshal(prop); nil == err {
		return string(bt)
	}
	return ""
}

// ParseJSON ParseJSON
func (prop PropJSON4DataNode) ParseJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), &prop)
}
