# fstools 工具集合

    使用`opensdk`编写的小工具.

## 程序基本信息

### 程序清单

|  程序名称 |  编译位置  |  描述 |
| ------ | ------ | ------ |
| shelltool | `${project}/shelltool` | 可连接存储进行文件的基本管理和服务管理, 模拟linux终端操作, 如: 上传(put)下载(get)文件、 查看目录(ls)信息等 |
| uploadtool | `${project}/uploadtool` | 单独的上传小工具, 可用于上传文件夹或随机上传测试等 |

### 配置清单

|  所属程序 |  KEY  |  默认值  |  可选值  |  描述 |
| ------ | ------ | ------ | ---- | ---- |
| shelltool | `datanodes` | `{"DN101": "http://127.0.0.1:5062"}` | `Key-Value键值对` | datanode地址信息 |
| uploadtool | `rpc.address` | 127.0.0.1:5051 | `*` | RPC服务地址 |
| uploadtool | `auth.user` | OPENAPI | `*` | RPC鉴权用户 |
| uploadtool | `auth.pwd` | `空` | `*` | RPC鉴权用户 |
| uploadtool | `datanodes` | `{"DN101": "http://127.0.0.1:5062"}` | `Key-Value键值对` | datanode地址信息 |

    . 配置使用json格式存储, 格式示例: 
        `{
            "rpc": {
                "address": "127.0.0.1:5051"
            }
        }` => rpc.address

### 命令行参数

|  所属程序 |  KEY  |  默认值   |  描述 |
| ------ | ------ | ------ | ---- |
| shelltool | `第一个参数` | `空` | 连接url`账号@IP:端口(默认5051)`, 如: `DATANODE@192.168.2.201` |
| uploadtool | `scandir` | `null` | The local folder that needs to be scanned, default null  |
| uploadtool | `destdir` | `/` | The destination folder to upload to, default /  |
| uploadtool | `override` | `false` | Whether to overwrite existing files, default false  |
| uploadtool | `random` | `false` | Whether to overwrite existing files, default false  |
| uploadtool | `logger` | `console` | logger: console, file or unset, default console |
| uploadtool | `loglevel` | `debug` | oglevel: debug, info, error or none, default debug |
| uploadtool | `logdir` | `./logs/{application name}.log` | default ./logs/{application name}.log |
