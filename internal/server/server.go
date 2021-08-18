package server

import (
	"fmt"
	"net/http"

	"goat-layout/internal/dao"
	"goat-layout/internal/server/router"
	"goat-layout/pkg/conf"

	"github.com/gin-gonic/gin"
)

var (
	server  *http.Server
	ginMode = map[string]string{
		"debug":   gin.DebugMode,
		"release": gin.ReleaseMode,
		"test":    gin.TestMode,
	}
)

func setGinMode() {
	if mode, ok := ginMode[conf.Server.RunMode]; ok {
		gin.SetMode(mode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func Start() {
	// 选择数据源实现
	store, err := dao.NewMock()
	//store, err := dao.NewMysql()
	if err != nil {
		panic(err)
	}
	setGinMode()
	httpPort := fmt.Sprintf(":%d", conf.Server.HttpPort)
	server = &http.Server{
		Addr:           httpPort,
		Handler:        router.New(store),
		ReadTimeout:    conf.Server.ReadTimeout,
		WriteTimeout:   conf.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if conf.Server.TLS {
		if err := server.ListenAndServeTLS(conf.Server.Crt, conf.Server.Key); err != nil {
			panic(err)
		}
	} else {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}
}

func Reset() {
	setGinMode()
	server.ReadTimeout = conf.Server.ReadTimeout
	server.WriteTimeout = conf.Server.WriteTimeout
}
