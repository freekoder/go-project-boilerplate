package logger

import "go.uber.org/zap"

func ConfigureLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
