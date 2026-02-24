package Components

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Routes/Ui/Components/ComponentHandlers"
	"Polybub/Utilities/Permissions"
	"net/http"
)

func AddComponentRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/comp/forgot-password", ComponentHandlers.ForgotPasswordHandler)
	OAuth2.JwtPermit(mux, "/comp/setup-mfa", ComponentHandlers.SetupMfaHandler, Permissions.MFA_CODE_CRU, nil)
	OAuth2.JwtPermit(mux, "/comp/validate-mfa", ComponentHandlers.ValidateMfaHandler, Permissions.MFA_CODE_R, nil)
	OAuth2.JwtPermit(mux, "/comp/foobar", ComponentHandlers.FooBarHandler, Permissions.FOOBARS_CRUD, nil)
	OAuth2.JwtPermit(mux, "/comp/create-user", ComponentHandlers.CreateUserHandler, Permissions.USERS_CRUD, nil)
}
