package Pages

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Routes/Ui/Pages/Dashboard"
	"Polybub/Routes/Ui/Pages/ForgotPassword"
	"Polybub/Routes/Ui/Pages/Login"
	"Polybub/Routes/Ui/Pages/SetupMfa"
	"Polybub/Utilities/Permissions"
	"net/http"
)

func AddPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", Login.Handler)
	mux.HandleFunc("/forgot-password", ForgotPassword.Handler)
	OAuth2.JwtPermit(mux, "/dashboard", Dashboard.Handler, Permissions.DASHBOARD_R, nil)
	OAuth2.JwtPermit(mux, "/setup-mfa", SetupMfa.Handler, Permissions.DASHBOARD_R, nil) // TODO: Handle this permission
}
