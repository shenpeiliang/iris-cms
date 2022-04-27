#### 简介

号称全宇宙最快的 Web 框架 Iris 

地址：

https://github.com/kataras/iris


#### 功能

- 用户登录状态

- 图文类常用的图片上传，常用的增删改查

#### 配置说明

- .env 配置静态文件版本
```
STATIC_VERSION=1.0.7
```

- config/config.yml other项服务配置等

- .air.conf 配置热重启，修改代码后自动编译，安装命令
```
go get -u github.com/cosmtrek/air
```

#### 运行

下载需要的扩展
```
go mod tidy
```

需要导入数据库文件

```
data/cms.sql
```

使用 GORM 的 gen 工具生成表结构体,表结构发生变化需要再次执行同步
```
go get -u gorm.io/gen

./generation.sh
```

启动服务
```
go run main.go
```

热编译方式：
```
air -c .air.conf
```

访问地址：
```
http://localhost:8080
```

#### 效果图

![image](static/images/action.gif)


#### 其他

国内优秀的框架推荐 goframe，框架开发者自称可以像 php 一样简单（确实是）

https://goframe.org/pages/viewpage.action?pageId=1114399

国内强大的开源框架 go-zero 

https://github.com/zeromicro/go-zero

![image](static/images/go-zero.png)