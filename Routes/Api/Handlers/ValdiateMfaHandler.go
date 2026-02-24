package Handlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Services"
	"Polybub/Routes/Jsend"
	"Polybub/Utilities/Permissions"
	"net/http"
)

func ValdiateMfaHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		postValdiateMfa(w, req)
	case http.MethodPut:
		putValdiateMfa(w, req)
	default:
		Jsend.MethodNotAllowed(w, req)
		return
	}
}

func postValdiateMfa(w http.ResponseWriter, req *http.Request) {
	status, err := OAuth2.JwtPermitRequest(req, Permissions.MFA_CODE_R, nil)
	if err != nil {
		Jsend.Error(w, req, err, err.Error(), status)
		return
	}

	// Users logging in
	code := req.Header.Get("Code")
	mfaTokenString, err := OAuth2.GetTokenStringFromHeader(req)
	if err != nil {
		Jsend.Error(w, req, err, "login failed", http.StatusBadRequest)
		return
	}

	claims, err := OAuth2.GetClaimsFromTokenString(mfaTokenString)
	if err != nil {
		Jsend.Error(w, req, err, "login failed", http.StatusBadRequest)
		return
	}
	userId := claims.Subject

	_, tokenString, err := Services.MfaLogin(req.Context(), userId, code)
	if err != nil {
		Jsend.Error(w, req, err, "login failed", http.StatusBadRequest)
		return
	}

	OAuth2.StoreTokenAndRedirect(w, tokenString, "dashboard")
}

func putValdiateMfa(w http.ResponseWriter, req *http.Request) {
	status, err := OAuth2.JwtPermitRequest(req, Permissions.MFA_CODE_CRU, nil)
	if err != nil {
		Jsend.Error(w, req, err, err.Error(), status)
		return
	}

	// Users updating their MFA TOTP key
	code := req.Header.Get("Code")
	key := req.Header.Get("Key")
	mfaTokenString, err := OAuth2.GetTokenStringFromHeader(req)
	if err != nil {
		Jsend.Error(w, req, err, "verification failed", http.StatusBadRequest)
		return
	}

	claims, err := OAuth2.GetClaimsFromTokenString(mfaTokenString)
	if err != nil {
		Jsend.Error(w, req, err, "verification failed", http.StatusBadRequest)
		return
	}
	userId := claims.Subject

	_, _, err = Services.MfaUpdate(req.Context(), userId, key, code)
	if err != nil {
		Jsend.Error(w, req, err, "verification failed", http.StatusBadRequest)
		return
	}

	Jsend.Success(w, req, "MFA Updated Successfully")
}
