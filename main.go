package main

import "github.com/jamieconnolly/mondas"

var (
	Name string
	Version string
)

func main() {
	mondas.Run(Name, Version)
}
