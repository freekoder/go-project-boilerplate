package logger

import (
	"github.com/freekoder/go-project-boilerplate/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Configure(cfg config.LoggingConfig) (*zap.Logger, error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	logFile, _ := os.OpenFile(cfg.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}
	core := zapcore.NewTee(zapcore.NewCore(fileEncoder, writer, defaultLogLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger, nil
}
