### 项目介绍
基于 Golang 的 go-gin项目，使用了 redis + mysql + swagger + jwt。

项目来自博客 [跟煎鱼学 Go](https://eddycjy.gitbook.io/golang/) 。

### Go 项目结构
```
go-gin-example/
 ├── conf
 ├── middleware
 ├── models
 ├── pkg
 ├── routers
 ├── runtime
 └── service
``` 

- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
- runtime：应用运行时数据
- service：biz层，访问缓存
