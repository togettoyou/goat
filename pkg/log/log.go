package log

import (
	"os"

	"goat/pkg/conf"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var level = zap.NewAtomicLevel()

// Setup 日志设置
func Setup() {
	setLevel()
	cores := make([]zapcore.Core, 0)
	// 创建控制台Core记录器接口
	cores = append(cores, zapcore.NewCore(
		getConsoleEncoder(),
		zapcore.Lock(os.Stdout),
		level,
	))
	// Option可选配置
	options := []zap.Option{
		zap.AddCaller(),                   // Caller调用显示
		zap.AddStacktrace(zap.ErrorLevel), // 堆栈跟踪级别
	}
	// 构建一个 zap 实例
	logger := zap.New(
		zapcore.NewTee(cores...),
		options...,
	)
	zap.ReplaceGlobals(logger) // 设为全局zap实例
}

var levels = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

// setLevel 日志级别设置
func setLevel() {
	if l, ok := levels[conf.Log.Level]; ok {
		level.SetLevel(l)
	} else {
		level.SetLevel(zapcore.DebugLevel)
	}
}

// getConsoleEncoder 控制台日志格式
func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func Reset() {
	setLevel()
}
