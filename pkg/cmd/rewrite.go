package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rewriteCommand = &cobra.Command{
		Use:   "rewrite",
		Short: "rewrite a carvel compatible semantic version",
		Long:  `rewrite a carvel compatible semantic version`,
		RunE:  rewrite,
		//Args:  require.MaxArgs(2),
	}
)

func rewrite(cmd *cobra.Command, args []string) error {
	in := ReadFromArgsOrStdin(args[1:], cmd.InOrStdin())

	fmt.Println(in[0])

	return nil
}
