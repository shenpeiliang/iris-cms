#### 功能

- 用户登录状态

- 图文类常用的图片上传，常用的增删改查

#### 配置说明

- .env 配置数据库等
```
STATIC_VERSION=1.0.7
HOST=127.0.0.1
PORT=3306
USER=root
PASSWORD=root
DATABASE=cms
CHARSET=UTF8
TABLE_PREFIX=hs_
```

- config/config.yml other项配置业务

- .air.conf 配置热重启，修改代码后自动编译，安装命令
```
go get -u github.com/cosmtrek/air
```

#### 运行

需要导入数据库文件

```
data/cms.sql
```

运行程序：
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