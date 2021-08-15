package router

import (
	"net/http"

	"goat/internal/server/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var swag gin.HandlerFunc

func New() *gin.Engine {
	r := gin.New()
	useMiddleware(r)
	initDebugRouter(r)
	initSwagRouter(r)
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
