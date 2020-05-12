package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level string `required:"true"`
}

func parseLevel(level string) (*zap.AtomicLevel, error) {
	atomicLevel := zap.NewAtomicLevel()
	levelInBytes := []byte(strings.ToLower(level))
	err := atomicLevel.UnmarshalText(levelInBytes)
	if err != nil {
		return nil, err
	}
	return &atomicLevel, nil
}

func NewLogger(config *Config) (*zap.Logger, error) {
	atomicLevel, err := parseLevel(config.Level)
	if err != nil {
		return nil, err
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapConfig := &zap.Config{
		Level:             *atomicLevel,
		DisableStacktrace: true,
		OutputPaths:       []string{"stdout"},
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
