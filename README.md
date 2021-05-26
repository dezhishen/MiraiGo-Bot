# MiraiGo-Bot
增强MiraiGo-Template,开箱可用,降低使用成本,降低开发成本

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
* [x] 提供工具类
  * [x] 键值对持久化存储 [storage](./pkg/storage)
    * [x] 数据隔离
  * [x] 键值对缓存 [cache](./pkg/cache)
  * [x] 命令行Parse工具 [command](./pkg/command)
  * [ ] 全局键值对
* [x] 在`MiraiGo-Template`插入中间层,调用实现的插件
  * [x] 消息监听插件支持
    * [x] 中间层将按照插件的顺序,依次触发插件的`OnMessageEvent`方法
    * [x] 中间层的可配置
  * [x] 定时任务插件支持
    * [x] 启动和运行定时任务 
* [x] 健康检查 运行目录下,文件`/data/health`是否存在

### 部分功能使用说明
#### 命令行解析
使用[https://github.com/jessevdk/go-flags](https://github.com/jessevdk/go-flags)实现

查看示例:[command_test.go](./pkg/command/command_test.go)

* 首先定义一个`struct`
* 补充 `short`,`long`,`default`,`description`等关键字信息
  ```
  type opts struct {
    Name string `short:"n" long:"name" description:"姓名"`
  }
  ```
* 调用
  ```
  func Test_Parse(t *testing.T) {
    var o = opts{}
    args, _ := Parse(&o, []string{"doSomething", "-n", "test"})
    print(len(args))
    print(o.Name)
  }
  ```

## 运行

编写启动类,引入需要加载的插件,调用启动方法

```
package main

import (
// 引入插件
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/calendar"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/haimage"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/hitokoto"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/jrrp"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/lpl"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/mc"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/pixiv"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/random"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/thecat"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/thedog"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/tips"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/caihongpi"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/dujitang"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/weather"
	_ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/sovietjokes"
 //启动入口
	"github.com/dezhiShen/MiraiGo-Bot/pkg/server"
)

func main() {
//启动
	server.Start()
}
```
## 插件仓库
https://github.com/dezhiShen/MiraiGo-Bot-Plugins

## 二次开发

1. 引入本项目,实现[一个或者多个`plugin`](./pkg/plugins/plugin.go)
2. 在`init()`方法中调用注册方法`plugins.RegisterXXXPlugin(YourPlugin{})`,将当前插件注册到系统中
3. 启动

    ```
    package main

    import (
      _ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/hitokoto"
      _ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/random"
      _ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/tips"
      _ "github.com/dezhiShen/MiraiGo-Bot-Plugins/pkg/plugins/weather"
      "github.com/dezhiShen/MiraiGo-Bot/pkg/server"
    )

    func main() {
      server.Start()
    }


    ```

## 依赖项目
* https://github.com/Mrs4s/MiraiGo
* https://github.com/Logiase/MiraiGo-Template
* ... 更多见[`go.mod`](go.mod)
