package commands_test

import (
	"bytes"
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/jamieconnolly/mondas/commands"
	"github.com/stretchr/testify/assert"
)

func TestCompletionsCommand(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	app := &cli.App{
		Commands: cli.Commands{
			{Name: "foo"},
			{Name: "bar"},
			{Name: "baz", Hidden: true},
		},
	}

	buf := new(bytes.Buffer)
	cli.Stdout = buf

	exitCode := commands.CompletionsCommand.Run(&cli.Context{App: app})
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, "bar\nfoo\n", buf.String())
}
