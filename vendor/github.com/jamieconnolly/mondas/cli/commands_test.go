package cli_test

import (
	"bytes"
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/stretchr/testify/assert"
)

func TestCommand_Parse_WithExecutable(t *testing.T) {
	cmd := &cli.Command{Path: "testdata/foo-bar"}
	assert.False(t, cmd.Parsed())

	err := cmd.Parse()
	if assert.NoError(t, err) {
		assert.Equal(t, "Display \"Hello, world!\"", cmd.Summary)
		assert.Equal(t, "foo bar <baz>", cmd.Usage)
		assert.True(t, cmd.Completions)
		assert.True(t, cmd.Hidden)
		assert.True(t, cmd.Parsed())
	}
}

func TestCommand_Parse_WithNoExecutable(t *testing.T) {
	cmd := &cli.Command{}
	assert.False(t, cmd.Parsed())

	err := cmd.Parse()
	if assert.Error(t, err, "An error was expected") {
		assert.EqualError(t, err, "open : no such file or directory")
		assert.True(t, cmd.Parsed())
	}
}

func TestCommand_Parse_WithNotExistingExecutable(t *testing.T) {
	cmd := &cli.Command{Path: "testdata/foo-not-found"}
	assert.False(t, cmd.Parsed())

	err := cmd.Parse()
	if assert.Error(t, err, "An error was expected") {
		assert.EqualError(t, err, "open testdata/foo-not-found: no such file or directory")
		assert.True(t, cmd.Parsed())
	}
}

func TestCommand_Run_WithAction(t *testing.T) {
	var s string

	cmd := &cli.Command{
		Action: func(ctx *cli.Context) int {
			s = ctx.Command.Name
			return 0
		},
		Name: "foo",
	}

	exitCode := cmd.Run(&cli.Context{})
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, "foo", s)
}

func TestCommand_Run_WithExecutable(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	buf := new(bytes.Buffer)
	cli.Stdout = buf

	cmd := &cli.Command{
		Name: "foo",
		Path: "testdata/bar-succeed",
	}

	exitCode := cmd.Run(&cli.Context{})
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, "Succeeded!\n", buf.String())
}

func TestCommand_Run_WithFailingExecutable(t *testing.T) {
	oldStderr := cli.Stderr
	defer func() { cli.Stderr = oldStderr }()

	buf := new(bytes.Buffer)
	cli.Stderr = buf

	cmd := &cli.Command{
		Name: "foo",
		Path: "testdata/bar-fail",
	}

	exitCode := cmd.Run(&cli.Context{})
	assert.Equal(t, 42, exitCode)
	assert.Equal(t, "Failed!\n", buf.String())
}

func TestCommand_Run_WithNotExistingExecutable(t *testing.T) {
	oldStderr := cli.Stderr
	defer func() { cli.Stderr = oldStderr }()

	buf := new(bytes.Buffer)
	cli.Stderr = buf

	app := &cli.App{Name: "bar"}
	cmd := &cli.Command{
		Name: "foo",
		Path: "testdata/bar-not-found",
	}

	exitCode := cmd.Run(&cli.Context{App: app})
	assert.Equal(t, 1, exitCode)
	assert.Equal(t, "bar: 'foo' appears to be a valid command, but we were not "+
		"able to execute it. Maybe bar-not-found is broken?\n", buf.String())
}

func TestCommand_Visible(t *testing.T) {
	cmd1 := &cli.Command{Hidden: false}
	cmd2 := &cli.Command{Hidden: true}

	assert.True(t, cmd1.Visible())
	assert.False(t, cmd2.Visible())
}

func TestCommands_Add(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}

	cmds := cli.Commands{cmd1}
	cmds.Add(cmd2)
	assert.Len(t, cmds, 2)
	assert.Equal(t, cmd1, cmds[0])
	assert.Equal(t, cmd2, cmds[1])
}

func TestCommands_Lookup(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}

	cmds := cli.Commands{cmd1, cmd2}
	assert.Equal(t, cmd1, cmds.Lookup("one"))
	assert.Equal(t, cmd2, cmds.Lookup("two"))
	assert.Nil(t, cmds.Lookup("three"))
}

func TestCommands_SuggestionsFor(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}
	cmd3 := &cli.Command{Name: "three"}
	cmds := cli.Commands{cmd1, cmd2, cmd3}

	suggestions1 := cmds.SuggestionsFor("neo")
	assert.Len(t, suggestions1, 1)
	assert.Equal(t, cmd1, suggestions1[0])

	suggestions2 := cmds.SuggestionsFor("t")
	assert.Len(t, suggestions2, 2)
	assert.Equal(t, cmd3, suggestions2[0])
	assert.Equal(t, cmd2, suggestions2[1])
}

func TestCommands_Visible(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two", Hidden: true}
	cmd3 := &cli.Command{Name: "three"}
	cmds := cli.Commands{cmd1, cmd2, cmd3}

	visible := cmds.Visible()
	assert.Len(t, visible, 2)
	assert.Equal(t, cmd1, visible[0])
	assert.Equal(t, cmd3, visible[1])
}
