package main

import (
	"fmt"

	"github.com/cloudnative-zoo/go-commons/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cli.NewCommand(cli.CommandConfig{
		Use:   "app",
		Short: "A sample CLI app",
		Long:  "This is a sample CLI app demonstrating the usage of the NewCommand function",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("App executed!")
		},
		Flags: []cli.FlagConfig{
			{
				Name:         "name",
				Short:        "n",
				DefaultValue: "World",
				Usage:        "Specify the name",
				Required:     false,
			},
			{
				Name:         "verbose",
				Short:        "v",
				DefaultValue: false,
				Usage:        "Enable verbose mode",
				Required:     false,
			},
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
