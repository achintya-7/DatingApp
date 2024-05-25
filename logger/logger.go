package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

func LoadLogger() {
	logger, err := zap.NewDevelopment(
		zap.AddStacktrace(zap.ErrorLevel),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		panic(err)
	}

	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	zap.L().Info("Logger initialized")
}

func Info(ctx context.Context, msg ...any) {
	if ctx == nil {
		zap.L().Info(fmt.Sprintf("%v", msg...))
		return
	}

	correlationID, ok := ctx.Value("correlation_id").(string)
	if !ok {
		zap.L().Info(fmt.Sprintf("%v %s", msg...))
		return
	}

	zap.L().Info(fmt.Sprintf("[Correlation ID: %s] %s", correlationID, fmt.Sprint(msg...)))
}

func Error(ctx context.Context, msg ...any) {
	if ctx == nil {
		zap.L().Error(fmt.Sprintf("%v", msg...))
		return
	}

	correlationID, ok := ctx.Value("correlation_id").(string)
	if !ok {
		zap.L().Error(fmt.Sprintf("%v", msg...))
		return
	}

	zap.L().Error(fmt.Sprintf("[Correlation ID: %s] %s", correlationID, fmt.Sprint(msg...)))
}

func Debug(ctx context.Context, msg ...any) {
	if ctx == nil {
		zap.L().Debug(fmt.Sprintf("%v", msg...))
		return
	}

	correlationID, ok := ctx.Value("correlation_id").(string)
	if !ok {
		zap.L().Debug(fmt.Sprintf("%v", msg...))
		return
	}

	zap.L().Debug(fmt.Sprintf("[Correlation ID: %s] %s", correlationID, fmt.Sprint(msg...)))
}

func Warn(ctx context.Context, msg ...any) {
	if ctx == nil {
		zap.L().Warn(fmt.Sprintf("%v", msg...))
		return
	}

	correlationID, ok := ctx.Value("correlation_id").(string)
	if !ok {
		zap.L().Warn(fmt.Sprintf("%v", msg...))
		return
	}

	zap.L().Warn(fmt.Sprintf("[Correlation ID: %s] %s", correlationID, fmt.Sprint(msg...)))
}

func Fatal(ctx context.Context, msg ...any) {
	if ctx == nil {
		zap.L().Fatal(fmt.Sprintf("%v", msg...))
		return
	}

	correlationID, ok := ctx.Value("correlation_id").(string)
	if !ok {
		zap.L().Fatal(fmt.Sprintf("%v", msg...))
		return
	}

	zap.L().Fatal(fmt.Sprintf("[Correlation ID: %s] %s", correlationID, fmt.Sprint(msg...)))
}
