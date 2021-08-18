package router

import (
	"net/http"

	"goat-layout/internal/api"
	"goat-layout/internal/api/v1beta1"
	"goat-layout/internal/model"
	"goat-layout/internal/server/middleware"
	"goat-layout/pkg/log"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var swag gin.HandlerFunc

func New(store *model.Store) *gin.Engine {
	r := gin.New()
	useMiddleware(r)
	initDebugRouter(r)
	initSwagRouter(r)
	initV1beta1(store, r)
	return r
}

// useMiddleware 全局中间件
func useMiddleware(r *gin.Engine) {
	r.Use(gin.Recovery(), middleware.Cors(), middleware.Logger())
}

// initDebugRouter debug/pprof
func initDebugRouter(r *gin.Engine) {
	debugRouter := r.Group("debug", func(c *gin.Context) {
		if !gin.IsDebugging() {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Next()
	})
	pprof.RouteRegister(debugRouter, "pprof")
}

// initSwagRouter swag 文档，可控制编译
func initSwagRouter(r *gin.Engine) {
	if swag != nil {
		r.GET("/swagger/*any", swag)
	}
}

func initV1beta1(store *model.Store, r *gin.Engine) {
	v1beta1Group := r.Group("api/v1beta1")
	{
		book := v1beta1.Book{
			Base: api.New(store, log.New("book").L()),
		}
		bookR := v1beta1Group.Group("/book")
		{
			bookR.GET("", book.GetList)
		}
	}

	{
		example := v1beta1.Example{
			Base: api.New(nil, log.New("example").L()),
		}
		exampleR := v1beta1Group.Group("/example")
		{
			exampleR.GET("", example.Get)
			exampleR.GET("/err", example.Err)
			exampleR.GET("/uri/:id", example.Uri)
			exampleR.GET("/query", example.Query)
			exampleR.POST("/form", example.FormData)
			exampleR.POST("/json", example.JSON)
		}
	}

}
