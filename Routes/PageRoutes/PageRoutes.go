package PageRoutes

import (
	"net/http"

	"Polybub/Routes/PageRoutes/Dashboard"
	"Polybub/Routes/PageRoutes/Login"
)

func AddPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", Login.PageHandler)
	mux.HandleFunc("/dashboard", Dashboard.Handler)
}
