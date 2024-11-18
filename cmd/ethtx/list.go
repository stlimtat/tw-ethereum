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
	"github.com/stlimtat/tw-ethereum/pkg/ethereum"
)

type listCmd struct {
	cmd    *cobra.Command
	cfg    config.RootConfig
	parser ethereum.IParser
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
			var lcfg config.RootConfig
			err := viper.Unmarshal(&lcfg)
			if err != nil {
				logger.Fatal().Err(err).Msg("viper.Unmarshal")
			}
			result.parser = ethereum.NewDefaultParser(cmd.Context())
			txns, err := result.parser.GetTransactions(cmd.Context(), lcfg.Addr)
			if err != nil {
				logger.Error().Err(err).Msg("result.parser.GetTransactions")
			}
			for idx, txn := range txns {
				logger.Info().
					Int("idx", idx).
					Interface("txn", txn).
					Msg("GetTransactions")
			}
		},
	}
	// result.cfg = config.NewListConfig(ctx, result.cmd)
	// logger.Info().Interface("result.cfg", result.cfg).Msg("NewListConfig")

	return result, result.cmd
}
