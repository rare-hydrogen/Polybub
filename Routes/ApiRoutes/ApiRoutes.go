package ApiRoutes

import (
	"net/http"

	"Polybub/Auth/BasicAuth"
	"Polybub/Routes/ApiHandlers"
)

func AddApiRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/login", ApiHandlers.LoginHandler)
	mux.HandleFunc("/api/user-password-reset", ApiHandlers.UserPasswordResetHandler)
	mux.HandleFunc("/api/foobar-basic", BasicAuth.BasicAuth(ApiHandlers.FooBarHandler, "username", "password"))
}
