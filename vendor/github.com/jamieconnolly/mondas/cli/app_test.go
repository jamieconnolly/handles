package cli_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	app := cli.NewApp("foo", "1.2.3")
	assert.Equal(t, "foo-", app.ExecPrefix)
	assert.Equal(t, "foo", app.Name)
	assert.Equal(t, "foo <command> [<args>]", app.Usage)
	assert.Equal(t, "1.2.3", app.Version)
}

func TestApp_AddCommand(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}
	cmds := cli.Commands{cmd1}
	app := &cli.App{Commands: cmds}

	app.AddCommand(cmd2)
	assert.Len(t, app.Commands, 2)
	assert.Equal(t, cmd1, app.Commands[0])
	assert.Equal(t, cmd2, app.Commands[1])
}

func TestApp_Initialize(t *testing.T) {
	envPath := os.Getenv("PATH")
	defer os.Setenv("PATH", envPath)

	os.Setenv("PATH", string(os.PathListSeparator)+envPath)

	app := &cli.App{
		ExecPath:   "testdata",
		ExecPrefix: "foo-",
		Name:       "foo",
	}
	assert.False(t, app.Initialized())

	app.Initialize()
	assert.Equal(t, os.Getenv("PATH"), strings.Join(
		[]string{app.ExecPath, "", envPath},
		string(os.PathListSeparator),
	))
	assert.Len(t, app.Commands, 1)
	assert.Equal(t, "bar", app.LookupCommand("bar").Name)
	assert.Equal(t, "testdata/foo-bar", app.LookupCommand("bar").Path)
	assert.True(t, app.Initialized())
}

func TestApp_Initialize_WithNoExecPath(t *testing.T) {
	envPath := os.Getenv("PATH")
	defer os.Setenv("PATH", envPath)

	app := &cli.App{
		ExecPrefix: "foo-",
		Name:       "foo",
	}
	assert.False(t, app.Initialized())

	app.Initialize()
	assert.Equal(t, envPath, os.Getenv("PATH"))
	assert.True(t, app.Initialized())
}

func TestApp_LookupCommand(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}
	cmds := cli.Commands{cmd1, cmd2}

	app := &cli.App{Commands: cmds}
	assert.Equal(t, cmd1, app.LookupCommand("one"))
	assert.Equal(t, cmd2, app.LookupCommand("two"))
	assert.Nil(t, app.LookupCommand("three"))
}

func TestApp_Run(t *testing.T) {
	var s string

	app := &cli.App{
		Commands: cli.Commands{
			{
				Action: func(ctx *cli.Context) int {
					s = ctx.Command.Name
					return 0
				},
				Name: "bar",
			},
		},
		ExecPath:   "testdata",
		ExecPrefix: "foo-",
		Name:       "foo",
	}

	exitCode := app.Run([]string{"bar"})
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, "bar", s)
}

func TestApp_Run_WithEmptyArguments(t *testing.T) {
	var s string

	app := &cli.App{
		Commands: cli.Commands{
			{
				Action: func(ctx *cli.Context) int {
					s = ctx.Command.Name
					return 0
				},
				Name: "help",
			},
		},
		ExecPath:   "testdata",
		ExecPrefix: "foo-",
		Name:       "foo",
	}

	exitCode := app.Run([]string{})
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, "help", s)
}

func TestApp_Run_WithHelpFlagInArguments(t *testing.T) {
	var s string

	app := &cli.App{
		Commands: cli.Commands{
			{
				Action: func(ctx *cli.Context) int {
					s = ctx.Args.First()
					return 0
				},
				Name: "help",
			},
		},
		ExecPath:   "testdata",
		ExecPrefix: "foo-",
		Name:       "foo",
	}

	exitCode := app.Run([]string{"bar", "baz", "--help"})
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, "bar", s)
}

func TestApp_Run_WithUnknownCommand(t *testing.T) {
	oldStderr := cli.Stderr
	defer func() { cli.Stderr = oldStderr }()

	buf := new(bytes.Buffer)
	cli.Stderr = buf

	app := &cli.App{
		ExecPath:   "testdata",
		ExecPrefix: "foo-",
		Name:       "foo",
	}

	exitCode := app.Run([]string{"baz"})
	assert.Equal(t, 1, exitCode)
	assert.Equal(t, "foo: 'baz' is not a valid command. See 'foo help'.\n", buf.String())
}

func TestApp_ShowUnknownCommandError_WithMultipleSuggestions(t *testing.T) {
	oldStderr := cli.Stderr
	defer func() { cli.Stderr = oldStderr }()

	buf := new(bytes.Buffer)
	cli.Stderr = buf

	app := &cli.App{
		Commands: cli.Commands{
			{Name: "bar"},
			{Name: "baz"},
		},
		Name: "foo",
	}

	exitCode := app.ShowUnknownCommandError("bat")
	assert.Equal(t, 1, exitCode)
	assert.Equal(t, "foo: 'bat' is not a valid command. See 'foo help'.\n"+
		"\nDid you mean one of these?\n\tbar\n\tbaz\n", buf.String())
}

func TestApp_ShowUnknownCommandError_WithNoSuggestions(t *testing.T) {
	oldStderr := cli.Stderr
	defer func() { cli.Stderr = oldStderr }()

	buf := new(bytes.Buffer)
	cli.Stderr = buf

	app := &cli.App{Name: "foo"}

	exitCode := app.ShowUnknownCommandError("bar")
	assert.Equal(t, 1, exitCode)
	assert.Equal(t, "foo: 'bar' is not a valid command. See 'foo help'.\n", buf.String())
}

func TestApp_ShowUnknownCommandError_WithSingleSuggestion(t *testing.T) {
	oldStderr := cli.Stderr
	defer func() { cli.Stderr = oldStderr }()

	buf := new(bytes.Buffer)
	cli.Stderr = buf

	app := &cli.App{
		Commands: cli.Commands{
			{Name: "bar"},
		},
		Name: "foo",
	}

	exitCode := app.ShowUnknownCommandError("baz")
	assert.Equal(t, 1, exitCode)
	assert.Equal(t, "foo: 'baz' is not a valid command. See 'foo help'.\n"+
		"\nDid you mean this?\n\tbar\n", buf.String())
}
