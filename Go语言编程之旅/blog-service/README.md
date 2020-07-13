## 项目目录结构
```
.
├── configs
├── doc
├── global
├── go.mod
├── go.sum
├── internal
│   ├── dao
│   ├── middleware
│   ├── model
│   ├── routers
│   └── service
├── main.go
├── pkg
├── scripts
├── storage
└── third_party
```
* configs: 配置文件
* doc: 文档集合
* global: 全局变量
* internal: 内部模块
    * dao: 数据库访问层(Database Access Object)
    * middleware: HTTP 中间件
    * model: 模型层,用户存放 model 对象
    * routers: 路由相关的逻辑
    * service: 项目核心业务逻辑
* pkg: 项目相关的模块包
* storage: 项目生成的临时文件
* scripts: 各类构建,安装,分析等操作的脚本
* third_party: 第三方的资源工具
 
