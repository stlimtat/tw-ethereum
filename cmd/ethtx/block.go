/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

type blockCmd struct {
	cmd *cobra.Command
}

func newBlockCmd(ctx context.Context) (*blockCmd, *cobra.Command) {
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msg("Testing")

	result := &blockCmd{}

	// serverCmd represents the server command
	result.cmd = &cobra.Command{
		Use:   "block",
		Short: "Get the current block number for ethereum",
		Long:  `Get the current block number for ethereum`,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug().Msg("block.run")
		},
	}

	return result, result.cmd
}
