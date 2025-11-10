package PageRoutes

import (
	"net/http"

	"Polybub/Routes/PageRoutes/Login"
)

func AddPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", Login.PageHandler)
}
