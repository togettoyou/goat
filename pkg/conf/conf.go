package conf

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type config struct {
	Server server `mapstructure:"SERVER"`
	Log    log    `mapstructure:"LOG"`
	Redis  redis  `mapstructure:"REDIS"`
	Mysql  mysql  `mapstructure:"MYSQL"`
}

type server struct {
	URL          string `mapstructure:"URL"`
	RunMode      string `mapstructure:"RUNMODE"`
	ReadTimeout  int    `mapstructure:"READTIMEOUT"`
	WriteTimeout int    `mapstructure:"WRITETIMEOUT"`
	HttpPort     int    `mapstructure:"HTTPPORT"`
	TLS          bool   `mapstructure:"TLS"`
	AutoTLS      bool   `mapstructure:"AUTOTLS"`
	ACMEEmail    string `mapstructure:"ACMEEMAIL"`
	Crt          string `mapstructure:"CRT"`
	Key          string `mapstructure:"KEY"`
}

type log struct {
	Level string `mapstructure:"LEVEL"`
}

type redis struct {
	DB       int    `mapstructure:"DB"`
	Addr     string `mapstructure:"ADDR"`
	Password string `mapstructure:"PASSWORD"`
}

type mysql struct {
	Dsn         string `mapstructure:"DSN"`
	MaxIdle     int    `mapstructure:"MAXIDLE"`
	MaxOpen     int    `mapstructure:"MAXOPEN"`
	MaxLifetime int    `mapstructure:"MAXLIFETIME"`
}

var (
	Server server
	Log    log
	Redis  redis
	Mysql  mysql
	Path   string
)

// Setup 配置文件设置
func Setup() {
	if Path != "" {
		viper.SetConfigFile(Path)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("default")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
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
	viper.OnConfigChange(func(in fsnotify.Event) { run() })
	viper.WatchConfig()
}

// setConfig 构造配置文件到结构体对象上
func setConfig() error {
	var config config
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}
	Server = config.Server
	Log = config.Log
	Redis = config.Redis
	Mysql = config.Mysql
	return nil
}
