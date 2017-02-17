package mondas

import (
	"os"
	"path/filepath"

	"github.com/jamieconnolly/mondas/cli"
	"github.com/jamieconnolly/mondas/commands"
)

// Commands is a list of commands for the default application to use.
var Commands cli.Commands

// AddCommand adds the command to the default application.
func AddCommand(cmd *cli.Command) {
	Commands.Add(cmd)
}

// Run is the main entry point for the command-line interface.
// It creates an application with some reasonable defaults
// and then runs it using the arguments from os.Args.
func Run(name string, version string) {
	app := cli.NewApp(name, version)

	app.Commands = Commands

	if exePath, err := os.Executable(); err == nil {
		realPath, _ := filepath.EvalSymlinks(exePath)
		app.ExecPath = filepath.Join(realPath, "../../libexec")
	}

	app.AddCommand(commands.CompletionsCommand)
	app.AddCommand(commands.HelpCommand)
	app.AddCommand(commands.VersionCommand)

	cli.Exit(app.Run(os.Args[1:]))
}
