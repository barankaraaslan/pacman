package main

import (
	"os"
	"pacman/pkg"
)

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "self-package" {
			pkg.SelfPackage()
		}
	}
}