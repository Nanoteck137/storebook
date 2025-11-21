package cmd

import (
	"github.com/nanoteck137/storebook"
	"github.com/nanoteck137/storebook/config"
	"github.com/nanoteck137/storebook/core"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
}

var upCmd = &cobra.Command{
	Use: "up",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.LoadedConfig
		conf.RunMigrations = false
		app := core.NewBaseApp(&conf)

		err := app.Bootstrap()
		if err != nil {
			app.Logger().Fatal("Failed to bootstrap app", "err", err)
		}

		// err = app.DB().RunMigrateUp()
		// if err != nil {
		// 	log.Fatal("Failed to run migrate up", "err", err)
		// }
	},
}

var downCmd = &cobra.Command{
	Use: "down",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.LoadedConfig
		conf.RunMigrations = false
		app := core.NewBaseApp(&conf)

		err := app.Bootstrap()
		if err != nil {
			app.Logger().Fatal("Failed to bootstrap app", "err", err)
		}

		// err = app.DB().RunMigrateDown()
		// if err != nil {
		// 	log.Fatal("Failed to run migrate down", "err", err)
		// }
	},
}

// TODO(patrik): Move to dev cmd
var createCmd = &cobra.Command{
	Use:  "create <MIGRATION_NAME>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		logger := storebook.DefaultLogger()

		err := goose.Create(nil, "./migrations", name, "sql")
		if err != nil {
			logger.Fatal("Failed to create migration", "err", err)
		}
	},
}

// TODO(patrik): Move to dev cmd?
var fixCmd = &cobra.Command{
	Use: "fix",
	Run: func(cmd *cobra.Command, args []string) {
		logger := storebook.DefaultLogger()

		err := goose.Fix("./migrations")
		if err != nil {
			logger.Fatal("Failed to fix migrations", "err", err)
		}
	},
}

func init() {
	migrateCmd.AddCommand(upCmd)
	migrateCmd.AddCommand(downCmd)
	migrateCmd.AddCommand(createCmd)
	migrateCmd.AddCommand(fixCmd)

	rootCmd.AddCommand(migrateCmd)
}
