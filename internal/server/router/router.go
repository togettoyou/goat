package router

import (
	"goat/internal/server/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(gin.Logger())
	// 开启性能分析
	// 可以根据需要使用 pprof.RouteRegister() 控制访问权限
	pprof.Register(r)
	return r
}
