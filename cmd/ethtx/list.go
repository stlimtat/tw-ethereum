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

type listCmd struct {
	cmd *cobra.Command
	cfg config.ListConfig
}

func newListCmd(ctx context.Context) (*listCmd, *cobra.Command) {
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msg("Testing")

	result := &listCmd{}

	// serverCmd represents the server command
	result.cmd = &cobra.Command{
		Use:   "list",
		Short: "Lists transactions for an ethereum account",
		Long:  `Lists the transactions for an ethereum user`,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug().Msg("list.run")
		},
	}
	result.cmd.PersistentFlags().StringP("addr", "a", "", "Address of the customer")
	err := viper.BindPFlag("addr", result.cmd.PersistentFlags().Lookup("addr"))
	if err != nil {
		logger.Fatal().Err(err).Msg("viper.BindPFlag.addr")
	}
	result.cfg = config.NewListConfig(ctx)

	return result, result.cmd
}
