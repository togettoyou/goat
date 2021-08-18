# api

api 接口层（或控制层），负责参数校验解析，请求响应处理，不应处理复杂逻辑

教程：

1. 创建 api 对象，需继承 api.Base

   ```go
   type Book struct {
   	api.Base
   }
   ```

2. 定义方法

   ```go
   // bad 切忌，不可使用指针接收器，对于每一个请求，应该使用值接收器创建新的对象
   // (g *Book)传递的是地址，而 MakeContext和MakeService会进行上下文和业务绑定操作
   func (g *Book) GetList(c *gin.Context) {
   	bookSvc := svc.Book{}
   	g.MakeContext(c).MakeService(&bookSvc.Service)
   	g.Log.Info("路由处理")
   	books, err := bookSvc.GetList()
   	if g.HasErrL(err) {
   		return
   	}
   	g.OK(books)
   }
   
   // good 正确做法
   func (g Book) GetList(c *gin.Context) {
   	bookSvc := svc.Book{}
   	g.MakeContext(c).MakeService(&bookSvc.Service)
   	g.Log.Info("路由处理")
   	books, err := bookSvc.GetList()
   	if g.HasErrL(err) {
   		return
   	}
   	g.OK(books)
   }
   ```

3. 路由绑定

   ```go
   book := v1beta1.Book{
       Base: api.New(store, log.New("book").L()),
   }
   r.GET("/", book.GetList)
   ```

   
