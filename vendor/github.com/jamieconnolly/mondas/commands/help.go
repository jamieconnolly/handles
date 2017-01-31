package commands

import (
	"github.com/jamieconnolly/mondas/cli"
	"github.com/kr/text"
)

// ShowAppHelp displays the help information for the given app.
func ShowAppHelp(ctx *cli.Context) int {
	cli.Printf("Usage: %s\n", ctx.App.Usage)

	if cmds := ctx.App.Commands.Visible(); len(cmds) > 0 {
		cli.Println("\nCommands:")
		for _, c := range cmds {
			cli.Printf("   %-15s   %s\n", c.Name, c.Summary)
		}
	}

	return 0
}

// ShowCommandHelp displays the help information for the given command.
func ShowCommandHelp(ctx *cli.Context) int {
	c := ctx.Command

	if !c.Parsed() {
		c.Parse()
	}

	cli.Println("Name:")
	cli.Printf("   %s - %s\n", c.Name, c.Summary)

	cli.Println("\nUsage:")
	if c.Usage != "" {
		cli.Println(text.Indent(c.Usage, "   "))
	} else {
		cli.Printf("   %s %s %s\n", ctx.App.Name, c.Name, c.ArgsUsage)
	}

	if c.Description != "" {
		cli.Println("\nDescription:")
		cli.Println(text.Indent(c.Description, "   "))
	}

	return 0
}

// HelpCommand displays the help information.
var HelpCommand = &cli.Command{
	Name:      "help",
	ArgsUsage: "<command>",
	Summary:   "Display help information",
	Action: func(ctx *cli.Context) int {
		if len(ctx.Args) == 0 {
			return ShowAppHelp(ctx)
		}

		if cmd := ctx.App.LookupCommand(ctx.Args.First()); cmd != nil {
			ctx.Command = cmd
			return ShowCommandHelp(ctx)
		}

		return ctx.App.ShowUnknownCommandError(ctx.Args.First())
	},
}
