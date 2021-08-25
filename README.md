## 毕业项目设计

### 项目介绍

现在的项目是一个媒体系统的BFF层，有若干上游系统。
整个公司的系统现在是建立在AWS上，已经分离出来了若干SOA层面，每个系统之间通过内部web service相互访问。

![系统结构图](docs/System-Current.svg)

我们项目提供了上游数据的缓存服务以及面向客户端的API服务，需要将上游数据聚合/转换成客户端需要的JSON格式。

公司所有的内容都是公共许可的，所以不需要用户认证，都可以公开访问。内部维护是通过另一套系统经过VPN访问的，不在我们相关的系统范围内。

### 微服务改造

系统本身已经是面向服务的架构，BFF层不适合做太多拆分，所以改造幅度不是很大。
* 将上游系统到BFF层的通信改为grpc，减少通信延时。
* 使用客户端负载均衡访问上游服务。
* 使用kafka替换现在的通知系统，增加系统容量。
* 上游内容聚合使用更好的并发结构。

![系统结构图](docs/System-Modified.svg)
