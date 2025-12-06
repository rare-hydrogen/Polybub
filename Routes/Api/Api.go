package Api

import (
	"Polybub/Auth/BasicAuth"
	"Polybub/Auth/OAuth2"
	"Polybub/Routes/Api/FoobarApi"
	"Polybub/Routes/Api/LoginApi"
	"Polybub/Routes/Api/UserPasswordResetApi"
	"Polybub/Routes/Api/ValidateMfaApi"
	"Polybub/Utilities/Permissions"
	"net/http"
)

func AddApiRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/user-login", LoginApi.Handler)
	mux.HandleFunc("/api/user-password-reset", UserPasswordResetApi.Handler)
	OAuth2.JwtPermit(mux, "/api/validate-mfa", ValidateMfaApi.Handler, Permissions.MFA_CODE_CRU, nil)
	mux.HandleFunc("/api/foobar-basic", BasicAuth.BasicAuth(FoobarApi.Handler, "username", "password"))
}
