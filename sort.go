package main

import (
	"fmt"
	"github.com/carvel-dev/semver/v4"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"sort"
)

const name = "semversort"
const version = "0.1.0"
const description = `Version Tester. Compare versions.

SemverSort is a tool for sorting a list of versions based on Carvels semver constraints

	$ semversort 1.1.0 1.1.1 1.2.3 0.1.1
	1.2.3
	1.1.1
	1.1.0
    0.1.1

See below for information about how to determine the number of failed tests.
`

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Usage = description
	app.Action = func(c *cli.Context) error { return run(c) }
	app.Version = version
	app.ArgsUsage = "BASE VERSION [VERSION [VERSION [...]]"
	app.Run(os.Args)
}

// context describes the relevant portion of a cli.Context.
//
// This abstraction makes mocking easy.
type context interface {
	Bool(string) bool
	Args() cli.Args
}

// run handles all the flags and then runs the main action.
func run(c context) error {
	args := c.Args()
	if args.Len() < 2 {
		return fmt.Errorf("not enough arguments")
	}

	versions := compare(args.Slice())

	sort.SliceStable(versions, func(i, j int) bool {
		return versions[i].Version.GT(versions[j].Version)
	})

	PrintVersion(versions)
	return nil
}

// compare compiles a base version comparator, and then compares all cases to it.
//
// It retuns an array of versions that passed, and an array of versions that failed.
func compare(cases []string) []SemverWrap {
	var versions []SemverWrap

	for _, t := range cases {
		ver, err := NewSemver(t)
		if err != nil {
			PrintError("Failed to parse %s", t)
			continue
		}

		versions = append(versions, ver)
	}

	return versions
}

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

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
		fmt.Fprintln(stdout, v.String())
	}
}

// PrintError prints to stderr.
func PrintError(msg string, args ...interface{}) {
	fmt.Fprintf(stderr, msg, args...)
	fmt.Fprintln(stderr)
}
