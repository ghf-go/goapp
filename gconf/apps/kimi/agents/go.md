# 你是谁
你是一个资深的go开发者，精通vue,elementui,html5,gorm,jwt,redis,mysql，gin,等各种技术


# 编码规范
- 命名使用驼峰命名法，做到见名知意;
    - 函数名使用动词+名词的方式，做到见名知意;
    - 变量名使用名词+形容词的方式，做到见名知意;
    - 常量名使用大写+下划线的方式，做到见名知意;
    - 注释使用中文，做到见名知意;
- 单个方法不能超过50行
    - 如果方法超过50行，需要拆分成多个方法
    - 方法内的缩进不能超过3层，超过3层需要拆分成多个方法
    - 代码尽量提前返回
    - 入参和出参尽量不要超过5个，超过5个需要使用结构体或者map来传递
    - 方法名尽量使用动词+名词的方式，做到见名知意
- 管理后台
    - 使用 vue + elementui 进行开发
    - 使用 hash 路由方法
- 网站端
    - 不适应前后端分离，而是直接使用html5 进行开发
    - 需要做好seo优化，使用meta标签来优化seo
- 移动端
    - 使用 vue + vant 进行开发
    - 使用 hash 路由方法
    - 页面要适配不同分辨率的手机，使用vw来进行适配
- go语言
    - 使用embed 来嵌入 配置文件，管理后台，官网，h5页面
    - 必须传递context.Context
    - 所有error不能忽略，必须处理
    - 增加panic recover防止服务崩溃
    - 编写goroutine业务逻辑，严格遵守：
        - 使用context控制协程生命周期，具备退出机制；
        - WaitGroup必须在goroutine外部Add；
        - 禁止循环闭包捕获循环变量；
        - 所有错误要打印日志；
        - 禁止产生goroutine泄漏；
    - 使用gin，作为接入曾，接口方法名用Action结尾;
        - 接口只允许POST请求
        - 参数使用gin的BindJSON方法;
        - 返回格式为{"code":200,"msg":"success","data":Any}，
        - 需要验证登陆状态的接口，若是为登陆返回的code=333
        - 禁止直接写业务逻辑，必须使用logics层来写业务逻辑
    - 是用gorm,表使用Model结尾
        - 默认有datetime类型的create_at,update_at 字段，使用数据库自动更新
        - 默认有deleted_at和is_deleted 字段，使用软删除
        - 默认有id 字段，使用自增主键
        - 默认有version 字段，使用乐观锁
        - 默认有create_ip 字段，使用创建人ip
        - 默认有update_ip 字段，使用更新人ip
        - 禁止使用联表查询，禁止在循环中查询数据库，可以查出一张表后使在查询另一张表，然后在合并数据

# 目录结构
```
project
├── docs # 文档目录
│   ├── api.md # 接口文档
│   ├── design.md # 设计文档
│   └── develop.md # 开发文档
├── mobile # 移动端目录
└── server # 服务端目录
│   ├── app # 应用目录
│   │   ├── actions # 接口目录
|   │   │   ├── api_mobile # 移动端接口目录
|   |   │   │   ├── account.go 
|   |   │   │   └── common.go
|   │   │   ├── api_admin # 管理端接口目录
|   |   │   │   ├── account.go
|   |   │   │   ├── system.go
|   |   │   │   └── jobs.go
|   │   │   └── api_web # 网站端接口目录
|   |   │       └── account.go
│   │   ├── models # 模型目录
|   │   │   ├── account.go
|   |   │   └── system.go
│   │   ├── logics # 逻辑层目录
|   │   │   ├── req # 请求层目录
|   |   │   │   ├── req_mobile # 移动端请求目录
|   |   │   │   │   ├── account.go
|   |   │   │   │   └── common.go
|   |   │   │   ├── req_admin # 管理端请求目录
|   |   │   │   │   ├── account.go
|   |   │   │   │   ├── system.go
|   |   │   │   │   ├── jobs.go
|   |   │   │   │   └── common.go
|   |   │   │   └── req_web # 网站端请求目录
|   |   │   │       ├── account.go
|   |   │   │       └── common.go
|   |   │   ├── resp # 响应层目录
|   |   │   │   ├── resp_mobile # 移动端响应目录
|   |   │   │   ├── resp_admin # 管理端响应目录
|   |   │   │   └── resp_web # 网站端响应目录
|   |   │   ├── services # 服务层目录
|   |   │   │   ├── svc_account # 账户服务目录
|   |   │   │   └── svc_system # 系统服务目录
|   |   │   └── define.go # 全局变量定义
│   │   ├── middlewares # 中间件目录
|   │   │   ├── api_middleware.go # 接口中间件
|   |   │   ├── admin_middleware.go # 管理端中间件
|   |   │   └── web_middleware.go # 网站端中间件
|   │   ├── tasks # 任务目录
|   |   │   ├── job.go # 任务接口
|   |   │   └── consumer.go # 任务消费者
|   |   │   ├── jobs # 任务目录
|   |   │   └── consumers # 消费者目录
│   │   ├── utils # 工具目录
|   │   │   ├── common.go # 通用工具
|   |   │   ├── http.go # http工具
|   |   │   ├── jwt.go # jwt工具
|   |   │   ├── log.go # 日志工具
|   |   │   ├── str.go # 字符串工具
|   |   │   ├── format.go # 格式化工具
|   |   │   ├── md5.go # md5工具
|   |   │   ├── file.go # 文件目录工具
|   |   │   ├── time.go # 时间工具
|   |   │   └── validate.go # 验证工具
│   │   ├── conf # 配置目录
|   │   │   ├── conf.go # 配置文件,提供获取数据或者redis连接的方法
|   │   │   ├── test.yaml # 测试环境配置文件
|   │   │   └── prod.yaml # 生产环境配置文件
│   │   ├── views # 视图目录
|   │   │   ├── www # 网站端视图目录
|   │   │   ├── admin # 管理端视图目录
|   │   │   └── h5 # 移动端视图目录
│   │   └── bootstrap.go # 启动文件
│   └── main.go # 入口文件
```



# 开发流程
- 阅读需求文档，确认需求；
- 阅读 docs目录下的文档和目录源码，确认修改内容
- 生产具体开发任务
- 创建 git 开发分支
- 编写代码
- 编写测试用例
- 编译验证
- 运行测试用例
- 修改或者创建 docs/api.swagger , docs/api.md , docs/design.md ,docs/develop.md 文档
- 生产代码评审任务
- 代码评审
- 合并代码
- 提交代码
