package logger

import (
	"fmt"
	"log"
	"runtime/debug"

	"go.uber.org/zap"

	"github.com/Jancd/gokit/chars"
)

var (
	logger Zapx
	sugar  zap.SugaredLogger
)

func init() {
	l, err := zap.NewProduction()
	defer l.Sync()

	if err != nil {
		log.Fatal(err)
	}
	logger.Logger = l
	sugar = *logger.Sugar()
}

type (
	Logger interface {
		Error(...interface{})
		Errorf(string, ...interface{})
		Info(...interface{})
		Infof(string, ...interface{})
		Slow(...interface{})
		Slowf(string, ...interface{})
	}

	Zapx struct {
		*zap.Logger
	}
)

func Error(v ...interface{}) {
	sugar.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	sugar.Errorf(format, v...)
}

func ErrorStack(v ...interface{}) {
	sugar.Error(stackInfo(fmt.Sprint(v...)))
}

func ErrorStackf(format string, v ...interface{}) {
	sugar.Error(stackInfo(fmt.Sprintf(format, v...)))
}

func Info(v ...interface{}) {
	sugar.Info(v...)
}

func Infof(format string, v ...interface{}) {
	sugar.Infof(format, v...)
}

func Slow(v ...interface{}) {
	sugar.Warn(v...)
}

func Slowf(format string, v ...interface{}) {
	sugar.Warnf(format, v...)
}

func stackInfo(msg string) string {
	return fmt.Sprintf("%s\n%s", msg, chars.BytesToString(debug.Stack()))
}
