package logger

import (
	"os"
	"strings"

	"github.com/WangHongshuo/acfun_comments_observer_backend/cfg"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	atom   zap.AtomicLevel
	core   zapcore.Core
	logger *zap.SugaredLogger
)

func init() {
	atom = zap.NewAtomicLevel()
	atom.SetLevel(convertCfgLevelToZapCoreLevel(cfg.GlobalConfig.Logger.Level))
	core = zapcore.NewCore(newEncoder(), newLogWriter(cfg.GlobalConfig.Logger.OnSave), atom)
	logger = zap.New(core, zap.AddCaller()).Sugar()
}

func NewLogger(name string) *zap.SugaredLogger {
	return logger.Named(name)
}

func WatiLogger() {
	logger.Sync()
}

func newLogWriter(onSave bool) zapcore.WriteSyncer {
	syncConsole := zapcore.AddSync(os.Stderr)
	if !onSave {
		return zapcore.AddSync(syncConsole)
	}
	config := cfg.GlobalConfig.Logger
	rollingFilesCfg := &lumberjack.Logger{
		Filename:   config.Path,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   true,
	}
	return zapcore.NewMultiWriteSyncer(syncConsole, zapcore.AddSync(rollingFilesCfg))
}

func newEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func convertCfgLevelToZapCoreLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.ErrorLevel
	}
}
