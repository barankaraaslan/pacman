package pkg

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestSelfPackage(t *testing.T) {
	path, err := filepath.Abs("../pacman")
	if err != nil {
		t.Fatal(err)
	}
	if err := exec.Command(path, "self-package").Run(); err != nil {
		t.Fatal(err)
	}
	archive_name := "package.tar.gz"
	_, err = os.Stat(archive_name)
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(archive_name)
}