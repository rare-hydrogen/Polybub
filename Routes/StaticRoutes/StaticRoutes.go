package StaticRoutes

import (
	"net/http"
)

func AddStaticRoutes(mux *http.ServeMux) {
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./Static"))))
}
