package log

import (
	"context"
)

type loggerKey struct{}

// NewContext context with tags logger
func NewContext(ctx context.Context, tags ...LogField) context.Context {
	return context.WithValue(ctx, loggerKey{}, global.With(tags...))
}

// NewContextWithLogger context with tags logger
func NewContextWithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// NewOrFromContext context with tags logger
func NewOrFromContext(ctx context.Context, logger Logger) context.Context {
	if _, ok := ctx.Value(loggerKey{}).(Logger); ok {
		return ctx
	}
	return context.WithValue(ctx, loggerKey{}, logger)
}

// Inject add tags to current context
func Inject(ctx context.Context, tags ...LogField) context.Context {
	if ctxLogger, ok := ctx.Value(loggerKey{}).(Logger); ok {
		ctxLogger = ctxLogger.With(tags...)
		ctx = context.WithValue(ctx, loggerKey{}, ctxLogger)
	}
	return ctx
}

// Extract logger from context
func Extract(ctx context.Context) Logger {
	if ctx == nil {
		return global
	}
	if ctxLogger, ok := ctx.Value(loggerKey{}).(Logger); ok {
		return ctxLogger
	}
	return global
}
