# namenode 文件元数据服务

    对文件的新增、删除、修改、移动、列表等基础操作RPC服务. 记录文件树结构信息以及管理datanode信息, 加载时会读取全部数据到内存, 因此需要注意内存大小的分配. 目前不支持多节点主从备份、故障自恢复等功能.

## 程序基本信息

### 监听端口

|  程序 |  描述  |  端口  |  配置  |
| ----- |------ | ------ | ---- |
| namenode | 记录文件数据信息的服务 | 127.0.0.1:5051(rpc) | `listen.rpc.address` |

### 服务鉴权

| 所属程序 |  账号密码 |  密码  |  角色  |  描述  |
| ------- | ------|------ | ------ | ------ |
| namenode | NAMENODE | `空` | NAMENODE | 管理账号, 可以操作账户信息, 但不可以操作文件类接口 |
| namenode | DATANODE | `空` | DATANODE | datanode连接时使用, 可以调用内部接口 |
| namenode | OPENAPI | `空` | OPENAPI |  调用开放api接口时使用, 可以调用文件部分接口 |

    1) 初次使用时, 在namenode启动成功后, 可使用`shelltool`连接服务器进行管理.
    2) 用户信息数据库配置 `store.user4rpc.driver=sqlite3` 和 `store.user4rpc.datasource=${appname}#user?cache=shared`, 可以更改成mysql的连接.

### 配置清单

|  所属程序 |  KEY  |  默认值  |  可选值  |  描述 |
| ------ | ------ | ------ | ---- | ---- |
| namenode | `listen.rpc.address` | 127.0.0.1:5051 | `*` | RPC服务监听地址 |
| namenode | `store.deletedaddr.driver` | sqlite3 | `sqlite3\|mysql` | 已删除的文件地址信息 |
| namenode | `store.deletedaddr.datasource` | ${appname}#deleted?cache=shared | `sqlite3\|mysql支持的地址` | 已删除的文件地址信息 |
| namenode | `store.filenames.driver` | sqlite3 | `sqlite3\|mysql` | 完整的文件树信息 |
| namenode | `store.filenames.datasource` | ${appname}#filenames?cache=shared | `sqlite3\|mysql支持的地址` | 完整的文件树信息 |

    . 配置使用json格式存储, 格式示例: 
        `{
            "listen": {
                "rpc": {
                    "address": "127.0.0.1:5051"
                }
            }
        }` => listen.rpc.address
    . 数据库默认使用sqlite, 但支持MySQL数据库, 连接示例: `user:pwd@tcp(192.168.2.8:3306)/filestorage`

### 参考内存配置

|  文件总量 |  文件夹  |  文件  |  平均长度(文件夹\|文件)  |  占用内存 |
| ------ | ------ | ------ | ---- | ---- |
| 1000046 | 46 | 1000000 |  8 \| 32 | 545.7M |
| 5000049 | 49 | 5000000 |  8 \| 32 | 3174.2M |

### RPC接口列表

#### 文件接口

|  Description |    Method    |
| ------------ | ------------ |
| IsDir | `NameNode.IsDir(src string) bool` |
| IsFile | `NameNode.IsFile(src string) bool` |
| IsExist | `NameNode.IsExist(src string) bool` |
| GetFileSize | `NameNode.GetFileSize(src string) int64` |
| | |
| GetNode | `NameNode.GetNode(src string) *TNode` |
| GetDirNameList | `NameNode.GetDirNameList(src string, limit, offset int) (*DirNameListDto, int)` |
| GetDirNodeList | `NameNode.GetDirNodeList(src string, limit, offset int) (*DirNodeListDto, int)` |
| | |
| DoMkDir | `NameNode.DoMkDir(path string) error` |
| DoDelete | `NameNode.DoDelete(src string) error` |
| DoRename | `NameNode.DoRename(src string, dst string) error` |
| DoCopy | `NameNode.DoCopy(src, dst string) error` |
| DoMove | `NameNode.DoMove(src, dst string) error` |
| | |
| DoQueryToken | `NameNode.DoQueryToken(token string) (*StreamToken, error)` |
| DoAskReadToken | `NameNode.DoAskReadToken(src string) (*StreamToken, error)` |
| DoAskWriteToken | `NameNode.DoAskWriteToken(src string) (*StreamToken, error)` |
| DoRefreshToken | `NameNode.DoRefreshToken(token string) (*StreamToken, error)` |
| DoSubmitWriteToken | `NameNode.DoSubmitWriteToken(token string, override bool) (node *TNode, err error)` |
