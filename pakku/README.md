# pakku 帕克

    pakku 实现了对象的生命周期管理、依赖自动注入等功能, 无三方引用依赖, 可快速搭建一个服务或程序.

## pakku 中的对象

    pakku 在加载实现了`ipakku.Module`、`ipakku.Controller`、`ipakku.Router`接口的对象以及RPC注册对象时, 会对内部的成员变量(包括私有)进行扫描, 对包含`@autowired:"接口名"`标签的字段进行自动赋值, 实现了模块间的解耦和依赖注入. 其中只有`ipakku.Module`可以注入其他对象和被注入别的`ipakku.Module`, 其他对象的只能被注入`ipakku.Module`.

* `Module` 即模块, 通常为有一定功能单元的对象, 在`Controller`、`rpc服务`、`其他Module`中被注入和调用.
* `Router` 即http路由定义模块, 通过定义接口的url、地址列表实现对url的路由.
* `Controller` 即http服务定义模块, 通过定义接口的url、地址列表、过滤器等实现对url的路由.
* `RpcService` 即rpc服务定义模块, 默认使用自带的rpc服务进行注册.

## 默认实现的模块

    pakku 默认实现了AppConfig(配置模块)、AppCache(缓存模块)、AppEvent(事件模块)、AppService(NET服务模块)以满足一个基本的服务运行环境. 所有默认实现的模块, 均通过接口调用, 在未使用的情况下, 不会对整体造成影响. 如需使用默认的接口又需要使用其他方式实现, 可通过重新实现`ipakku/...`下的接口后, 重新指定默认调用实例即可(AppService模块暂不支持自定义实现). 也可以不使用默认实现的模块, 使用手工指定模块的方式加载.

### AppConfig 配置模块

    默认json格式的配置器, 文件存放在启动目录下`./.conf/{appName}.json`中, 可以通过启动前指定符合`ipakku.IConfig`接口的实现类来复写

### AppCache 缓存模块

    默认实现的本机缓存模块, 可以通过启动前指定符合`ipakku.ICache`接口的实现类来复写, 如: redis等

### AppEvent 事件模块

    默认的本机事件模块, 默认实现异步事件, 可以通过启动前指定符合`ipakku.IEvent`接口的实现类来复写, 如: kafka等

### AppService NET服务模块

    HTTP服务
        .Get(url string, fun ipakku.HandlerFunc) error 
        .Post(url string, fun ipakku.HandlerFunc) error
        .Put(url string, fun ipakku.HandlerFunc) error
        .Patch(url string, fun ipakku.HandlerFunc) error
        .Head(url string, fun ipakku.HandlerFunc) error
        .Options(url string, fun ipakku.HandlerFunc) error
        .Delete(url string, fun ipakku.HandlerFunc) error
        .Any(url string, fun ipakku.HandlerFunc) error
        .AsRouter(url string, router ipakku.AsRouter) error
        .AsController(router ipakku.AsController) error
        .Filter(url string, fun ipakku.FilterFunc) error
        .SetStaticDIR(path, dir string, fun ipakku.FilterFunc)  error
        .SetStaticFile(path, file string, fun ipakku.FilterFunc) error

    RPC服务
        .RegisteRPC(rcvr interface{}) error

## 如何使用 pakku

    pakku 是一个核心库, 提供灵活的编码方式的同时也具有一定的复杂性, 可参考`pakku_test.go`和`modules//*`下的模块编写. 
