package router

import (
	"goat-layout/internal/api"
	"goat-layout/internal/api/v1beta1"
	"goat-layout/internal/model"
	"goat-layout/pkg/log"

	"github.com/gin-gonic/gin"
)

func registerBookRouter(store *model.Store, r *gin.RouterGroup) {
	book := v1beta1.Book{
		Base: api.New(store, log.New("book").L()),
	}
	bookR := r.Group("/book")

	{
		bookR.GET("", book.GetList)
		bookR.POST("", book.Add)
	}
}
