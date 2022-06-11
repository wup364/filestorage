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

    用户信息数据库配置 `store.user4rpc.driver=sqlite3` 和 `store.user4rpc.datasource=${appname}#user?cache=shared`, 可以更改成mysql的连接;
    datanode的节点名称可以动过配置改变 `nodeno=DN101`;

### 配置清单

| 所属程序 | KEY                         | 默认值                           | 可选值                     | 描述                                       |
| -------- | --------------------------- | -------------------------------- | -------------------------- | ------------------------------------------ |
| datanode | `nodeno`                    | DN101                            | `在同一套系统中名字不重复` | 用于在 namenode 中注册节点                 |
| datanode | `listen.rpc.address`        | 127.0.0.1:5061                   | `*`                        | RPC 服务监听地址                           |
| datanode | `listen.rpc.endpoint`       | `{listen.http.address}`          | `*`                        | RPC 服务访问地址, 可设为仅 namenode 可访问 |
| datanode | `listen.http.address`       | 127.0.0.1:5062                   | `*`                        | HTTP 服务监听地址                          |
| datanode | `listen.http.endpoint`      | `http://{listen.http.address}`   | `*`                        | HTTP 服务访问地址, 带协议头                |
| datanode | `store.datanode.driver`     | sqlite3                          | `sqlite3\|mysql`           | 用于记录文件分片信息                       |
| datanode | `store.datanode.datasource` | ${appname}#datanode?cache=shared | `sqlite3\|mysql支持的地址` | 用于记录文件分片信息                       |
| datanode | `store.datahash.driver`     | sqlite3                          | `sqlite3\|mysql`           | 文件 hash 信息索引                         |
| datanode | `store.datahash.datasource` | ${appname}#datahash?cache=shared | `sqlite3\|mysql支持的地址` | 文件 hash 信息索引                         |
| datanode | `clearsetting.deletefile`   | false                            | `false\|true`              | 标记是否清理已删除文件                     |
| datanode | `namenode.rpc.address`      | 127.0.0.1:5051                   | `namenode服务地址`         | 用于 datanode 连接 namenode                |
| datanode | `namenode.rpc.pwd`          | `空`                             | `密码`                     | 标记是否清理已删除文件                     |

    . 配置使用json格式存储, 格式示例:
        `{
            "listen": {
                "rpc": {
                    "address": "127.0.0.1:5051"
                }
            }
        }` => listen.rpc.address
    . 数据库默认使用sqlite, 但支持MySQL数据库, 连接示例: `user:pwd@tcp(192.168.2.8:3306)/filestorage`

### HTTP 接口列表

| Description | URL (`{listen.http.endpoint}/*`)                    |
| ----------- | --------------------------------------------------- |
| 文件上传    | `/stream/put/{token}?number={number}&hash={sha256}` |
| 文件下载    | `/stream/read/{token}?name={name}`                  |
