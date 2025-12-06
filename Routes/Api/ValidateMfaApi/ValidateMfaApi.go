package ValidateMfaApi

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Services"
	"Polybub/Jsend"
	"Polybub/Utilities/Permissions"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		post(w, req)
	case http.MethodPut:
		put(w, req)
	default:
		Jsend.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func post(w http.ResponseWriter, req *http.Request) {
	status, err := OAuth2.JwtPermitRequest(req, Permissions.MFA_CODE_R, nil)
	if err != nil {
		Jsend.Error(w, err.Error(), status)
		return
	}

	// Users logging in
	code := req.Header.Get("Code")
	mfaTokenString, err := OAuth2.GetTokenStringFromHeader(req)
	if err != nil {
		Jsend.Error(w, "login failed", http.StatusBadRequest)
		return
	}

	claims, err := OAuth2.GetClaimsFromTokenString(mfaTokenString)
	if err != nil {
		Jsend.Error(w, "login failed", http.StatusBadRequest)
		return
	}
	userId := claims.Subject

	_, tokenString, err := Services.MfaLogin(userId, code)
	if err != nil {
		Jsend.Error(w, "login failed", http.StatusBadRequest)
		return
	}

	OAuth2.StoreTokenAndRedirect(w, tokenString, "dashboard")
}

func put(w http.ResponseWriter, req *http.Request) {
	status, err := OAuth2.JwtPermitRequest(req, Permissions.MFA_CODE_CRU, nil)
	if err != nil {
		Jsend.Error(w, err.Error(), status)
		return
	}

	// Users updating their MFA TOTP key
	code := req.Header.Get("Code")
	key := req.Header.Get("Key")
	mfaTokenString, err := OAuth2.GetTokenStringFromHeader(req)
	if err != nil {
		Jsend.Error(w, "verification failed", http.StatusBadRequest)
		return
	}

	claims, err := OAuth2.GetClaimsFromTokenString(mfaTokenString)
	if err != nil {
		Jsend.Error(w, "verification failed", http.StatusBadRequest)
		return
	}
	userId := claims.Subject

	_, _, err = Services.MfaUpdate(userId, key, code)
	if err != nil {
		Jsend.Error(w, "verification failed", http.StatusBadRequest)
		return
	}

	Jsend.Success(w, "MFA Updated Successfully")
}
