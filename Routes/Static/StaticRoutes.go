package Static

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed Files/*
var Files embed.FS

func AddStaticRoutes(mux *http.ServeMux) {
	sub, err := fs.Sub(Files, "Files")
	if err == nil {
		mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(sub))))
	}
}
