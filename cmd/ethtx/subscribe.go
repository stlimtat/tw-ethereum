/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/stlimtat/tw-ethereum/internal/config"
)

type subscribeCmd struct {
	cmd *cobra.Command
	cfg config.RootConfig
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
			// 1. maintain a list of address vs id
			// 2. run eth_newFilter
			// 3. run eth_getFilterLogs
		},
	}

	return result, result.cmd
}
