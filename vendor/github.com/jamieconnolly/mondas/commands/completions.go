package commands

import "github.com/jamieconnolly/mondas/cli"

// CompletionsCommand displays the list of commands for autocompletion.
var CompletionsCommand = &cli.Command{
	Name:    "completions",
	Hidden:  true,
	Summary: "Display the list of commands for autocompletion",
	Action: func(ctx *cli.Context) int {
		for _, cmd := range ctx.App.Commands.Visible() {
			cli.Println(cmd.Name)
		}
		return 0
	},
}
