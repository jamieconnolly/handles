package cli_test

import (
	"bytes"
	"testing"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	oldStderr := cli.Stderr
	defer func() { cli.Stderr = oldStderr }()

	buf := new(bytes.Buffer)
	cli.Stderr = buf

	cli.Error("foo")

	assert.Equal(t, "foo", buf.String())
}

func TestErrorf(t *testing.T) {
	oldStderr := cli.Stderr
	defer func() { cli.Stderr = oldStderr }()

	buf := new(bytes.Buffer)
	cli.Stderr = buf

	cli.Errorf("%s %d", "foo", 42)

	assert.Equal(t, "foo 42", buf.String())
}

func TestErrorln(t *testing.T) {
	oldStderr := cli.Stderr
	defer func() { cli.Stderr = oldStderr }()

	buf := new(bytes.Buffer)
	cli.Stderr = buf

	cli.Errorln("foo")

	assert.Equal(t, "foo\n", buf.String())
}

func TestPrint(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	buf := new(bytes.Buffer)
	cli.Stdout = buf

	cli.Print("foo")

	assert.Equal(t, "foo", buf.String())
}

func TestPrintf(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	buf := new(bytes.Buffer)
	cli.Stdout = buf

	cli.Printf("%s %d", "foo", 42)

	assert.Equal(t, "foo 42", buf.String())
}

func TestPrintln(t *testing.T) {
	oldStdout := cli.Stdout
	defer func() { cli.Stdout = oldStdout }()

	buf := new(bytes.Buffer)
	cli.Stdout = buf

	cli.Println("foo")

	assert.Equal(t, "foo\n", buf.String())
}
