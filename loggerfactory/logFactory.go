// package loggerfactory handles creating concrete logger with factory method pattern
package loggerfactory

import (
	"loggo/config"
)

// logger mapp to map logger code to logger builder
var logfactoryBuilderMap = map[string]logFbInterface{
	config.LOGRUS: &LogrusFactory{},
	config.ZAP:    &ZapFactory{},
}

// interface for logger factory
type logFbInterface interface {
	Build(*config.LogConfig) error
}

// accessors for factoryBuilderMap
func GetLogFactoryBuilder(key string) logFbInterface {
	return logfactoryBuilderMap[key]
}
