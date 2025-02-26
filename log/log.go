package log

import (
	"os"

	"frec.kr/tdoo/config"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var Log *Logger

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

	switch cfg.Format {
	case "json":
		l.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		l.SetFormatter(&logrus.TextFormatter{})
	default:
		l.SetFormatter(&logrus.TextFormatter{})
	}

	return &Logger{l}
}

func Setup(cfg config.LoggerConfig) {
	l := NewLogger(cfg)
	Log = l
}
