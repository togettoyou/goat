package main

import (
	"encoding/json"
	"fmt"
	"os"

	"goat-layout/internal/server"
	"goat-layout/pkg"
	"goat-layout/pkg/conf"
	"goat-layout/pkg/log"
	"goat-layout/pkg/validatorer"
	"goat-layout/pkg/version"

	"github.com/spf13/pflag"
)

var (
	v        = pflag.BoolP("version", "v", false, "显示版本信息")
	confPath = pflag.StringP("conf", "c", "conf/default.yaml", "指定配置文件路径")
)

func setup() {
	conf.Path = *confPath
	conf.Setup()
	conf.OnChange(func() {
		if err := pkg.Reset(); err != nil {
			return
		}
		server.Reset()
	})
	log.Setup()
	validatorer.Setup()
}

// @title 接口文档
// @version 1.0
// @description 🐐 基于 gin + gorm 的轻量级工程项目
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	pflag.Parse()
	if *v {
		info := version.Get()
		marshalled, err := json.MarshalIndent(&info, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshalled))
		return
	}
	setup()
	server.Start()
}
