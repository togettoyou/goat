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
	zap.ReplaceGlobals(zap.New(
		zapcore.NewTee(zapcore.NewCore(
			getConsoleEncoder(),
			zapcore.Lock(os.Stdout),
			level,
		)),
	))
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
