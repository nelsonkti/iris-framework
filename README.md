```
 _      _            __                                             _    
(_)_ __(_)___       / _|_ __ __ _ _ __ ___   _____      _____  _ __| | __
| | '__| / __|_____| |_| '__/ _` | '_ ` _ \ / _ \ \ /\ / / _ \| '__| |/ /
| | |  | \__ \_____|  _| | | (_| | | | | | |  __/\ V  V / (_) | |  |   < 
|_|_|  |_|___/     |_| |_|  \__,_|_| |_| |_|\___| \_/\_/ \___/|_|  |_|\_\
                                                                                                                                                        
```
[iris-framework](https://github.com/nelsonkti/iris-framework) 是基于 iris 搭建用于快速开发的项目框架

## 安装
```
go get -u github.com/nelsonkti/iris-framework
```

运行
```
cd iris-framework

go run main.go
```

## 功能叙述
- 支持 gorm、xorm、logger 日志、jwt 等
- 日志切割
- 支持mysql

## 文件夹结构 
  - app 应用目录，包含了http、模型等
  - config 项目的应用配置文件
  - database 数据迁移
  - public 资源文件（图片、文件等）
  - routes 路由
  - storage 缓存和日志文件
、
## 环境要求 

- go >= 1.13