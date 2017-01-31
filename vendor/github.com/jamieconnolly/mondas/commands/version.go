package commands

import "github.com/jamieconnolly/mondas/cli"

// VersionCommand displays the version information.
var VersionCommand = &cli.Command{
	Name:    "version",
	Hidden:  true,
	Summary: "Display version information",
	Action: func(ctx *cli.Context) int {
		cli.Printf("%s version %s\n", ctx.App.Name, ctx.App.Version)
		return 0
	},
}
