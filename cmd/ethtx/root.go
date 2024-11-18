/*
Copyright © 2024 Swee Tat Lim <st_lim@stlim.net>
*/
package main

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var addr string

type rootCmd struct {
	cmd *cobra.Command
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func newRootCmd(ctx context.Context) *rootCmd {
	logger := zerolog.Ctx(ctx)

	result := &rootCmd{}
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// cobra.OnInitialize(NewConfig)

	// rootCmd represents the base command when called without any subcommands
	result.cmd = &cobra.Command{
		Use:   "ethtx",
		Short: "An ethereum transaction cli",
		Long:  `An ethereum transaction cli for trustwallet take home test`,
	}
	err := result.cmd.ExecuteContext(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("rootCmd.Execute")
	}
	_, blockCmd := newBlockCmd(ctx)
	_, listCmd := newListCmd(ctx)
	_, subscribeCmd := newSubscribeCmd(ctx)

	result.cmd.AddCommand(
		blockCmd,
		listCmd,
		subscribeCmd,
	)
	return result
}
