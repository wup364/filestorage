# FileStorage 文件存储服务

    这是一个存储服务, 使用namenode存储元数据信息, datanode负责数据的存储和读取. 每个组件的详细说明在对应的目录下

## 目录结构

| 组件          | 描述                                                                                   |
| ------------- | -------------------------------------------------------------------------------------- |
| namenode     | [源码] 文件元数据服务, 管理文件的元数据信息和文件信息                                       |
| datanode     | [源码] 文件存取服务, 数据存储服务, 可部署多个                                              |
| opensdk      | [源码] 开放SDK, 提供了系统的核心服务接口封装                                               |
| tools        | [源码] 工具集合包, 可对系统进行管理和测试                                                  |
| bin          | [编译结果] 所有服务编译好的二进制可执行文件                                                 |
