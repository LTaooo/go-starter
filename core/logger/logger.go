package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 全局日志实例
var Logger *zap.Logger

// Config 日志配置
type Config struct {
	Level         string   `json:"level" yaml:"level"`                 // 日志级别: debug, info, warn, error, fatal
	Format        string   `json:"format" yaml:"format"`               // 日志格式: json, console
	OutputPaths   []string `json:"outputPaths" yaml:"outputPaths"`     // 输出路径列表
	ErrorOutputs  []string `json:"errorOutputs" yaml:"errorOutputs"`   // 错误输出路径列表
	Filename      string   `json:"filename" yaml:"filename"`           // 日志文件名
	MaxSize       int      `json:"maxSize" yaml:"maxSize"`             // 单个文件最大大小 (MB)
	MaxBackups    int      `json:"maxBackups" yaml:"maxBackups"`       // 保留旧文件的最大数量
	MaxAge        int      `json:"maxAge" yaml:"maxAge"`               // 保留旧文件的最大天数
	Compress      bool     `json:"compress" yaml:"compress"`           // 是否压缩旧文件
	Development   bool     `json:"development" yaml:"development"`     // 是否为开发环境
	DisableCaller bool     `json:"disableCaller" yaml:"disableCaller"` // 是否禁用调用者信息
	DisableStack  bool     `json:"disableStack" yaml:"disableStack"`   // 是否禁用堆栈跟踪
}

func Get() *zap.Logger {
	return Logger
}

func NewConfig() Config {
	return Config{
		Level:         "info",
		Format:        "console",
		OutputPaths:   []string{"stdout"},
		ErrorOutputs:  []string{"stderr"},
		Filename:      "logs/app.log",
		MaxSize:       10,
		MaxBackups:    5,
		MaxAge:        30,
		Compress:      true,
		Development:   false,
		DisableCaller: false,
		DisableStack:  false,
	}
}

// Init 初始化日志
func Init(config Config) error {

	// 解析日志级别
	var level zapcore.Level
	err := level.UnmarshalText([]byte(config.Level))
	if err != nil {
		level = zapcore.InfoLevel
	}

	// 创建多个 Core
	var cores []zapcore.Core

	// 控制台输出（彩色）
	if contains(config.OutputPaths, "stdout") {
		consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     customTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})
		cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level))
	}

	// 文件输出（JSON）
	if contains(config.OutputPaths, "file") || contains(config.OutputPaths, config.Filename) {
		fileEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     customTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})

		// 使用 lumberjack 实现日志轮转
		fileWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.Filename,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		})

		cores = append(cores, zapcore.NewCore(fileEncoder, fileWriter, level))
	}

	// 创建 Tee Core
	core := zapcore.NewTee(cores...)

	// 构建 Logger
	loggerOptions := []zap.Option{}
	if !config.DisableCaller {
		loggerOptions = append(loggerOptions, zap.AddCaller())
	}
	if !config.DisableStack {
		loggerOptions = append(loggerOptions, zap.AddStacktrace(zap.ErrorLevel))
	}
	if config.Development {
		loggerOptions = append(loggerOptions, zap.Development())
	}

	Logger = zap.New(core, loggerOptions...)
	zap.RedirectStdLog(Logger)

	return nil
}

// Sync 刷新日志缓冲区
func Sync() {
	Logger.Sync()
}

// 自定义时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05.000Z0700"))
}

// 判断切片是否包含元素
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
