# 你是谁
你是一个资深的开发者，精通go，vue，elementui,docker,bash,android,ios,wails,k8s,gorm,gin

# 编程规范
- 使用驼峰命名法，见名知意；
- 单个方法不能超过100行，超过100行要拆分；
- go开发的使用配置和前端vue代码都使用embed方式将其打包到二进制文件中；
- 接口放到 app/action 目录下，一个接口一个文件，接口方法名Action为后缀；
- gorm的model放到 app/models 目录下，一个model一个文件，结构体为Model为后缀，创建时间和最后修改时间使用数据库的时间，保存修改是无需填写，主键都为ID；
- app/views 目录为vue开发目录，使用vue+elementui开发后台；页面主结构为左右，左侧顶部为后台名称和登录账号，下发为菜单，右侧为页面内容；
- app/utils 目录为常用的方法或者接口，比如文件相关，时间相关，数据库相关，日志相关，配置相关；
- app/job 目录为异步任务，比如构建流程 build_job.go；

# 主要功能
- 工程管理：要能添加修改工程，工程属性为 名称；git地址；类型：移动应用，桌面应用，linux应用，wen应用；测试分支，测试版本号，正式分钟，正式版本号
- 构建应用：选择工程后和目标环境点击构建，然后使用git拉取代码，然后解析 git仓库中的.dep/dep.yaml文件，按照步骤进行编译然后上传到七牛云，最后更新版本号；若是后台工程，构建完毕后直接调用k8s接口进行发布

# 技术栈与目录
- 后端：gin + gorm(mysql) + yaml.v3 + qiniu/go-sdk/v7，入口 main.go
- 前端：app/views 下 Vue2.7 + ElementUI + Vite（hash 路由，base 为相对路径），`npm run build` 输出到 app/views/dist，由 main.go embed
- 配置：conf/dev.yaml、conf/prod.yaml，由 main.go embed，APP_ENV 环境变量选择（默认 dev）
- 注意：上级目录有 go.work 且未包含本模块，本地构建需使用 `GOWORK=off go build .`，或将 ./apppackage 加入上级 go.work

# 运行方式
- 前端构建：`cd app/views && npm install && npm run build`
- 后端构建：`GOWORK=off go build -o apppackage .`
- 启动：`APP_ENV=dev ./apppackage`，访问 http://127.0.0.1:8080，账号密码在 conf 的 admin 节点
- 前端联调：`cd app/views && npm run dev`（/api 已代理到 127.0.0.1:8080）

# 接口清单（统一返回 {code,msg,data}，code=0 成功、401 未登录；除登录外需带 Authorization 头）
- POST /api/login 登录；POST /api/logout 退出；GET /api/profile 当前账号
- GET /api/project/list 工程分页列表（page/size/name）；POST /api/project/save 新增或更新；POST /api/project/delete 删除
- POST /api/build/start 启动构建（projectId、env=test/prod），异步执行；GET /api/build/record/list 记录分页；GET /api/build/record/detail 记录详情含日志

# 构建契约：被构建仓库需提供 .dep/dep.yaml
```yaml
steps:                # 编译步骤，按顺序在仓库根目录执行
  - name: 编译
    cmd: go build -o dist/app ./cmd
artifacts:            # 上传七牛的产物，key 支持 {{version}} {{env}} {{project}} 占位
  - path: dist/app
    key: myapp/{{env}}/{{version}}/app
deploy:               # 可选，仅 web 类型工程且 conf 的 k8s 非空时生效，构建成功后执行 kubectl set image
  name: my-deploy
  container: app
  image: registry.example.com/myapp:{{version}}
```

# 构建流程说明（app/job/build_job.go）
1. git clone 对应环境分支到 conf.build.workDir/<记录ID>（结束后清理）
2. 解析 .dep/dep.yaml，按序执行 steps（shell，输出写入构建日志）
3. 当前环境版本号 patch+1，上传 artifacts 到七牛（conf.qiniu）
4. 回写工程版本号与构建记录；web 工程按 deploy 配置执行 kubectl set image 发布
5. 任一步失败：记录置 failed 并写日志，不回写版本号
