package cli

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// App represents the entire command-line interface.
type App struct {
	// Commands is the list of available commands
	Commands Commands
	// ExecPath is the directory where program executables are stored
	ExecPath string
	// ExecPrefix is the prefix for program executables
	ExecPrefix string
	// Name is the name of the application
	Name string
	// Usage is the one-line usage message
	Usage string
	// Version is the version of the application
	Version string

	initialized bool
}

// NewApp creates a new App object.
func NewApp(name string, version string) *App {
	return &App{
		ExecPrefix: name + "-",
		Name:       name,
		Usage:      name + " <command> [<args>]",
		Version:    version,
	}
}

// AddCommand adds a command to the application.
func (a *App) AddCommand(cmd *Command) {
	a.Commands.Add(cmd)
}

// Initialize prepends the exec path to PATH then populates the list
// of commands with program executables and the help command.
func (a *App) Initialize() {
	a.initialized = true

	if a.ExecPath != "" {
		os.Setenv("PATH", strings.Join(
			[]string{a.ExecPath, os.Getenv("PATH")},
			string(os.PathListSeparator),
		))
	}

	for _, dir := range filepath.SplitList(os.Getenv("PATH")) {
		if dir == "" {
			continue
		}
		files, _ := filepath.Glob(filepath.Join(dir, a.ExecPrefix+"*"))
		for _, file := range files {
			if _, err := exec.LookPath(file); err == nil {
				a.AddCommand(&Command{
					Name: strings.TrimPrefix(filepath.Base(file), a.ExecPrefix),
					Path: file,
				})
			}
		}
	}
}

// Initialized returns true if the application has been initialized.
func (a *App) Initialized() bool {
	return a.initialized
}

// LookupCommand returns a command matching the given name.
func (a *App) LookupCommand(name string) *Command {
	return a.Commands.Lookup(name)
}

// Run parses the given argument list and runs the matching command.
func (a *App) Run(arguments []string) int {
	if !a.Initialized() {
		a.Initialize()
	}

	args := Args(arguments)

	if len(args) == 0 {
		args = Args([]string{"help"})
	}

	args[0] = strings.TrimPrefix(args.First(), "--")

	if len(args) > 1 && args.Contains("--help") {
		args = Args([]string{"help", args.First()})
	}

	if cmd := a.LookupCommand(args.First()); cmd != nil {
		return cmd.Run(NewContext(a, args[1:], os.Environ()))
	}

	return a.ShowUnknownCommandError(args.First())
}

// ShowUnknownCommandError shows a list of suggested commands
// based on the given name then exits with status code 1.
func (a *App) ShowUnknownCommandError(typedName string) int {
	Errorf("%s: '%s' is not a valid command. See '%s help'.\n",
		a.Name, typedName, a.Name)

	if suggestions := a.Commands.SuggestionsFor(typedName); len(suggestions) > 0 {
		if len(suggestions) == 1 {
			Errorln("\nDid you mean this?")
		} else {
			Errorln("\nDid you mean one of these?")
		}
		for _, cmd := range suggestions {
			Errorf("\t%s\n", cmd.Name)
		}
	}

	return 1
}
