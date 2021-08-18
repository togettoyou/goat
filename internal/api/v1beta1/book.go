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
