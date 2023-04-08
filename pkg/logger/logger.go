package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.EpochTimeEncoder

	log, _ := config.Build()

	Logger = log.Sugar()
}

func Info(msg string, keys ...interface{}) {
	Logger.Infow(msg, keys...)
}

func Warn(msg string, keys ...interface{}) {
	Logger.Infow(msg, keys...)
}

func Fatal(msg string, keys ...interface{}) {
	Logger.Fatalw(msg, keys...)
}

func Panic(msg string, keys ...interface{}) {
	Logger.Panicw(msg, keys...)
}
