package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func get_directory_to_serve_from_args() string {
	for index, arg := range os.Args {
		if strings.HasPrefix(arg, "--dir") {
			return os.Args[index + 1]
		}
	}
	return "packages-to-serve/"
}

func Server() error {
	fs := http.FileServer(http.Dir(get_directory_to_serve_from_args()))

	port := ":5001"
    fmt.Println("Server is running on port" + port)

	return http.ListenAndServe(port, fs)
}