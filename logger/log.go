package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	log, err := config.Build()
	if err != nil {
		zap.S().Errorf("failed to create logger: %v", err)
		os.Exit(1)
	}
	defer log.Sync()
	zap.ReplaceGlobals(log)
}
