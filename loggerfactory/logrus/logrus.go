// package logrus handles creating logrus logger
package logrus

import (
	"github.com/GuilleOr/loggo/config"
	"github.com/GuilleOr/loggo/logger"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

func RegisterLog(lc config.LogConfig) error {

	log := logrus.New()

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		log.Panic(err)
	}

	hook, err := elogrus.NewAsyncElasticHook(client, "localhost", logrus.DebugLevel, "app_kyru_log")
	if err != nil {
		log.Panic(err)
	}


	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)
	//log.SetOutput(os.Stdout)
	//customize it from configuration file
	err = customizeLogFromConfig(log, lc)
	if err != nil {
		return errors.Wrap(err, "")
	}

	if err == nil {
		log.AddHook(hook)
	} else {
		logrus.Errorf(" error: %q", err)
	}

	logger.SetLogger(log)
	return nil
}

// customizeLogFromConfig customize log based on parameters from configuration file
func customizeLogFromConfig(log *logrus.Logger, lc config.LogConfig) error {
	log.SetReportCaller(lc.EnableCaller)
	//log.SetOutput(os.Stdout)
	l := &log.Level
	err := l.UnmarshalText([]byte(lc.Level))
	if err != nil {
		return errors.Wrap(err, "")
	}
	log.SetLevel(*l)
	return nil
}
