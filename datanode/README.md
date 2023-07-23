# datanode 文件存取服务

    文件的归档和读取, 提供文件的上传和下载HTTP服务, 启动datanode前需要先启动连接的namenode. 目前不支持多节点主从备份、故障自恢复等功能.

## 程序基本信息

### 监听端口

| 程序     | 描述         | 端口                 | 配置                           |
| -------- | ------------ | -------------------- | ------------------------------ |
| datanode | 内部服务接口 | 127.0.0.1:5061(rpc)  | `listen.rpc.address`  |
| datanode | 上传下载接口 | 127.0.0.1:5062(http) | `listen.http.address` |

### 服务鉴权

| 所属程序 | 账号密码            | 密码   | 描述                         |
| -------- | ------------------- | ------ | ---------------------------- |
| datanode | 节点名称,默认 DN101 | `随机` | 仅供内部调用, 每次启动时生成 |


### 配置清单

| 所属程序 | KEY                         | 默认值                           | 可选值                     | 描述                                       |
| -------- | --------------------------- | -------------------------------- | -------------------------- | ------------------------------------------ |
| datanode | `nodeno`                    | DN101                            | `在同一套系统中名字不重复` | 用于在 namenode 中注册节点                 |
| datanode | `namenode.rpc.address`      | 127.0.0.1:5051                   | `namenode服务地址`         | 用于 datanode 连接 namenode                |
| datanode | `namenode.rpc.user`         | `DATANODE`                       | `namenode上存在的DATANODE角色账户`| 用于 datanode 连接 namenode          |
| datanode | `namenode.rpc.pwd`          | `空`                             | `密码`                     | 用于 datanode 连接 namenode                 |
| datanode | `listen.rpc.address`        | 127.0.0.1:5061                   | `*`                        | RPC 服务监听地址                           |
| datanode | `listen.rpc.endpoint`       | `{listen.http.address}`          | `*`                        | RPC 服务访问地址, 可设为仅 namenode 可访问 |
| datanode | `listen.http.address`       | 127.0.0.1:5062                   | `*`                        | HTTP 服务监听地址                          |
| datanode | `listen.http.endpoint`      | `http://{listen.http.address}`   | `*`                        | HTTP 服务访问地址, 带协议头                |
| datanode | `store.datanode.datasource` | ${appname}#datanode?cache=shared | `sqlite3支持的地址`         | 文件对应的块信息                       |
| datanode | `store.datahash.datasource` | ${appname}#datahash?cache=shared | `sqlite3支持的地址`         | 文件块hash值信息索引                         |
| datanode |`filedatas.encryption.passwd`| `空`                             | `*`                        | 文件加密密钥, 设置后不能更改, 会造成更改之前的文件无法正常解密 |
| datanode | `clearsetting.deletefile`   | false                            | `false\|true`              | 标记在清理已删除文件时, 是否物理删除磁盘文件      |

    . 配置使用json格式存储, 格式示例:
        `{
            "listen": {
                "rpc": {
                    "address": "127.0.0.1:5051"
                }
            }
        }` => listen.rpc.address

### HTTP 接口列表

| Description | URL (`{listen.http.endpoint}/*`)                    |
| ----------- | --------------------------------------------------- |
| 文件上传    | `/stream/put/{token}?number={number}&hash={sha256}` |
| 文件下载    | `/stream/read/{token}?name={name}`                  |
