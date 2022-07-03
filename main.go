package main

import (
	"fmt"
	"os"
	"pacman/pkg"
	"pacman/server"
)

const version = "0.0.1"

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
	for _, args := range os.Args {
		if args == "--version" {
			fmt.Println(version)
		}
	}
}