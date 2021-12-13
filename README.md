# goat

goat 是基于 gin 进行快速构建 RESTFUL API 的项目框架

# 脚手架安装

```shell
go install github.com/togettoyou/goat/cmd/goatkit@latest
```

Or

```shell
go get -u github.com/togettoyou/goat/cmd/goatkit
```

# 使用

```
# 创建项目模板
goatkit new helloworld

cd helloworld
# 运行程序
goatkit run

# 生成 swag 文档
goatkit swag

# 使用镜像代理
export GOAT_LAYOUT_REPO=https://ghproxy.com/https://github.com/togettoyou/goat.git
goatkit new helloworld

# 更多帮助
goatkit -h
```

# 文档

项目的依赖路径为: [model 层](internal/model/README.md) -> [dao 层](internal/dao/README.md) -> [svc 层](internal/svc/README.md)
-> [api 层](internal/api/README.md)

其中 model struct 串联各个层

[router 层](internal/server/router/README.md) 负责路由注册
