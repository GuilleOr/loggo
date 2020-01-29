package loggerfactory

import (
	"github.com/pkg/errors"
	"sa-logging/config"
	"sa-logging/loggerfactory/logrus"
)

// receiver for logrus factory
type LogrusFactory struct{}

// build logrus logger
func (mf *LogrusFactory) Build(lc *config.LogConfig) error {
	err := logrus.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
