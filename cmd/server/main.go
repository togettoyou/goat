package main

import (
	"encoding/json"
	"fmt"
	"os"

	"goat/internal/server"
	"goat/pkg"
	"goat/pkg/conf"
	"goat/pkg/log"
	"goat/pkg/validatorer"
	"goat/pkg/version"

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
