package Pages

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Routes/Ui/Pages/Handlers"
	"Polybub/Utilities/Permissions"
	"net/http"
)

func AddPageRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", Handlers.LoginHandler)
	mux.HandleFunc("/forgot-password", Handlers.ForgotPasswordHandler)
	OAuth2.JwtPermit(mux, "/setup-mfa", Handlers.SetupMfaHandler, Permissions.MFA_CODE_CRU, nil)
	OAuth2.JwtPermit(mux, "/validate-mfa", Handlers.ValidateMfaHandler, Permissions.MFA_CODE_R, nil)
	OAuth2.JwtPermit(mux, "/dashboard", Handlers.DashboardHandler, Permissions.DASHBOARD_R, nil)
	OAuth2.JwtPermit(mux, "/create-user", Handlers.CreateUserHandler, Permissions.DASHBOARD_R, nil) // Fix perms
}
