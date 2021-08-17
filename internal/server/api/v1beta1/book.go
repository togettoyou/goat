package v1beta1

import (
	"goat/internal/server/api"

	"github.com/gin-gonic/gin"
)

type Book struct {
	api.Base
}

// GetList
// @Tags v1beta1
// @Summary 获取书籍列表
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} api.Response
// @Router /api/v1beta1/book [get]
func (b *Book) GetList(c *gin.Context) {
	b.Log.Info("路由处理")
	books, err := b.Svc.GetBookList()
	if b.HasErrL(c, err) {
		return
	}
	b.OK(c, books)
}
