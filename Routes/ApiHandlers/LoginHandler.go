package ApiHandlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Services"
	"Polybub/Jsend"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		login(w, req)
	case http.MethodDelete:
		logout(w, req)
	default:
		Jsend.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	username := req.Header.Get("Username")
	password := req.Header.Get("Password")

	jwtString, err := Services.Login(username, password)
	if err != nil {
		Jsend.Error(w, "login failed", http.StatusBadRequest)
		return
	}

	OAuth2.StoreTokenAndRedirect(w, jwtString, "dashboard")
}

func logout(w http.ResponseWriter, req *http.Request) {
	OAuth2.DeleteTokenAndRedirect(w, "login")
}
