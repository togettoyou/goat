package router

import (
	"net/http"

	"goat/internal/model"
	"goat/internal/server/api"
	"goat/internal/server/api/v1beta1"
	"goat/internal/server/middleware"
	"goat/pkg/log"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var swag gin.HandlerFunc

func New(store *model.Store) *gin.Engine {
	r := gin.New()
	useMiddleware(r)
	initDebugRouter(r)
	initSwagRouter(r)

	v1beta1Group := r.Group("api/v1beta1")
	initV1beta1(store, v1beta1Group)
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

func initV1beta1(store *model.Store, r *gin.RouterGroup) {
	book := v1beta1.Book{
		Base: api.New(store, log.New("book").L()),
	}
	bookR := r.Group("/book")
	{
		bookR.GET("", book.GetList)
	}
}
