package log

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	gormlogger "gorm.io/gorm/logger"
)

type Logger interface {
	Configure(level string, format string, filePath string) error

	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	WithError(err error) Logger

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})

	GetTracingLogger() TracingLogger
	GetLockDistributorLogger() LockDistributorLogger
	GetGormLogger() gormlogger.Interface
}

type TracingLogger interface {
	Error(msg string)
	Infof(msg string, args ...interface{})
	Debugf(msg string, args ...interface{})
}

type LockDistributorLogger interface {
	Println(args ...interface{})
}

type GormLogger interface {
	LogMode(gormlogger.LogLevel) GormLogger
	Info(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}

func NewLogger() Logger {
	return &loggerIns{
		FieldLogger: logrus.New(),
	}
}
