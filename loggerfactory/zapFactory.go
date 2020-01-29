package loggerfactory

import (
	"github.com/GuilleOr/loggo/config"
	"github.com/GuilleOr/loggo/loggerfactory/zap"
	"github.com/pkg/errors"
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
