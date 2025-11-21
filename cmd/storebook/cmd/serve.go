package cmd

import (
	"github.com/nanoteck137/storebook/apis"
	"github.com/nanoteck137/storebook/config"
	"github.com/nanoteck137/storebook/core"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		app := core.NewBaseApp(&config.LoadedConfig)

		err := app.Bootstrap()
		if err != nil {
			app.Logger().Fatal("Failed to bootstrap app", "err", err)
		}

		e, err := apis.Server(app)
		if err != nil {
			app.Logger().Fatal("Failed to create server", "err", err)
		}

		err = e.Start(app.Config().ListenAddr)
		if err != nil {
			app.Logger().Fatal("Failed to start server", "err", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
