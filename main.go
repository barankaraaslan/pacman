package main

import (
	"os"
	"pacman/pkg"
)

func main() {
	if os.Args[1] == "self-package" {
		pkg.SelfPackage()
	}
}