package Ui

import (
	"net/http"

	"Polybub/Auth/OAuth2"
	Ui "Polybub/Routes/Ui/Pages/Handlers"

	"Polybub/Utilities/Permissions"
)

func AddUiRoutes(mux *http.ServeMux) {
	addComponentRoutes(mux)
	addPageRoutes(mux)
}

func addComponentRoutes(mux *http.ServeMux) {
	//mux.HandleFunc("/login", Ui.LoginHandler)
	//OAuth2.JwtPermit(mux, "/dashboard", Ui.DashboardHandler, Permissions.DASHBOARD_R, nil)
}

func addPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", Ui.LoginHandler)
	mux.HandleFunc("/forgot-password", Ui.ForgotPasswordHandler)
	OAuth2.JwtPermit(mux, "/dashboard", Ui.DashboardHandler, Permissions.DASHBOARD_R, nil)

}
