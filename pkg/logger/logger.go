package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerObj *zap.SugaredLogger

func defaultConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel), // 设置日志等级
		Development: false,                               // 非开发模式
		Encoding:    "console",                           // 输出格式（json 或 console）
		OutputPaths: []string{getOutPutFile()},           // 输出目标
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			MessageKey:     "msg",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder, // 日志等级大写
			EncodeTime:     zapcore.ISO8601TimeEncoder,  // 时间格式
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
}

func initLogger() {
	config := defaultConfig()
	logger, err := config.Build(zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCallerSkip(1)) // 创建基础 Logger
	if err != nil {
		panic(err)
	}

	loggerObj = logger.Sugar()
}

func getOutPutFile() string {
	return fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02"))
}

// updateLogger update the output file for LogrusObj
func updateLogger() {
	oneday := int64(3600 * 24)
	for {
		now := time.Now().Unix()
		remain := oneday - (now % oneday)
		initLogger()
		time.Sleep(time.Duration(remain) * time.Second)
	}
}
