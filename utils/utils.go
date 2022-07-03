package utils

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func Archive(baseFolder string, archive_name string) error {
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

func Extract(path string) error {

    r, err := os.Open(path)
    if err != nil {
        fmt.Println("error")
    }
	
    uncompressedStream, err := gzip.NewReader(r)
    if err != nil {
        log.Fatal("ExtractTarGz: NewReader failed")
    }

    tarReader := tar.NewReader(uncompressedStream)

    for {
        header, err := tarReader.Next()

        if err == io.EOF {
            break
        }

        if err != nil {
			return err
        }

        switch header.Typeflag {
        case tar.TypeDir:
            if err := os.Mkdir(header.Name, 0755); err != nil {
				return err
            }
        case tar.TypeReg:
            outFile, err := os.Create(header.Name)
            if err != nil {
				return err
            }
            if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
            }
            outFile.Close()

        default:
			return errors.New(fmt.Sprintf("ExtractTarGz: uknown type: %s in %s", header.Typeflag, header.Name))
        }
    }
	return nil
}