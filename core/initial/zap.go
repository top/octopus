package initial

import (
	"octopus/core/global"
	"octopus/core/helper"
	"os"
	"path"
	"time"

	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var level zapcore.Level

func Zap() *zap.Logger {
	println("initializing zap logger: " + global.CONFIG.System.Zap.Directory)
	if err := helper.InitLogDir(global.CONFIG.System.Zap.Directory); err != nil {
		println("initializing zap logger failed: " + err.Error())
		return nil
	}

	switch global.CONFIG.System.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	var logger *zap.Logger
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if global.CONFIG.System.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func getEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: global.CONFIG.System.Zap.StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(global.CONFIG.System.Zap.Prefix + " " + time.RFC3339))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.CONFIG.System.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.CONFIG.System.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.CONFIG.System.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.CONFIG.System.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

func getWriterSyncer() (zapcore.WriteSyncer, error) {
	writer, err := zaprotatelogs.New(
		path.Join(global.CONFIG.System.Zap.Directory, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(global.CONFIG.System.Zap.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if global.CONFIG.System.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writer)), err
	}
	return zapcore.AddSync(writer), err
}

func getEncoderCore() zapcore.Core {
	writer, err := getWriterSyncer()
	if err != nil {
		println("got write syncer error: ", err)
		return nil
	}
	enc := func(format string) zapcore.Encoder {
		if format == "json" {
			return zapcore.NewJSONEncoder(getEncoderConfig())
		}
		return zapcore.NewConsoleEncoder(getEncoderConfig())
	}(global.CONFIG.System.Zap.Format)
	return zapcore.NewCore(enc, writer, level)
}
