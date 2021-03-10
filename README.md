# MiraiGo-Bot
增强MiraiGo-Bot,开箱可用,降低使用成本
## 运行
### 宿主机
release页面直接下载对应的程序,运行
### docker
* 自行构建docker镜像
* `docker run -it -v ${数据目录}:/data xxx:xxx`
* 根据完成账号密码填写和设备验证,停止容器
* `docker run -d --restart=always -v ${数据目录}:/data xxxx`

## 功能

* [ ] 实现[MiraiGo-Template](https://github.com/Logiase/MiraiGo-Template)
* [ ] 提供`plugins`,统一消息输入输出
  * [ ] `plugins`接口设计
  * [ ] `plugins`注册功能
* [ ] 在`MiraiGo-Template`插入中间层,监听消息,封装`struct`后,调用实现的插件
  * [ ] 中间层将实现`命令`=>`自定义插件`的配置
* [ ] 提供默认插件
  * [ ] 配置插件
  * [ ] 天气插件
  * [ ] 一言插件
  * [ ] 骰子插件
* [ ] 提供配置命令
  * [ ] 当前插件列表
  * [ ] 配置插件命令
* [ ] ~~提供配置界面~~
## 二次开发

1. 引入本项目,实现[`Plugin接口`](./pkg/plugins/entity.go)
2. 调用注册方法,将当前插件注册到接口中心
3. 仿照[`cmd/main.go`](./cmd/main.go)编写项目启动类
4. 对插件和命令进行配置,绑定`命令`=>`插件`


## 依赖项目
* https://github.com/Mrs4s/MiraiGo
* https://github.com/Logiase/MiraiGo-Template
* https://github.com/golang-migrate/migrate/v4
* https://github.com/mattn/go-sqlite3
* ... 更多见[`go.mod`](go.mod)