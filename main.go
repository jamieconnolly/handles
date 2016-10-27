package main

import "github.com/jamieconnolly/mondas"

var Version string

func main() {
	mondas.SetVersion(Version)
	mondas.Run()
}
