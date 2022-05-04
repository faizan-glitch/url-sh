package main

import (
	"log"

	"github.com/faizan-glitch/url-sh/cmd/cli/commands"
	"github.com/faizan-glitch/url-sh/pkg/config"
	"github.com/faizan-glitch/url-sh/pkg/database"
	"github.com/spf13/cobra"
)

var (
	RootCmd = cobra.Command(cobra.Command{
		Use:   "url-sh",
		Short: "url-sh is a command line tool to shorten and resolve urls",
		Long: `A Simple and Blazingly Fast URL Shortner and Resolver built with love by faizan-glitch in Go.
Complete source code is available at https://github.com/faizan-glitch/url-sh/`,
		DisableAutoGenTag: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	})
)

func init() {
	config.Load()

	database.Connect()

	RootCmd.AddCommand(commands.GetShortenCommand())

	RootCmd.AddCommand(commands.GetResolveCommand())
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
