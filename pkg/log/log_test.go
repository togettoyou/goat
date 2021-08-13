package log

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestZap(t *testing.T) {
	Setup()

	userLog := NewModule("user").L()
	userLog.Info("unknown user err")

	loginLog := userLog.Named("login")
	loginLog.Info("password err")

	level.SetLevel(zapcore.ErrorLevel)

	loginLog.Info("password err")
	userLog.Info("unknown user err")
}
