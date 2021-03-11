# MiraiGo-Bot
增强MiraiGo-Bot,开箱可用,降低使用成本,降低开发成本
## 运行
### 宿主机
release页面直接下载对应的程序,运行
### docker
* 自行构建docker镜像 | 或者使用镜像`1179551960sdniu/miraigo:0.02`
* `docker run -it -v ${数据目录}:/data 1179551960sdniu/miraigo:0.02`
* 根据提示完成账号密码填写和设备验证,停止容器
* `docker run -d --restart=always -v ${数据目录}:/data 1179551960sdniu/miraigo:0.02`

## 功能

* [x] 实现[MiraiGo-Template](https://github.com/Logiase/MiraiGo-Template)
* [x] 提供`plugins`接口
  * [x] `plugin`基类
    * [x] `PluginInfo()`插件基础信息
    * [x] `PluginInit()`插件初始化
  * [x] `OnMessagePlugins`消息处理插件
    * [x] 接口设计
      * [x] `SortNum()`触发的顺序
      * [x] `IsFireEvent(msg *MessageRequest) bool`是否触发
      * [x] `OnMessageEvent(msg *MessageRequest) (*MessageResponse, error)`方法
      * [x] `IsFireNextEvent(msg *MessageRequest) bool`是否继续触发后续的插件
    * [x] 注册功能
  * [x] `SchedulerPlugins`定时器插件
    * [x] 接口设计
      * [x] `Cron()` 返回cron表达式
      * [x] `Run(bot *bot.Bot) error` 执行回调
    * [x] 注册功能
* [ ] 提供工具类
  * [ ] 键值对持久化存储
  * [ ] 键值对缓存
  * [ ] 数据隔离
  * [ ] 全局键值对
* [x] 在`MiraiGo-Template`插入中间层,监听消息,封装`struct`后,调用实现的插件
  * [x] 中间层将按照插件的顺序,依次触发插件的`OnMessageEvent`方法
  * [x] 中间层的可配置
* [ ] 提供默认插件
  * [ ] 配置插件(包含权限要求)
  * [x] 天气插件
  * [x] 一言插件
  * [x] 骰子插件
* [ ] ~~提供配置界面~~

## 二次开发

1. 引入本项目,实现[一个或者多个`plugin`](./pkg/plugins/plugin.go)
2. 在`init()`方法中调用注册方法`plugins.RegisterXXXPlugin(YourPlugin{})`,将当前插件注册到系统中
3. 启动

    ```
    package main

    import (
      // 引入插件
      _ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins"
      "github.com/dezhiShen/MiraiGo-Bot/pkg/server"
    )

    func main() {
      server.Start()
    }


    ```

## 插件仓库
* https://github.com/dezhiShen/MiraiGo-Bot-Plugins

## 依赖项目
* https://github.com/Mrs4s/MiraiGo
* https://github.com/Logiase/MiraiGo-Template
* ... 更多见[`go.mod`](go.mod)
