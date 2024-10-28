package logger

import "go.uber.org/zap"

func init() {
	initLogger()
	go updateLogger()
}

func Fatalf(template string, args ...interface{}) {
	loggerObj.Fatal(template, args)
}

func Errorf(template string, args ...interface{}) {
	loggerObj.Errorf(template, args...)
}

func Infof(template string, args ...interface{}) {
	loggerObj.Infof(template, args...)
}

func Debugf(template string, args ...interface{}) {
	loggerObj.Debugf(template, args...)
}

func Fatal(args ...interface{}) {
	loggerObj.Fatal(args)
}

func Info(args ...interface{}) {
	loggerObj.Info(args)
}

func Error(args ...interface{}) {
	loggerObj.Error(args)
}

func Debug(args ...interface{}) {
	loggerObj.Debug(args)
}

var skip = zap.AddCallerSkip(3)

func SDebugf(template string, args ...interface{}) {
	loggerObj.WithOptions(skip).Debugf(template, args)
}

func SDebug(args ...interface{}) {
	loggerObj.WithOptions(skip).Debug(args)
}

func SErrorf(template string, args ...interface{}) {
	loggerObj.WithOptions(skip).Errorf(template, args)
}
func SError(args ...interface{}) {
	loggerObj.WithOptions(skip).Error(args)
}
