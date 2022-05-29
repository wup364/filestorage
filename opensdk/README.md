# openSDK 文件存储开放SDK

    实现了对文件存储服务的开放接口封装, 可实现文件的存取和管理.

## 基本信息

### 服务鉴权

| 所属程序 |  账号密码 |  密码  |  描述  |
| ------- | ------|------ | ------ |
| namenode | OPENAPI | `空` | 可以调用开放的api接口, opensdk使用 |

### 方法列表

|  Description |    Method    |
| ------------ | ------------ |
| IsDir | `IsDir(src string) bool` |
| IsFile | `IsFile(src string) bool` |
| IsExist | `IsExist(src string) bool` |
| GetFileSize | `GetFileSize(src string) int64` |
| | |
| GetNode | `GetNode(src string) *TNode` |
| GetDirNameList | `GetDirNameList(src string, limit, offset int) (*DirNameListDto, int)` |
| GetDirNodeList | `GetDirNodeList(src string, limit, offset int) (*DirNodeListDto, int)` |
| | |
| DoMkDir | `DoMkDir(path string) error` |
| DoDelete | `DoDelete(src string) error` |
| DoRename | `DoRename(src string, dst string) error` |
| DoCopy | `DoCopy(src, dst string) error` |
| DoMove | `DoMove(src, dst string) error` |
| | |
| DoQueryToken | `DoQueryToken(token string) (*StreamToken, error)` |
| DoAskReadToken | `DoAskReadToken(src string) (*StreamToken, error)` |
| DoAskWriteToken | `DoAskWriteToken(src string) (*StreamToken, error)` |
| DoRefreshToken | `DoRefreshToken(token string) (*StreamToken, error)` |
| DoSubmitWriteToken | `DoSubmitWriteToken(token string, override bool) (node *TNode, err error)` |
| SetNodeStreamAddr | `SetNodeStreamAddr(nodes map[string]string)` |
| DoWriteToken | `DoWriteToken(nodeNo, token string, pieceNumber int, sha256 string, reader io.Reader) (err error)` |
    sdk暂时仅支持golang语言. 可以单独拷贝到其他项目使用. `example`目录下包含一些调用示例供参考使用.
