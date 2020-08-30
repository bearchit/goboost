package app

import (
	"fmt"

	"github.com/bearchit/goboost/config"
	"github.com/creasty/defaults"
	"github.com/sirupsen/logrus"
)

type App struct {
	Logger logrus.Logger
}

type Config interface {
	GetLogLevel() string
}

type BaseConfig struct {
	LogLevel string `default:"debug"`
}

func New(
	stage string,
	envPrefix string,
	v Config,
) (*App, error) {
	if err := defaults.Set(v); err != nil {
		return nil, err
	}

	c := config.New(
		config.WithScanners(
			config.NewEnvScanner(envPrefix, true),
			config.NewYMLScanner(fmt.Sprintf("config/%s.yml", stage), false),
		),
	)

	if err := c.Unmarshal(v); err != nil {
		return nil, err
	}

	logger := *logrus.New()
	if l, err := logrus.ParseLevel(v.GetLogLevel()); err != nil {
		logger.SetLevel(l)
	} else {
		logger.SetLevel(logrus.DebugLevel)
	}

	return &App{
		Logger: logger,
	}, nil
}
