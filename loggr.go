package loggr

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type marker struct{}

var (
	markerKey = &marker{}
)

func New(fields ...interface{}) *zap.SugaredLogger {
	cfgEncoder := zap.NewProductionEncoderConfig()
	cfgProd := zap.NewProductionConfig()

	cfgEncoder.TimeKey = "ts"
	cfgEncoder.MessageKey = "message"
	cfgEncoder.LevelKey = "level"
	cfgEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

	cfgProd.EncoderConfig = cfgEncoder
	lg, _ := cfgProd.Build()
	if lg != nil {
		defer func() {
			_ = lg.Sync()
		}()
	}
	return lg.Sugar().With(fields...)
}

func WithContext(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		return New()
	}
	if ctxLogger, ok := ctx.Value(markerKey).(*zap.SugaredLogger); ok {
		return ctxLogger
	}
	return New()
}

func ToContext(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, markerKey, logger)
}

func Marshal(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
