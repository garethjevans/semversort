package cmd

import (
	"fmt"
	"github.com/carvel-dev/semver/v4"
	"github.com/garethjevans/semver/pkg/bump"
	"github.com/spf13/cobra"
)

var (
	bumpCommand = &cobra.Command{
		Use:   "bump",
		Short: "bump a part of a carvel compatible semantic version",
		Long:  `bump a part of a carvel compatible semantic version`,
		RunE:  run,
	}
)

func init() {
	bumpCommand.Flags().StringVarP(&Out, "out", "o", "", "Write the output to a file")
}

func run(cmd *cobra.Command, args []string) error {
	part := args[0]
	in := ReadFromArgsOrStdin(args[len(args)-1:], cmd.InOrStdin())

	v := semver.MustParse(in[0])
	switch part {
	case "major":
		b := bump.MajorBump{}
		v = b.Apply(v)

		fmt.Printf("%s", v)
		return nil
	case "minor":
		b := bump.MinorBump{}
		v = b.Apply(v)

		fmt.Printf("%s", v)
		return nil
	case "patch":
		b := bump.PatchBump{}
		v = b.Apply(v)

		fmt.Printf("%s", v)
		return nil
	case "pre":
		b := bump.PreBump{
			Pre: args[1],
		}
		v = b.Apply(v)

		fmt.Printf("%s", v)
		return nil
	default:
		return fmt.Errorf("unknown part '%s'", part)
	}
}
