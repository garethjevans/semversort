package main

import (
	"bytes"
	"github.com/urfave/cli/v2"
	"testing"
)

type mockContext struct {
	bools map[string]bool
	args  cli.Args
}

func (c *mockContext) Args() cli.Args {
	return c.args
}
func (c *mockContext) Bool(name string) bool {
	return c.bools[name]
}

func TestRun(t *testing.T) {
	var b bytes.Buffer

	c := &mockContext{
		args: NewMockArgs([]string{"1.0.0", "1.1.1", "1.2.3", "1.0.1", "0.9.0"}),
		bools: map[string]bool{
			"failed": false,
			"sort":   true,
		},
	}

	// Set the package defaults
	stdout = &b
	stderr = &b

	tests := []struct {
		args cli.Args
		out  string
	}{
		{
			args: NewMockArgs([]string{"1.0.0", "1.1.1", "1.2.3", "1.0.1", "0.9.0"}),
			out:  "1.2.3\n1.1.1\n1.0.1\n1.0.0\n0.9.0\n",
		},
	}

	for _, tt := range tests {
		c.args = tt.args
		res := run(c)
		if res != nil {
			t.Errorf("Expected no error, got %v", res)
		}
		if b.String() != tt.out {
			t.Errorf("Expected:%s\nGot:%s", tt.out, b.String())
		}
		b.Reset()
	}

}

func NewMockArgs(args []string) cli.Args {
	return &mockArgs{values: args}
}

type mockArgs struct {
	values []string
}

func (a *mockArgs) Get(n int) string {
	if len(a.values) > n {
		return (a.values)[n]
	}
	return ""
}

func (a *mockArgs) First() string {
	return a.Get(0)
}

func (a *mockArgs) Tail() []string {
	if a.Len() >= 2 {
		tail := []string((a.values)[1:])
		ret := make([]string, len(tail))
		copy(ret, tail)
		return ret
	}
	return []string{}
}

func (a *mockArgs) Len() int {
	return len(a.values)
}

func (a *mockArgs) Present() bool {
	return a.Len() != 0
}

func (a *mockArgs) Slice() []string {
	ret := make([]string, len(a.values))
	copy(ret, a.values)
	return ret
}
