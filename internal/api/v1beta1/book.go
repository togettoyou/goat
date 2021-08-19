package v1beta1

import (
	"goat-layout/internal/api"
	"goat-layout/internal/svc"

	"github.com/gin-gonic/gin"
)

type Book struct {
	api.Base
}

// GetList
// @Tags book
// @Summary 获取书籍列表
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} api.Response
// @Router /api/v1beta1/book [get]
func (g Book) GetList(c *gin.Context) {
	bookSvc := svc.Book{}
	g.Named("GetList").MakeContext(c).MakeService(&bookSvc.Service)
	g.Log.Info("路由处理")
	books, err := bookSvc.GetList()
	if g.HasErrL(err) {
		return
	}
	g.OK(books)
}

type bookBody struct {
	Name string `json:"name" binding:"required"`
	Url  string `json:"url" binding:"required"`
}

// Add
// @Tags book
// @Summary 新增书籍
// @Security ApiKeyAuth
// @Produce json
// @Param data body bookBody true "测试请求json参数"
// @Success 200 {object} api.Response
// @Router /api/v1beta1/book [post]
func (g Book) Add(c *gin.Context) {
	var body bookBody
	bookSvc := svc.Book{}
	if !g.Named("Add").MakeContext(c).MakeService(&bookSvc.Service).ParseJSON(&body) {
		return
	}
	g.Log.Info("路由处理")
	if g.HasErrL(bookSvc.Add(body.Name, body.Url)) {
		return
	}
	g.OK()
}
