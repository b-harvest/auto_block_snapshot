package cmd

import (
	"auto_block_snapshot/pkg/config"
	"auto_block_snapshot/pkg/server"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "auto_block_snapshot",
	Short: "auto_block_snapshot is an automated tool for snapshot and pruning",
	Long:  `auto_block_snapshot is a CLI tool that periodically takes snapshot, prunes and uploads the data to an S3 bucket.`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")

		cfg := config.NewConfig()
		cfg.Load(configFile)

		server := server.NewServer(cfg)
		server.Run()
	},
}

func Execute() error {
	rootCmd.PersistentFlags().StringP("config", "c", "./config.toml", "config file (default is ./config.toml)")
	return rootCmd.Execute()
}
