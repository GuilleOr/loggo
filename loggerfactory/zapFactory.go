package loggerfactory

import (
	"github.com/pkg/errors"
	"sa-logging/config"
	"sa-logging/loggerfactory/zap"
)

// receiver for zap factory
type ZapFactory struct{}

// build zap logger
func (mf *ZapFactory) Build(lc *config.LogConfig) error {
	err := zap.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
