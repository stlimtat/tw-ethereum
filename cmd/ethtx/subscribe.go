/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stlimtat/tw-ethereum/internal/config"
)

type subscribeCmd struct {
	cmd *cobra.Command
	cfg config.SubscribeConfig
}

func newSubscribeCmd(ctx context.Context) (*subscribeCmd, *cobra.Command) {
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msg("Testing")

	result := &subscribeCmd{}

	// serverCmd represents the server command
	result.cmd = &cobra.Command{
		Use:   "subscribe",
		Short: "Subscribe to transaction updates for ethereum",
		Long:  `Subscribe to transaction updates for ethereum`,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug().Msg("subscribe.run")
		},
	}
	result.cmd.PersistentFlags().StringP("addr", "a", "", "Address of the customer")
	err := viper.BindPFlag("addr", result.cmd.PersistentFlags().Lookup("addr"))
	if err != nil {
		logger.Fatal().Err(err).Msg("viper.BindPFlag.addr")
	}
	result.cfg = config.NewSubscribeConfig(ctx)

	return result, result.cmd
}
