/*
Copyright © 2024 Swee Tat Lim <st_lim@stlim.net>
*/
package main

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stlimtat/tw-ethereum/internal/config"
)

type rootCmd struct {
	cfg config.RootConfig
	cmd *cobra.Command
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func newRootCmd(ctx context.Context) *rootCmd {
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msg("newRootCmd")

	result := &rootCmd{}

	cobra.OnInitialize(config.NewRootConfig)
	// rootCmd represents the base command when called without any subcommands
	result.cmd = &cobra.Command{
		Use:   "ethtx",
		Short: "An ethereum transaction cli",
		Long:  `An ethereum transaction cli for trustwallet take home test`,
	}
	result.cmd.PersistentFlags().StringP("addr", "a", "", "Address of the customer")
	err := viper.BindPFlag("addr", result.cmd.PersistentFlags().Lookup("addr"))
	if err != nil {
		logger.Fatal().Err(err).Msg("viper.BindPFlag.addr")
	}
	_, blockCmd := newBlockCmd(ctx)
	_, listCmd := newListCmd(ctx)
	// _, subscribeCmd := newSubscribeCmd(ctx)

	result.cmd.AddCommand(
		blockCmd,
		listCmd,
		// 	subscribeCmd,
	)
	return result
}
