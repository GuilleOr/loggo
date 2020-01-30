// Package config reasd configurations from a YAML file and load them into a AppConfig type to save the configuration
// information for the application.
// Configuration for different environment can be saved in files with different suffix, for example [Dev], [Prod]
package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// AppConfig represents the application config
type AppConfig struct {
	ZapConfig       LogConfig       `yaml:"zapConfig"`
	LogrusConfig    LogConfig      `yaml:"logrusConfig"`
	Log             LogConfig       `yaml:"logConfig"`
}


// LogConfig represents logger handler
// Logger has many parameters can be set or changed. Currently, only three are listed here. Can add more into it to
// fits your needs.
type LogConfig struct {
	// log library name
	Code string `yaml:"code"`
	// log level
	Level string `yaml:"level"`
	// show caller in log message
	EnableCaller bool `yaml:"enableCaller"`
}

// ReadConfig reads the file of the filename (in the same folder) and put it into the AppConfig
func ReadConfig(filename string) (*AppConfig, error) {

	var ac AppConfig
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read error")
	}
	err = yaml.Unmarshal(file, &ac)

	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}
	/*
	err = validateConfig(ac)
	if err != nil {
		return nil, errors.Wrap(err, "validate config")
	}
	fmt.Println("appConfig:", ac)
	 */
	return &ac, nil
}
