package loggo

import (
	"github.com/GuilleOr/loggo/loggerfactory"
	"github.com/pkg/errors"
	"loggo/config"
	loggo "loggo/logger"
)

func NewLogo() loggo.Logger {
	err := initApp()
	if err != nil {
		println("init error")
		return nil
	}
	return loggo.Log
}

// TODO load external config file
func initApp() error {
	var err error
	appConfig, err := loadConfig("config/config.yml")
	if err != nil {
		return errors.Wrap(err, "loadConfig")
	}
	err = loadLogger(appConfig.LogrusConfig)
	if err != nil {
		return errors.Wrap(err, "loadLogger")
	}

	return nil
}

// loads the application configurations
func loadConfig(filename string) (*config.AppConfig, error) {

	ac, err := config.ReadConfig(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read container")
	}
	return ac, nil
}

// loads the logger
func loadLogger(lc config.LogConfig) error {
	loggerType := lc.Code
	err := loggerfactory.GetLogFactoryBuilder(loggerType).Build(&lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
