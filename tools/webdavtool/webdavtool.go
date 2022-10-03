package main

import (
	"net/http"

	"golang.org/x/net/webdav"
)

func main() {
	http.ListenAndServe(":8080", &webdav.Handler{
		FileSystem: webdav.Dir("D:/T"),
		LockSystem: webdav.NewMemLS(),
	})
}
