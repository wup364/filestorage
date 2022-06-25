// Copyright (C) 2020 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 挂载管理器, 文件虚拟目录解析&路径驱动获取

package dirmount

import (
	"datanode/business/modules/filedatas/ifiledatas"
	"errors"
	"strings"

	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
)

// 常量
const (
	keyMountType   string = "type"   // 配置文件-挂载类别
	keymountAddr   string = "addr"   // 配置文件-挂载地址
	keymountpasswd string = "passwd" // 配置文件-挂载密码
)

// DIRMount 挂载管理器
type DIRMount struct {
	nodes   []*ifiledatas.MountNode
	drivers map[string]ifiledatas.DIRMountRegister
}

// RegisterFileDriver 注册文件驱动
func (mount *DIRMount) RegisterFileDriver(drivers ...ifiledatas.DIRMountRegister) ifiledatas.DIRMount {
	if nil == mount.drivers {
		mount.drivers = make(map[string]ifiledatas.DIRMountRegister)
	}
	if nil != drivers && len(drivers) > 0 {
		for i := 0; i < len(drivers); i++ {
			mount.drivers[drivers[i].GetDriverType()] = drivers[i]
		}
	}
	return mount
}

// LoadAllMount 初始化挂载节点 mnodes: {'/':{type:'local', addr:'./datas'}}
func (mount *DIRMount) LoadAllMount(mnodes map[string]interface{}) ifiledatas.DIRMount {
	if len(mnodes) == 0 {
		logs.Panicln("mnodes is nil")
	}
	if nil == mount.drivers || len(mount.drivers) == 0 {
		logs.Panicln("drivers is nil")
	}
	nodes := make([]*ifiledatas.MountNode, len(mnodes))
	count := 0
	for key, val := range mnodes {
		confVal := val.(map[string]interface{})
		mtnode := ifiledatas.MountNode{
			Path:  key,
			Type:  confVal[keyMountType].(string),
			Addr:  confVal[keymountAddr].(string),
			Depth: 0,
		}
		if val, ok := confVal[keymountpasswd]; ok {
			mtnode.Passwd = val.(string)
		}
		nodes[count] = mount.parseMountNode(&mtnode)
		count++
	}
	mount.nodes = nodes
	return mount
}

// GetFileDriver 根据相对路径获取对应驱动类
func (mount *DIRMount) GetFileDriver(relativePath string) ifiledatas.FileDriver {
	if len(strings.Replace(relativePath, " ", "", -1)) == 0 {
		relativePath = "/"
	}
	// 挂载节点
	mtnode := mount.GetMountNode(relativePath)
	// 解析 mtnode
	if nil == mtnode || mtnode.Path == "" {
		logs.Panicln(errors.New("mount path is not find"))
	}
	if mtnode.Addr == "" {
		logs.Panicln(errors.New("mount address is nil, at mount path: " + mtnode.Path))
	}
	if mtnode.Type == "" {
		logs.Panicln(errors.New("mount Type is not find"))
	}
	// 不支持的分区挂载类型
	if mtnode.Driver == nil {
		logs.Panicln(errors.New("Unsupported partition mount type: " + mtnode.Type))
	}
	return mtnode.Driver
}

// GetMountNode 查找相对路径下的分区挂载信息
func (mount *DIRMount) GetMountNode(relativePath string) *ifiledatas.MountNode {
	var rootMount *ifiledatas.MountNode
	for i := 0; i < len(mount.nodes); i++ {
		// 相等|startWith情况
		if mount.nodes[i].Path == relativePath || strings.HasPrefix(relativePath, mount.nodes[i].Path+"/") {
			return mount.nodes[i]
		}
		// 根路径情况
		if mount.nodes[i].Path == "/" {
			rootMount = mount.nodes[i]
		}
	}
	return rootMount
}

// ListChildMount 查找符合当前路径下的子挂载分区路径 /==>/Mount
func (mount *DIRMount) ListChildMount(relativePath string) (res []string) {
	if relativePath != "/" {
		return res
	}
	depth := len(strings.Split(relativePath, "/")) // 这个地方实质上+1了
	for i := 0; i < len(mount.nodes); i++ {
		if relativePath == "/" {
			// 如果为 / 则取挂载目录深度为 1 的 /==>/mount1 /mount2
			if mount.nodes[i].Depth == 1 && mount.nodes[i].Path != "/" {
				res = append(res, mount.nodes[i].Path)
			}
		} else
		// 其他目录则取当前目录深度加一目录&以他开头的 /ps==>/ps/mount1 /ps/mount2
		if mount.nodes[i].Depth == depth && mount.nodes[i].Path != "/" &&
			strings.HasPrefix(mount.nodes[i].Path, relativePath+"/") {
			res = append(res, mount.nodes[i].Path)
		}
	}
	return res
}

// parseMountNode 转换配置信息, 如: 相对路径转绝对路径
func (mount *DIRMount) parseMountNode(mtnode *ifiledatas.MountNode) *ifiledatas.MountNode {
	// 统一挂载类型大写
	mtnode.Type = strings.ToUpper(mtnode.Type)
	if driver, ok := mount.drivers[mtnode.Type]; ok {
		if instance, err := driver.InstanceDriver(mount, mtnode); nil == err {
			mtnode.Driver = instance
			mtnode.Path = strutil.Parse2UnixPath(mtnode.Path)
			mtnode.Depth = len(strings.Split(mtnode.Path, "/")) - 1
		} else {
			logs.Panicln(err, mtnode)
		}
	} else {
		logs.Panicln(errors.New("Unsupported partition mount type: " + mtnode.Type))
	}
	logs.Infoln("> Mounting partition: ", mtnode)
	return mtnode
}
