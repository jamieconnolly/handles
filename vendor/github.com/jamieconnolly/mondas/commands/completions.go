package commands

import "github.com/jamieconnolly/mondas/cli"

// ShowAppCompletions displays the completions for the given app.
func ShowAppCompletions(ctx *cli.Context) int {
	if cmds := ctx.App.Commands.Visible(); len(cmds) > 0 {
		for _, cmd := range cmds {
			cli.Println(cmd.Name)
		}
	}

	return 0
}

// ShowCommandCompletions displays the completions for the given command.
func ShowCommandCompletions(ctx *cli.Context) int {
	cmd := ctx.Command

	if !cmd.Parsed() {
		cmd.Parse()
	}

	if cmd.Completions {
		ctx.Args = append(ctx.Args[1:], "--complete")
		return cmd.Run(ctx)
	}

	return 0
}

// CompletionsCommand displays the completions.
var CompletionsCommand = &cli.Command{
	Name:      "completions",
	ArgsUsage: "<command>",
	Hidden:    true,
	Summary:   "Display completions",
	Action: func(ctx *cli.Context) int {
		if len(ctx.Args) == 0 {
			return ShowAppCompletions(ctx)
		}

		if cmd := ctx.App.LookupCommand(ctx.Args.First()); cmd != nil {
			ctx.Command = cmd
			return ShowCommandCompletions(ctx)
		}

		return 0
	},
}
