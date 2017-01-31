package commands_test

import (
	"bytes"
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/jamieconnolly/mondas/commands"
	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	app := &cli.App{Name: "foo", Version: "1.2.3"}
	buf := new(bytes.Buffer)
	cli.Stdout = buf

	exitCode := commands.VersionCommand.Run(&cli.Context{App: app})
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, "foo version 1.2.3\n", buf.String())
}
