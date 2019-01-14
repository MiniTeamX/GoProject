package main

import (
	"oc_res_control/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var zapCfgLoggerLevel zap.AtomicLevel = zap.NewAtomicLevel()

func  Build(cfg *conf.LoggerConfig) (*zap.Logger, error) {
	zapCfg := zap.NewProductionConfig()
	zapCfg.Level = cfg.Level
	zapCfg.OutputPaths = cfg.OutputPaths
	localLoc, err := time.LoadLocation("Asian/Shanghai")
	if err != nil {
		localLoc = time.Now().Location()
	}
	zapCfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.In(localLoc).Format("2006-01-02 15:04:05") + "]")
	}
	return zapCfg.Build()
}

func SetLogger(log *zap.Logger) func() {
	undoGlobals := zap.ReplaceGlobals(log)
	undoHijack := zap.RedirectStdLog(log)
	return func() {
		undoGlobals()
		undoHijack()
	}
}

func BuildLogger(cfg *conf.LoggerConfig) error {
	logger, err := Build(cfg)
	if err != nil {
		return err
	}
	zapCfgLoggerLevel = cfg.Level
	SetLogger(logger)
	return nil
}

func enabledLogDebug() bool {
	return zapCfgLoggerLevel.Enabled(zap.DebugLevel)
}

func enabledLogInfo() bool {
	return zapCfgLoggerLevel.Enabled(zap.InfoLevel)
}

func enabledLogWarn() bool {
	return zapCfgLoggerLevel.Enabled(zap.WarnLevel)
}

func enabledLogError() bool {
	return zapCfgLoggerLevel.Enabled(zap.ErrorLevel)
}

func enabledLogDPanic() bool {
	return zapCfgLoggerLevel.Enabled(zap.DPanicLevel)
}

func enabledLogPanic() bool {
	return zapCfgLoggerLevel.Enabled(zap.PanicLevel)
}

func enabledLogFatal() bool {
	return zapCfgLoggerLevel.Enabled(zap.FatalLevel)
}




