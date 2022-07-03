package pkg

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path"
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
	if err = package_directory("package", "package.tar.gz"); err != nil {
        return err
    }
    if err = os.RemoveAll("package"); err != nil{
        return err
    }
    return nil
}

func package_directory(baseFolder string, archive_name string) error {
    // Get a Buffer to Write To
    outFile, err := os.Create(archive_name)
    if err != nil {
		return err
	}
    defer outFile.Close()

	gw := gzip.NewWriter(outFile)
	defer gw.Close()
	w := tar.NewWriter(gw)
	defer w.Close()


    // Add some files to the archive.
    add_file_to_archive(w, baseFolder, "")

    if err != nil {
		return err
    }

    // Make sure to check the error on Close.
    err = w.Close()
    if err != nil {
		return err
	}
    return nil 
}

func add_file_to_archive(w *tar.Writer, basePath, baseInZip string) error {
    // Open the Directory
    files, err := ioutil.ReadDir(basePath)
    if err != nil {
        return err
    }

    for _, file := range files {
        if !file.IsDir() {
            // Open the file which will be written into the archive
            file, err := os.Open(path.Join(basePath, file.Name()))
            if err != nil {
                return err
            }
            defer file.Close()

            // Get FileInfo about our file providing file size, mode, etc.
            info, err := file.Stat()
            if err != nil {
                return err
            }

            // Create a tar Header from the FileInfo data
            header, err := tar.FileInfoHeader(info, info.Name())
            if err != nil {
                return err
            }

            // Use full path as name (FileInfoHeader only takes the basename)
            // If we don't do this the directory strucuture would
            // not be preserved
            // https://golang.org/src/archive/tar/common.go?#L626
            header.Name = path.Join(baseInZip, info.Name())

            // Write file header to the tar archive
            err = w.WriteHeader(header)
            if err != nil {
                return err
            }

            // Copy file content to tar archive
            _, err = io.Copy(w, file)
            if err != nil {
                return err
            }

        } else if file.IsDir() {

            // Recurse
            newBase := path.Join(basePath, file.Name())
            add_file_to_archive(w, newBase, path.Join(baseInZip, file.Name()))
        }
    }
    return nil
}