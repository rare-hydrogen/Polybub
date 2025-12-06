package Api

import (
	"Polybub/Auth/BasicAuth"
	"Polybub/Routes/Api/FoobarApi"
	"Polybub/Routes/Api/LoginApi"
	"Polybub/Routes/Api/UserPasswordResetApi"
	"Polybub/Routes/Api/ValidateMfaApi"
	"net/http"
)

func AddApiRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/user-login", LoginApi.Handler)
	mux.HandleFunc("/api/user-password-reset", UserPasswordResetApi.Handler)
	mux.HandleFunc("/api/validate-mfa", ValidateMfaApi.Handler)
	mux.HandleFunc("/api/foobar-basic", BasicAuth.BasicAuth(FoobarApi.Handler, "username", "password"))
}
