package zap

import (
	"path/filepath"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger() *zap.Logger  {
	config := zap.NewProductionConfig()
	_, b, _, _ := runtime.Caller(0)
	logPath   := filepath.Join(filepath.Dir(b), "../..", "/logs/development.log") 
	config.OutputPaths = []string{logPath,"stdout"}
	// 環境変数で管理出来たほうが良いかも
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	logger, err := config.Build()
	if err != nil {
		logger.Error("logger build error.")
	}

	return logger
}
