package main

import (
	"fmt"
	"github.com/carvel-dev/semver/v4"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

const name = "semversort"
const version = "0.1.0"
const description = `Version Tester. Compare versions.

SemverSort is a tool for comparing two version strings, or comparing a version string
to a version range.

	$ semversort ">1.0.0" 1.1.0
	1.1.0

	$ semversort "<1.0.0" 1.1.0
	   # No output because nothing matched.

	$ semversort -f "<1.0.0" 1.1.0
	1.1.0   # -f returns all failures, rather than matches.

	$ semversort ">1.0.0" 1.1.0 1.1.1 1.2.3 0.1.1
	1.1.0
	1.1.1
	1.2.3

See below for information about how to determine the number of failed tests.

EXIT CODES:

vert returns exit codes based on the number of failed matches. There are a few
reserved exit codes:

- 128: The command was not called correctly.
- 256: A version failed to parse, and comparisons could not continue. This will
  occur if the original constraint version cannot be parsed. If any
  subsequent version fails to parse, it will simply be counted as a failure.

Any other error codes indicate the number of failed tests. For example:

	$ vert 1.2.3 1.2.3 1.2.4 1.2.5
	1.2.3
	$ echo $?
	2   # <-- Two tests failed.

BASE VERSIONS:

The base version may be in any of the following formats:

- An exact semantic version number
	- 1.2.3
	- v1.2.3
	- 1.2.3-alpha.1+10212015
- A semantic version range
	- *
	- !=1.0.0
	- >=1.2.3
	- >1.2.3,<1.3.2
	- ~1.2.0
	- ^2.3

VERSIONS:

Other than the base version, all other supplied versions must follow the
SemVer 2 spec. Examples:

	- 1.2.3
	- v1.2.3
	- 1.2.3-alpha.1+10212015
	- v1.2.3-alpha.1+10212015
	- 1 (equivalent to 1.0.0)
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

	pass := compare(args.Slice())

	out := pass

	PrintVersion(out)
	return nil
}

// compare compiles a base version comparator, and then compares all cases to it.
//
// It retuns an array of versions that passed, and an array of versions that failed.
func compare(cases []string) []semver.Version {
	var passed []semver.Version

	for _, t := range cases {
		ver, err := semver.Make(t)
		if err != nil {
			PrintError("Failed to parse %s", t)
			continue
		}

		passed = append(passed, ver)
	}

	return passed
}

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

// PrintVersion prints a list of versions to standard out.
func PrintVersion(vers []semver.Version) {
	for _, v := range vers {
		fmt.Fprintln(stdout, v.String())
	}
}

// PrintError prints to stderr.
func PrintError(msg string, args ...interface{}) {
	fmt.Fprintf(stderr, msg, args...)
	fmt.Fprintln(stderr)
}
