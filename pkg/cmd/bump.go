package cmd

import (
	"fmt"
	"github.com/carvel-dev/semver/v4"
	"github.com/spf13/cobra"
)

var (
	bumpCommand = &cobra.Command{
		Use:   "bump",
		Short: "bump a part of a carvel compatible semantic version",
		Long:  `bump a part of a carvel compatible semantic version`,
		RunE:  bump,
	}
)

func init() {
	bumpCommand.Flags().StringVarP(&Out, "out", "o", "", "Write the output to a file")
}

func bump(cmd *cobra.Command, args []string) error {
	part := args[0]
	in := ReadFromArgsOrStdin(args[1:], cmd.InOrStdin())

	v := semver.MustParse(in[0])
	switch part {
	case "major":
		err := v.IncrementMajor()
		if err != nil {
			return err
		}

		fmt.Printf("%d.%d.%d", v.Major, v.Minor, v.Patch)
		return nil
	case "minor":
		err := v.IncrementMinor()
		if err != nil {
			return err
		}

		fmt.Printf("%d.%d.%d", v.Major, v.Minor, v.Patch)
		return nil
	case "patch":
		err := v.IncrementPatch()
		if err != nil {
			return err
		}

		fmt.Printf("%d.%d.%d", v.Major, v.Minor, v.Patch)
		return nil
	default:
		return fmt.Errorf("unknown part '%s'", part)
	}
}
