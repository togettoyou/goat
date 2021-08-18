# svc 即 service

业务逻辑层，负责处理复杂业务逻辑，处于 api 层和 dao 层之间。svc 只能通过 dao 层获取数据


教程：

1. 创建 svc 对象，需继承 Service

   ```go
   type Book struct {
   	Service
   }
   ```

2. 定义方法

   ```go
   func (b *Book) GetList() ([]model.Book, error) {
   	b.log.Info("业务处理")
   	// 使用store调用dao层
   	books, err := b.store.Book.List()
   	if err != nil {
   		// 返回包装错误，包含调用栈信息
   		return nil, e.New(e.DBError, err)
   	}
   	return books, nil
   }
   ```
   
