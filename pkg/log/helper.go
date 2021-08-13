package log

import (
	"go.uber.org/zap"
)

type module struct {
	name string
}

func NewModule(name string) *module {
	return &module{
		name: "[MODULE]" + name,
	}
}

func (m *module) L() *zap.Logger {
	return zap.L().Named(m.name)
}

func (m *module) S() *zap.SugaredLogger {
	return zap.S().Named(m.name)
}
