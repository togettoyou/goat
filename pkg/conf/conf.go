package conf

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type config struct {
	App    app    `yaml:"app"`
	Server server `yaml:"server"`
	Log    log    `yaml:"log"`
	Mysql  mysql  `yaml:"mysql"`
}

type app struct {
	JwtSecret string `yaml:"jwtSecret"`
}

type server struct {
	RunMode      string        `yaml:"runMode"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	HttpPort     int           `yaml:"httpPort"`
	TLS          bool          `yaml:"tls"`
	Crt          string        `yaml:"crt"`
	Key          string        `yaml:"key"`
}

type log struct {
	Level string `yaml:"level"`
}

type mysql struct {
	Dsn         string        `yaml:"dsn"`
	MaxIdle     int           `yaml:"maxIdle"`
	MaxOpen     int           `yaml:"maxOpen"`
	MaxLifetime time.Duration `yaml:"maxLifetime"`
}

var (
	App    app
	Server server
	Log    log
	Mysql  mysql
	Path   string
	v      *viper.Viper
)

// Setup 配置文件设置
func Setup() {
	v = viper.New()
	v.SetConfigFile(Path)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := setConfig(); err != nil {
		panic(err)
	}
}

// Reset 配置文件重设
func Reset() error {
	return setConfig()
}

// OnChange 配置文件热加载回调
func OnChange(run func()) {
	v.OnConfigChange(func(in fsnotify.Event) { run() })
	v.WatchConfig()
}

// setConfig 构造配置文件到结构体对象上
func setConfig() error {
	var config config
	if err := v.Unmarshal(&config); err != nil {
		return err
	}
	App = config.App
	Server = config.Server
	Log = config.Log
	Mysql = config.Mysql
	Server.ReadTimeout *= time.Second
	Server.WriteTimeout *= time.Second
	Mysql.MaxLifetime *= time.Minute
	return nil
}
