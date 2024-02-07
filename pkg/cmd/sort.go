package cmd

import (
	"fmt"
	"github.com/carvel-dev/semver/v4"
	"github.com/spf13/cobra"
	"sort"
)

var (
	sortCommand = &cobra.Command{
		Use:   "sort",
		Short: "sort a series of carvel compatible semantic versions",
		Long:  `sort a series of carvel compatible semantic versions.  This will also remove any versions that are not compatible.`,
		RunE:  sortCmd,
		// FIXME two args
	}
	rangeString string
)

func init() {
	sortCommand.Flags().StringVarP(&rangeString, "range", "r", "", "Use a range to filter versions")
}

func sortCmd(cmd *cobra.Command, args []string) error {
	in := ReadFromArgsOrStdin(args[0:], cmd.InOrStdin())
	versions := compare(in)

	sort.SliceStable(versions, func(i, j int) bool {
		return versions[i].Version.GT(versions[j].Version)
	})

	PrintVersion(versions)
	return nil
}

// compare compiles a base version comparator, and then compares all cases to it.
//
// It returns an array of versions that passed, and an array of versions that failed.
func compare(cases []string) []SemverWrap {
	var err error
	var r semver.Range

	if rangeString != "" {
		r, err = semver.ParseRange(rangeString)
		if err != nil {
			fmt.Printf("Unable to parse range '%s'\n", rangeString)
		}
	}

	var versions []SemverWrap
	for _, t := range cases {
		ver, err := NewSemver(t)
		if err != nil {
			// ignore
			continue
		}

		if rangeString != "" {
			if r(ver.Version) {
				versions = append(versions, ver)
			}
		} else {
			versions = append(versions, ver)
		}
	}

	return versions
}

type SemverWrap struct {
	semver.Version
	Original string
}

func NewSemver(version string) (SemverWrap, error) {
	parsedVersion, err := semver.Parse(version)
	if err != nil {
		return SemverWrap{}, err
	}

	return SemverWrap{parsedVersion, version}, nil
}

// PrintVersion prints a list of versions to standard out.
func PrintVersion(vers []SemverWrap) {
	for _, v := range vers {
		fmt.Println(v.Original)
	}
}
