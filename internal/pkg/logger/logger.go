package logger

import (

	"os"
	
	"miniBilling/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)


type LoggerZap struct{
	*zap.Logger
}
func NewLogger(logConfig config.Logger) *LoggerZap{
	logLevel := logConfig.Log_level // debug->info-> warning -> error -> fatal ->panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level =  zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "panic":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEcoderLog()
	hook := lumberjack.Logger{
		Filename:   logConfig.File_log,
		MaxSize:    logConfig.Max_size, // megabytes
		MaxBackups: logConfig.Max_backup,
		MaxAge:     logConfig.Max_age, //days
		Compress:   logConfig.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	if logger == nil {
		panic("Failed to initialize logger")
	}

	return &LoggerZap{logger}
}

func getEcoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}