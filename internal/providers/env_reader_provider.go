package providers

import (
	"fmt"
	"github.com/spf13/viper"
)

type EnvReaderProvider struct {
}

func NewEnvReaderProvider() *EnvReaderProvider {
	return &EnvReaderProvider{}
}

func (envReaderProvider *EnvReaderProvider) InitViper() {
	const (
		configType = "env"
		configFile = ".env"
	)

	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
