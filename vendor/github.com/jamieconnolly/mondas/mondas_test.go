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
	mondas.Commands = cli.Commands{cmd1}

	mondas.AddCommand(cmd2)
	assert.Equal(t, 2, len(mondas.Commands))
	assert.Equal(t, cmd1.Name, mondas.Commands[0].Name)
	assert.Equal(t, cmd2.Name, mondas.Commands[1].Name)
}

func TestRun(t *testing.T) {
	var s string

	cli.Exit = func(code int) {
		s = "exited"
	}
	cli.Stderr = ioutil.Discard
	cli.Stdout = ioutil.Discard
	os.Args = []string{"foo", "bar"}

	mondas.Run("foo", "1.2.3")
	assert.Equal(t, "exited", s)
}
