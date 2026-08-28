# 你是谁



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
- go语言
    - 使用gin，作为接入曾，接口方法名用Action结尾;
        - 接口只允许POST请求
        - 参数使用gin的BindJSON方法;
        - 返回格式为{"code":200,"msg":"success","data":Any}，
        - 需要验证登陆状态的接口，若是为登陆返回的code=333
        
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
├── docs
│   ├── api.md
│   ├── design.md
│   └── develop.md
├── mobile
└── server
│   ├── app
│   │   ├── actions
|   │   │   ├── api_mobile
|   |   │   │   ├── account.go
|   |   │   │   └── common.go
|   │   │   ├── api_admin
|   |   │   │   ├── account.go
|   |   │   │   ├── system.go
|   |   │   │   └── jobs.go
|   │   │   └── api_web
|   |   │       └── account.go
│   │   ├── models
|   │   │   ├── account.go
|   |   │   └── system.go
│   │   ├── logics
|   │   │   ├── req
|   |   │   │   ├── req_mobile
|   |   │   │   │   ├── account.go
|   |   │   │   │   └── common.go
|   |   │   │   ├── req_admin
|   |   │   │   │   ├── account.go
|   |   │   │   │   ├── system.go
|   |   │   │   │   ├── jobs.go
|   |   │   │   │   └── common.go
|   |   │   │   └── req_web
|   |   │   │       ├── account.go
|   |   │   │       └── common.go
|   |   │   ├── resp
|   |   │   │   ├── resp_mobile
|   |   │   │   ├── resp_admin
|   |   │   │   └── resp_web
|   |   │   ├── services
|   |   │   │   ├── svc_account
|   |   │   │   └── svc_system
|   |   │   └── define.go
│   │   ├── middlewares
|   │   │   ├── api_middleware.go
|   |   │   ├── admin_middleware.go
|   |   │   └── web_middleware.go
|   │   ├── tasks
|   |   │   ├── job.go
|   |   │   └── consumer.go
|   |   │   ├── jobs
|   |   │   └── consumers
│   │   ├── utils
|   │   │   ├── common.go
|   |   │   ├── http.go
|   |   │   ├── jwt.go
|   |   │   ├── log.go
|   |   │   ├── str.go
|   |   │   ├── format.go
|   |   │   ├── md5.go
|   |   │   ├── file.go
|   |   │   ├── time.go
|   |   │   └── validate.go
│   │   ├── conf
|   │   │   ├── conf.go
|   │   │   ├── test.yaml
|   │   │   └── prod.yaml
│   │   ├── views
|   │   │   ├── www
|   │   │   ├── admin
|   │   │   └── h5
│   │   └── bootstrap.go
│   └── main.go
```

# 约束
约束：
1. 必须传递context.Context；
2. 所有error不能忽略，必须处理；
3. 统一返回结构体；
4. 增加panic recover防止服务崩溃；
5. 不要使用_忽略任何错误；
输出完整可编译代码。

编写goroutine业务逻辑，严格遵守：
1. 使用context控制协程生命周期，具备退出机制；
2. WaitGroup必须在goroutine外部Add；
3. 禁止循环闭包捕获循环变量；
4. 所有错误要打印日志；
5. 禁止产生goroutine泄漏；
输出完整Go代码。

评审下面Go代码，输出问题清单：并发风险、错误处理遗漏、内存泄漏风险、安全问题、性能问题；给出修复后的diff代码。
代码：
【粘贴代码】

# 开发流程
- 阅读需求文档，确认需求；
- 阅读 docs目录下的文档和目录源码，确认修改内容
- 生产具体开发任务
- 创建 git 开发分支
- 编写代码
- 编写测试用例
- 编译验证
- 运行测试用例
- 修改 /docs/swagger 文档，修改 /docs/api.md 文档，修改 docs/design.md 文档,修改 docs/develop.md 文档
- 生产代码评审任务
- 代码评审
- 合并代码
- 提交代码
