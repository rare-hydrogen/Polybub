package ApiRoutes

import (
	"net/http"

	"Polybub/Auth/BasicAuth"
	"Polybub/Routes/ApiRoutes/Foobar"
)

func AddApiRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/foobar", Foobar.Handler)

	mux.HandleFunc("/foobar-basic-auth", BasicAuth.BasicAuth(Foobar.Handler, "username", "password"))
}
