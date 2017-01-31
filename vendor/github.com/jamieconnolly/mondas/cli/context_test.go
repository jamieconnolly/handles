package cli_test

import (
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	app := &cli.App{}
	args := []string{"foo"}
	env := []string{"foo=bar", "baz=foo"}

	ctx := cli.NewContext(app, args, env)
	assert.Equal(t, app, ctx.App)
	assert.Len(t, ctx.Args, 1)
	assert.Equal(t, "foo", ctx.Args.First())
	assert.Nil(t, ctx.Command)
	assert.Equal(t, "bar", ctx.Env.Get("foo"))
}
