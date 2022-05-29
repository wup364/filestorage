// Copyright (C) 2020 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 虚拟路径挂载

package ifiledatas

// MountNode 挂载节点信息
type MountNode struct {
	Depth  int               // 深度
	Path   string            // 挂载路径-虚拟路径
	Type   string            // 挂载类型
	Addr   string            // 实际挂载路径
	Passwd string            // 部分连接可能有密码
	Props  map[string]string // 其他部分拓展数据
	Driver FileDriver        // 驱动实例
}

// DIRMount 虚拟路径挂载管理
type DIRMount interface {

	// RegisterFileDriver 注册文件驱动
	RegisterFileDriver(drivers ...DIRMountRegister) DIRMount

	// LoadAllMount 挂载所有挂载接节点
	LoadAllMount(mounts map[string]interface{}) DIRMount

	// GetFileDriver 根据路径获取对应驱动类
	GetFileDriver(relativePath string) FileDriver

	// GetMountNode 查找相对路径下的分区挂载信息
	GetMountNode(relativePath string) *MountNode

	// ListChildMount 查找符合当前路径下的子挂载分区路径 /==>/Mount
	ListChildMount(relativePath string) (res []string)
}

// DIRMountRegister 路径挂载注册器
type DIRMountRegister interface {
	// GetDriverType GetDriverType
	GetDriverType() string
	// InstanceDriver 当驱动实例化时调用
	InstanceDriver(dirMount DIRMount, mtnode *MountNode) (FileDriver, error)
}
