package cmd

import (
	"fmt"
	"github.com/carvel-dev/semver/v4"
	"github.com/spf13/cobra"
)

var (
	getCommand = &cobra.Command{
		Use:   "get",
		Short: "get a part of a carvel compatible semantic version",
		Long:  `get a part of a carvel compatible semantic version`,
		RunE:  get,
		//Args:  require.MaxArgs(2),
	}
)

func get(cmd *cobra.Command, args []string) error {
	part := args[0]
	in := ReadFromArgsOrStdin(args[1:], cmd.InOrStdin())

	v := semver.MustParse(in[0])
	switch part {
	case "major":
		fmt.Printf("%d", v.Major)
		return nil
	case "minor":
		fmt.Printf("%d", v.Minor)
		return nil
	case "patch":
		fmt.Printf("%d", v.Patch)
		return nil
	case "pre":
		p := v.Pre
		if len(p) > 0 {
			fmt.Printf("%s", p)
		} else {
			fmt.Printf("")
		}
		return nil
	case "build":
		b := v.Build
		if len(b) > 0 {
			fmt.Printf("%s", b)
		} else {
			fmt.Printf("")
		}
		return nil
	}

	return nil
}
