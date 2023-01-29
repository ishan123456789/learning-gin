package utils

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	Log, _ = zap.NewProduction()
	defer Log.Sync()
}

func SyncBeforeDying() {
	defer Log.Sync()
}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

func StringField(key string, message string) zapcore.Field {
	return zap.String(key, message)
}

func LogError(where string, err error) {
	Log.Error(fmt.Sprintf("Error occured at:: %s", where), zap.Error(err))
}
