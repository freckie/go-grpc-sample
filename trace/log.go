package trace

import (
	"context"
	"io"
	"os"

	"frec.kr/tdoo/config"
	"frec.kr/tdoo/internal/ctxkey"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(cfg config.LoggerConfig) *Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)

	switch cfg.Level {
	case "debug":
		l.SetLevel(logrus.DebugLevel)
	case "info":
		l.SetLevel(logrus.InfoLevel)
	case "warn":
		l.SetLevel(logrus.WarnLevel)
	case "error":
		l.SetLevel(logrus.ErrorLevel)
	default:
		l.SetLevel(logrus.InfoLevel)
	}

	txt := logrus.TextFormatter{
		DisableSorting: true,
		PadLevelText:   false,
	}
	switch cfg.Format {
	case "json":
		l.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		l.SetFormatter(&txt)
	default:
		l.SetFormatter(&txt)
	}

	return &Logger{l}
}

func WithLogger(ctx context.Context, l *Logger) context.Context {
	return ctxkey.WithKey[*Logger](ctx, l)
}

func LoggerFrom(ctx context.Context) *Logger {
	l := ctxkey.From[*Logger](ctx)
	if l == nil {
		return discard()
	}
	return l
}

func discard() *Logger {
	l := logrus.New()
	l.SetOutput(io.Discard)
	return &Logger{l}
}
