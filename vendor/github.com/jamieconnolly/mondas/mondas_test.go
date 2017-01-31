package mondas_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/jamieconnolly/mondas"
	"github.com/jamieconnolly/mondas/cli"
	"github.com/stretchr/testify/assert"
)

func TestAddCommand(t *testing.T) {
	cmd1 := &cli.Command{Name: "one"}
	cmd2 := &cli.Command{Name: "two"}
	cmds := cli.Commands{cmd1}
	mondas.CommandLine = &cli.App{Commands: cmds}

	mondas.AddCommand(cmd2)
	assert.Equal(t, 2, len(mondas.CommandLine.Commands))
	assert.Equal(t, cmd1.Name, mondas.CommandLine.Commands[0].Name)
	assert.Equal(t, cmd2.Name, mondas.CommandLine.Commands[1].Name)
}

func TestNew(t *testing.T) {
	app := mondas.New("foo", "1.2.3")
	assert.Equal(t, "foo", app.Name)
	assert.Equal(t, "1.2.3", app.Version)
}

func TestRun(t *testing.T) {
	var s string

	cli.Exit = func(code int) {
		s = "exited"
	}
	cli.Stderr = ioutil.Discard
	cli.Stdout = ioutil.Discard
	mondas.CommandLine = &cli.App{ExecPrefix: "foo-", Name: "foo"}
	os.Args = []string{"foo", "bar"}

	mondas.Run()
	assert.Equal(t, "exited", s)
}
