package Pages

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Routes/Ui/Pages/PageHandlers"
	"Polybub/Utilities/Permissions"
	"net/http"
)

func AddPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", PageHandlers.LoginHandler)
	mux.HandleFunc("/forgot-password", PageHandlers.ForgotPasswordHandler)
	OAuth2.JwtPermit(mux, "/setup-mfa", PageHandlers.SetupMfaHandler, Permissions.MFA_CODE_CRU, nil)
	OAuth2.JwtPermit(mux, "/validate-mfa", PageHandlers.ValidateMfaHandler, Permissions.MFA_CODE_R, nil)
	OAuth2.JwtPermit(mux, "/dashboard", PageHandlers.DashboardHandler, Permissions.DASHBOARD_R, nil)
	OAuth2.JwtPermit(mux, "/create-user", PageHandlers.CreateUserHandler, Permissions.USERS_CRUD, nil)
}
