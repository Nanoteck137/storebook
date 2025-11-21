package main

import (
	"github.com/nanoteck137/pyrin/spark"
	"github.com/nanoteck137/pyrin/spark/typescript"
	"github.com/nanoteck137/pyrin/trail"
	"github.com/nanoteck137/storebook/apis"
	"github.com/spf13/cobra"
)

var logger = trail.NewLogger(&trail.Options{
	Debug: true,
})

var rootCmd = &cobra.Command{
	Use: "internal",
}

var genCmd = &cobra.Command{
	Use: "gen",
	Run: func(cmd *cobra.Command, args []string) {
		router := spark.Router{}
		apis.RegisterHandlers(nil, &router)

		nameFilter := spark.NameFilter{}

		serverDef, err := spark.CreateServerDef(&router, nameFilter)
		if err != nil {
			logger.Fatal("failed to create server def", "err", err)
		}

		err = serverDef.SaveToFile("misc/pyrin.json")
		if err != nil {
			logger.Fatal("failed save server def", "err", err)
		}

		logger.Info("Wrote 'misc/pyrin.json'")

		resolver, err := spark.CreateResolverFromServerDef(&serverDef)
		if err != nil {
			logger.Fatal("failed to create resolver", "err", err)
		}

		{
			gen := typescript.TypescriptGenerator{}

			err = gen.Generate(&serverDef, resolver, "web/src/lib/api")
			if err != nil {
				logger.Fatal("failed to generate typescript client", "err", err)
			}
		}

		// {
		// 	gen := golang.GolangGenerator{}
		//
		// 	err = gen.Generate(&serverDef, resolver, "cmd/storebook-cli/api")
		// 	if err != nil {
		// 		logger.Fatal("failed to generate golang client", "err", err)
		// 	}
		// }
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		logger.Fatal("failed to execute", "err", err)
	}
}
