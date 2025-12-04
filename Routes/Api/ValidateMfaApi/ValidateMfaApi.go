package ValidateMfaApi

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Services"
	"Polybub/Jsend"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		post(w, req)
	default:
		Jsend.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func post(w http.ResponseWriter, req *http.Request) {
	// 1. get code and mfaToken
	// 2. get the user from the mfaToken
	// 3. validate the code against the user's TOTP KEY
	// 4. create a new full jwtToken
	// 5. redirect the user to the dashboard

	// 1
	code := req.Header.Get("Code")
	mfaTokenString, err := OAuth2.GetTokenStringFromHeader(req)
	if err != nil {
		Jsend.Error(w, "login failed", http.StatusBadRequest)
		return
	}

	// 2
	claims, err := OAuth2.GetClaimsFromTokenString(mfaTokenString)
	if err != nil {
		Jsend.Error(w, "login failed", http.StatusBadRequest)
		return
	}
	userId := claims.Subject

	// 3 and 4
	_, tokenString, err := Services.MfaLogin(userId, code)
	if err != nil {
		Jsend.Error(w, "login failed", http.StatusBadRequest)
		return
	}

	// 5
	OAuth2.StoreTokenAndRedirect(w, tokenString, "dashboard")
}
