// log/logger.go
package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 全局日志实例
var Logger *zap.Logger

// SugaredLogger 全局 sugared logger 实例
var SugaredLogger *zap.SugaredLogger

func InitLogger() error {
	// 1. 创建日志目录
	if err := os.MkdirAll("./logs", 0755); err != nil {
		return err
	}

	// 2. 配置日志级别
	level := zap.NewAtomicLevelAt(zap.InfoLevel)

	// 3. 配置控制台输出（文本格式）
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.TimeKey = "time"
	consoleEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)

	// 4. 配置文件输出（JSON 格式）
	fileWriter := &lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    10, // 单个文件最大 MB
		MaxBackups: 5,  // 保留旧日志数量
		MaxAge:     30, // 保留天数
		Compress:   true,
	}
	fileEncoderConfig := zap.NewProductionEncoderConfig()
	fileEncoderConfig.TimeKey = "time"
	fileEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)
	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(fileWriter), level)

	// 5. 创建多输出 core
	core := zapcore.NewTee(consoleCore, fileCore)

	// 6. 创建 logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	SugaredLogger = Logger.Sugar()

	return nil
}
