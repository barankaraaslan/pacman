package pkg

import (
	"io/ioutil"
	"os"
	"pacman/utils"
)

func SelfPackage() error {
	// TODO: consider system umask for creating the file
	if err := os.Mkdir("package", 0755); err != nil {
        return err
    }
    if err := os.Mkdir("package/bin", 0755); err != nil {
        return err
    }

	input, err := ioutil.ReadFile(os.Args[0])
    if err != nil {
        return err
    }
	// TODO: copied filename should be taken from os.Args[0]
    err = ioutil.WriteFile("package/bin/pacman", input, 0644)
    if err != nil {
        return err
    }
	if err = utils.Archive("package", "package.tar.gz"); err != nil {
        return err
    }
    if err = os.RemoveAll("package"); err != nil{
        return err
    }
    return nil
}
