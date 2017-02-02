package commands_test

import (
	"bytes"
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/jamieconnolly/mondas/commands"
	"github.com/stretchr/testify/assert"
)

func TestCompletionsCommand_ForApp(t *testing.T) {
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

func TestCompletionsCommand_ForCommand(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	var a []string

	app := &cli.App{
		Commands: cli.Commands{
			{
				Action: func(ctx *cli.Context) int {
					a = ctx.Args
					return 0
				},
				Name:        "foo",
				Completions: true,
			},
		},
	}
	args := cli.Args([]string{"foo", "bar", "baz"})

	buf := new(bytes.Buffer)
	cli.Stdout = buf

	exitCode := commands.CompletionsCommand.Run(&cli.Context{App: app, Args: args})
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, []string{"bar", "baz", "--complete"}, a)
	assert.Empty(t, buf.String())
}

func TestCompletionsCommand_ForCommandWithNoCompletions(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	app := &cli.App{
		Commands: cli.Commands{
			{Name: "foo"},
		},
	}
	args := cli.Args([]string{"foo", "bar", "baz"})

	buf := new(bytes.Buffer)
	cli.Stdout = buf

	exitCode := commands.CompletionsCommand.Run(&cli.Context{App: app, Args: args})
	assert.Equal(t, 0, exitCode)
	assert.Empty(t, buf.String())
}

func TestCompletionsCommand_ForUnknownCommand(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	buf := new(bytes.Buffer)
	cli.Stdout = buf

	app := &cli.App{Name: "foo"}
	args := cli.Args([]string{"bar"})

	exitCode := commands.CompletionsCommand.Run(&cli.Context{App: app, Args: args})
	assert.Equal(t, 0, exitCode)
	assert.Empty(t, buf.String())
}
