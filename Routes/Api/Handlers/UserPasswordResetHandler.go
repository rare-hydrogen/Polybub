package Handlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Services"
	"Polybub/Data/Validators"
	"Polybub/Routes/Jsend"
	"errors"
	"net/http"
	"strconv"
)

func UserPasswordResetHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		postUserPasswordReset(w, req)
	case http.MethodPut:
		putUserPasswordReset(w, req)
	default:
		Jsend.MethodNotAllowed(w, req)
		return
	}
}

func postUserPasswordReset(w http.ResponseWriter, req *http.Request) {
	givenEmail := req.Header.Get("email")
	if givenEmail == "" {
		Jsend.Error(w, req, errors.New("invalid email"), "invalid email.", http.StatusBadRequest)
		return
	}

	userId, err := Services.GetIdByEmail(req.Context(), givenEmail)
	if err == nil {
		// If there is a real user:
		err2 := Services.AddResetKeyThenDeleteOthers(req.Context(), userId, givenEmail)
		if err2 != nil {
			// TODO: Continue silently handling the common error, but address the other error types
		}
	}

	Jsend.Success(w, req, nil)
	OAuth2.DeleteTokenAndRedirect(w, "/forgot-password?confirm")
}

func putUserPasswordReset(w http.ResponseWriter, req *http.Request) {
	queryId := req.Header.Get("id")
	givenUserId, err := strconv.Atoi(queryId)
	if err != nil {
		Jsend.Error(w, req, err, "invalid request", http.StatusBadRequest)
		return
	}

	givenKey := req.Header.Get("key")
	newPassword := req.Header.Get("newPassword")
	checkPassword := req.Header.Get("checkPassword")

	err = Validators.UserPassword(newPassword, checkPassword)
	if err != nil {
		Jsend.Error(w, req, err, err.Error(), http.StatusBadRequest)
		return
	}

	actualKey, err := Services.GetResetKey(req.Context(), int32(givenUserId))
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	if givenKey != actualKey {
		Jsend.Error(w, req, err, "invalid or expired key.", http.StatusBadRequest)
		return
	}

	err = Services.DeleteAllResetKeys(req.Context(), int32(givenUserId))
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	err = Services.UpdatePasswordAndSalt(req.Context(), int32(givenUserId), newPassword)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	Jsend.Success(w, req, nil)
	OAuth2.DeleteTokenAndRedirect(w, "/login")
}
