package cli_test

import (
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/stretchr/testify/assert"
)

func TestArgs_Contains(t *testing.T) {
	args := cli.Args{"foo", "bar"}
	assert.True(t, args.Contains("foo"))
	assert.False(t, args.Contains("baz"))
}

func TestArgs_First(t *testing.T) {
	args := cli.Args{"foo", "bar", "baz"}
	assert.Equal(t, "foo", args.First())
}

func TestArgs_Index(t *testing.T) {
	args := cli.Args{"foo", "bar", "baz"}
	assert.Equal(t, "foo", args.Index(0))
	assert.Equal(t, "bar", args.Index(1))
	assert.Equal(t, "baz", args.Index(2))
	assert.Empty(t, args.Index(-1))
	assert.Empty(t, args.Index(3))
}
