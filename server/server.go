package server

import (
	"fmt"
	"net/http"
)

func Server() error {
	fs := http.FileServer(http.Dir("packages-to-serve/"))

	port := ":5001"
    fmt.Println("Server is running on port" + port)

	return http.ListenAndServe(port, fs)
}