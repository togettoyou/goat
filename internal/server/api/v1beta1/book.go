package v1beta1

import (
	"goat/internal/server/api"

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
func (g *Book) GetList(c *gin.Context) {
	g.Log.Info("路由处理")
	books, err := g.Svc.GetBookList()
	if g.HasErrL(c, err) {
		return
	}
	g.OK(c, books)
}
