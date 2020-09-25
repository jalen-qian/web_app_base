## 基于Gin框架的Go Web开发脚手架
这个仓库是基于Gin框架搭建的Go Web开发通用脚手架，里面根据CLD架构做了分层。

其中对一些常用的第三方库做了初始化操作，例如：

- 集成了基于Zap日志库的日志系统
- 集成了基于Viper的配置管理
- 基于http.Server.Shutdown()方法做了优雅关机或者重启
- 基于Gorm做的mysql连接和数据分层
- 集成了Redis连接和常用操作
- 集成了基于雪花算法的ID生成工具