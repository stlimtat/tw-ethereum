package config

import (
	"context"

	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type ListConfig struct {
	Addr string `mapstructure:"addr"`
}

func NewListConfig(ctx context.Context) ListConfig {
	logger := zerolog.Ctx(ctx)

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
	var result ListConfig
	err = viper.Unmarshal(&result)
	if err != nil {
		logger.Fatal().Err(err).Msg("Unmarshal")
	}

	logger.Debug().
		Interface("result", result).
		Msg("ListConfig init")

	return result
}
