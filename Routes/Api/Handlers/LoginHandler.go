package Handlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Models"
	"Polybub/Data/Services"
	"Polybub/Jsend"
	"Polybub/Utilities/Permissions"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		postLogin(w, req)
	case http.MethodDelete:
		deleteLogout(w, req)
	default:
		Jsend.Error(req.Context(), w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func postLogin(w http.ResponseWriter, req *http.Request) {
	username := req.Header.Get("Username")
	password := req.Header.Get("Password")

	user, jwtString, err := Services.Login(username, password)
	if err != nil {
		Jsend.Error(req.Context(), w, "login failed", http.StatusBadRequest)
		return
	}

	// TODO: This can probably be refactored to be cleaner
	// If 2FA is disabled, finish authenticating the user
	if user.TotpKey == "" {
		OAuth2.StoreTokenAndRedirect(w, jwtString, "dashboard")
		return
	} else {
		// Supply temporary MFA auth token
		name := user.FirstName + " " + user.LastName
		permissions := []Models.Permission{Permissions.MFA_CODE_R}
		mfaJwtString, err := OAuth2.NewJwt(name, user.Id, user.UserGroup, permissions)
		if err != nil {
			Jsend.Error(req.Context(), w, "mfa validation failed", http.StatusBadRequest)
			return
		}
		OAuth2.StoreTokenAndRedirect(w, mfaJwtString, "validate-mfa")
	}
}

func deleteLogout(w http.ResponseWriter, _ *http.Request) {
	OAuth2.DeleteTokenAndRedirect(w, "login")
}
