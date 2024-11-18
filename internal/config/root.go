package config

import (
	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type RootConfig struct {
	Addr string `mapstructure:"addr"`
}

func NewRootConfig() {
	logger := log.Logger
	var err error

	home, err := homedir.Dir()
	if err != nil {
		logger.Fatal().Err(err).Msg("homedir.Dir")
	}

	viper.AddConfigPath(home)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("ETX")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("ReadInConfig")
	}

	var result RootConfig
	err = viper.Unmarshal(&result)
	if err != nil {
		logger.Fatal().Err(err).Msg("Unmarshal")
	}
	logger.Info().
		Interface("result", result).
		Interface("viper", viper.AllSettings()).
		Msg("NewRootConfig")
}
