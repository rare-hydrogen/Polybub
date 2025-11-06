package BasicAuth

import (
	"crypto/subtle"
	"net/http"
)

func BasicAuth(
	handler http.HandlerFunc, username string, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		realm := "Basic Auth credentials required."

		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}

		handler(w, r)
	}
}
