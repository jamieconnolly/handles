package mondas

import (
	"os"
	"path/filepath"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/jamieconnolly/mondas/commands"
	"github.com/kardianos/osext"
)

// CommandLine is the default application.
var CommandLine = New(filepath.Base(os.Args[0]), Version)

// Version is the version used for the default application.
var Version string

// AddCommand adds a command to the default application.
func AddCommand(cmd *cli.Command) {
	CommandLine.AddCommand(cmd)
}

// New creates a new application with some default commands.
func New(name string, version string) *cli.App {
	app := cli.NewApp(name, version)

	if exePath, err := osext.Executable(); err == nil {
		app.ExecPath = filepath.Join(exePath, "../../libexec")
	}

	app.AddCommand(commands.CompletionsCommand)
	app.AddCommand(commands.HelpCommand)
	app.AddCommand(commands.VersionCommand)

	return app
}

// Run runs the default application using the arguments from os.Args.
func Run() {
	cli.Exit(CommandLine.Run(os.Args[1:]))
}
