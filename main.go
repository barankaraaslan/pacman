package main

import (
	"os"
	"pacman/pkg"
	"pacman/server"
)

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "self-package" {
			pkg.SelfPackage()
		}
	}
	if len(os.Args) == 2 {
		if os.Args[1] == "server" {
			server.Server()
		}
	}
}