package Api

import (
	"Polybub/Auth/BasicAuth"
	Api "Polybub/Routes/Api/Handlers"
	"net/http"
)

func AddApiRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/user-login", Api.LoginHandler)
	mux.HandleFunc("/api/user-password-reset", Api.UserPasswordResetHandler)
	mux.HandleFunc("/api/foobar-basic", BasicAuth.BasicAuth(Api.FooBarHandler, "username", "password"))
}
