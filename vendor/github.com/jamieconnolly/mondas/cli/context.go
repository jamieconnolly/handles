package cli

// Context represents the context of a process.
type Context struct {
	App     *App
	Args    Args
	Command *Command
	Env     Env
}

// NewContext creates a new Context object.
func NewContext(app *App, args []string, env []string) *Context {
	return &Context{
		App:  app,
		Args: Args(args),
		Env:  NewEnvFromEnviron(env),
	}
}
