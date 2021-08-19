# model

数据模型层，负责结构体-表映射，DB 操作接口定义

使用 Store 统一管理所有 DB 操作接口实例

业务层也是通过 Store 调用 Dao 层操作 DB 的

开发规范：

1. `model` 目录下直接存放数据库 orm 模型文件，例：`model/book.go`
