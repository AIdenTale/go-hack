// Package logger содержит инфраструктурные функции для работы с логированием.
package logger

import (
	"go.uber.org/zap"
)

// New создает zap.Logger с заданным уровнем логирования.
func New(level string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	if level == "debug" {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	return cfg.Build()
}
