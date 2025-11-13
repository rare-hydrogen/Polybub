package PageRoutes

import (
	"net/http"

	"Polybub/Auth/OAuth2"
	"Polybub/Routes/PageRoutes/Dashboard"
	"Polybub/Routes/PageRoutes/Login"
	"Polybub/Utilities/Permissions"
)

func AddPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", Login.PageHandler)
	OAuth2.JwtPermit(mux, "/dashboard", Dashboard.Handler, Permissions.DASHBOARD_R, nil)
}
