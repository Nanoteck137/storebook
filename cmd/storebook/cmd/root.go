package cmd

import (
	"log/slog"

	"github.com/nanoteck137/pyrin/trail"
	"github.com/nanoteck137/storebook"
	"github.com/nanoteck137/storebook/config"
	"github.com/spf13/cobra"
)

var logger = trail.NewLogger(&trail.Options{Debug: true, Level: slog.LevelInfo})

var rootCmd = &cobra.Command{
	Use:     storebook.AppName,
	Version: storebook.Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal("Failed to run root command", "err", err)
	}
}

func init() {
	rootCmd.SetVersionTemplate(storebook.VersionTemplate(storebook.AppName))

	cobra.OnInitialize(config.InitConfig)

	rootCmd.PersistentFlags().StringVarP(&config.ConfigFile, "config", "c", "", "Config File")
}
