package config

import (
	"github.com/pkg/errors"
)

// constant for logger code, it needs to match log code (logConfig)in configuration
const (
	LOGRUS string = "logrus"
	ZAP    string = "zap"
)


func validateConfig(appConfig AppConfig) error {
	err := validateLogger(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateLogger(appConfig AppConfig) error {
	zc := appConfig.ZapConfig
	key := zc.Code
	zcMsg := " in validateLogger doesn't match key = "
	if ZAP != key {
		errMsg := ZAP + zcMsg + key
		return errors.New(errMsg)
	}
	lc := appConfig.LogrusConfig
	key = lc.Code
	if LOGRUS != lc.Code {
		errMsg := LOGRUS + zcMsg + key
		return errors.New(errMsg)
	}
	return nil
}

