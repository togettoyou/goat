package log

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestZap(t *testing.T) {
	Setup()

	userLog := NewModule("user").L()
	userLog.Info("hello")

	loginLog := userLog.Named("login")
	loginLog.Info("login success")

	level.SetLevel(zapcore.ErrorLevel)

	loginLog.Error("password err")
	userLog.Info("unknown user err")
}
