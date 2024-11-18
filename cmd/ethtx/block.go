/*
Copyright © 2024 Swee Tat Lim <st_lim@stlim.net>
*/
package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/stlimtat/tw-ethereum/pkg/ethereum"
)

type blockCmd struct {
	cmd    *cobra.Command
	parser ethereum.IParser
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
		Run: func(cmd *cobra.Command, _ []string) {
			logger.Debug().Msg("block.run")
			result.parser = ethereum.NewDefaultParser(cmd.Context())
			blockResult := result.parser.GetCurrentBlock(ctx)
			fmt.Printf("blockResult = (%d)", blockResult)
		},
	}

	return result, result.cmd
}
