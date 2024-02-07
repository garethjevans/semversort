package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "semver",
		Short: "handle carvel based semver correctly",
		Long:  `a CLI that can correctly manipulate the semantic versions that carvel uses`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(getCommand)
	rootCmd.AddCommand(bumpCommand)
	rootCmd.AddCommand(sortCommand)
	rootCmd.AddCommand(rewriteCommand)
}
