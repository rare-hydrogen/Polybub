package Api

import (
	"Polybub/Auth/BasicAuth"
	"Polybub/Routes/Api/Handlers"
	"net/http"
)

func AddApiRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/user-login", Handlers.LoginHandler)
	mux.HandleFunc("/api/user-password-reset", Handlers.UserPasswordResetHandler)
	mux.HandleFunc("/api/validate-mfa", Handlers.ValdiateMfaHandler)
	mux.HandleFunc("/api/foobar", Handlers.FooBarHandler)
	mux.HandleFunc("/api/foobar-basic", BasicAuth.BasicAuth(Handlers.FooBarHandler, "username", "password"))
}
