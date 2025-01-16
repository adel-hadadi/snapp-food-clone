package logger

import (
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const logFileNameTemplate = "/tmp/logs/log-%Y-%m-%d.log"

func NewLogger() *zap.Logger {
	rotatort, err := rotatelogs.New(
		logFileNameTemplate,
		rotatelogs.WithMaxAge(60*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(fmt.Errorf("error on creating a new rotator: %w", err))
	}

	w := zapcore.AddSync(rotatort)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	return zap.New(core, zap.AddStacktrace(zap.ErrorLevel))
}
