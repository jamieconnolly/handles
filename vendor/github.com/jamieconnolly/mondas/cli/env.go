package cli

import (
	"fmt"
	"strings"
)

// Env represents the environment of a process.
type Env map[string]string

// NewEnvFromEnviron creates a new Env object from a list of strings
// in the form of "key=value".
func NewEnvFromEnviron(environ []string) Env {
	env := make(Env)
	for _, variable := range environ {
		parts := strings.SplitN(variable, "=", 2)
		if len(parts) > 1 {
			env[parts[0]] = parts[1]
		} else {
			env[parts[0]] = ""
		}
	}
	return env
}

// Environ returns a list of strings representing the environment,
// in the form of "key=value".
func (e Env) Environ() []string {
	result := make([]string, 0, len(e))
	for key, value := range e {
		result = append(result, fmt.Sprintf("%s=%s", key, value))
	}
	return result
}

// Get returns the value of the named environment variable.
func (e Env) Get(key string) string {
	return e[key]
}

// Set sets the value of the named environment variable.
func (e Env) Set(key string, value string) {
	e[key] = value
}

// Unset removes the named environment variable.
func (e Env) Unset(key string) {
	delete(e, key)
}
