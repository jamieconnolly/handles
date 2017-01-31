package commands_test

import (
	"bytes"
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/jamieconnolly/mondas/commands"
	"github.com/stretchr/testify/assert"
)

func TestHelpCommand_ForApp(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	app := &cli.App{
		Commands: cli.Commands{
			{Name: "foo"},
			{Name: "bar"},
			{Name: "baz", Hidden: true},
		},
		Usage: "test <command> [<args>]",
	}

	buf := new(bytes.Buffer)
	cli.Stdout = buf

	exitCode := commands.HelpCommand.Run(&cli.Context{App: app})
	assert.Equal(t, 0, exitCode)
	assert.Contains(t, buf.String(), "Usage: test <command> [<args>]")
	assert.Contains(t, buf.String(), "bar")
	assert.Contains(t, buf.String(), "foo")
	assert.NotContains(t, buf.String(), "baz")
}

func TestHelpCommand_ForCommand(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	app := &cli.App{
		Commands: cli.Commands{
			{
				Description: "baz",
				Name:        "foo",
				Summary:     "bar",
				Usage:       "foo bar <baz>",
			},
		},
	}
	args := cli.Args([]string{"foo"})

	buf := new(bytes.Buffer)
	cli.Stdout = buf

	exitCode := commands.HelpCommand.Run(&cli.Context{App: app, Args: args})
	assert.Equal(t, 0, exitCode)
	assert.Contains(t, buf.String(), "Name:\n   foo - bar\n")
	assert.Contains(t, buf.String(), "Usage:\n   foo bar <baz>\n")
	assert.Contains(t, buf.String(), "Description:\n   baz")
}

func TestHelpCommand_ForCommandWithNoDescription(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	app := &cli.App{
		Commands: cli.Commands{
			{Name: "foo"},
		},
	}
	args := cli.Args([]string{"foo"})

	buf := new(bytes.Buffer)
	cli.Stdout = buf

	exitCode := commands.HelpCommand.Run(&cli.Context{App: app, Args: args})
	assert.Equal(t, 0, exitCode)
	assert.NotContains(t, buf.String(), "Description:")
}

func TestHelpCommand_ForCommandWithNoUsageMessage(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	app := &cli.App{
		Commands: cli.Commands{
			{
				Name:      "bar",
				ArgsUsage: "<baz>",
			},
		},
		Name: "foo",
	}
	args := cli.Args([]string{"bar"})

	buf := new(bytes.Buffer)
	cli.Stdout = buf

	exitCode := commands.HelpCommand.Run(&cli.Context{App: app, Args: args})
	assert.Equal(t, 0, exitCode)
	assert.Contains(t, buf.String(), "Usage:\n   foo bar <baz>")
}

func TestHelpCommand_ForUnknownCommand(t *testing.T) {
	oldStderr := cli.Stderr
	defer func() { cli.Stderr = oldStderr }()

	buf := new(bytes.Buffer)
	cli.Stderr = buf

	app := &cli.App{Name: "foo"}
	args := cli.Args([]string{"bar"})

	exitCode := commands.HelpCommand.Run(&cli.Context{App: app, Args: args})
	assert.Equal(t, 1, exitCode)
	assert.Equal(t, "foo: 'bar' is not a valid command. See 'foo help'.\n", buf.String())
}
