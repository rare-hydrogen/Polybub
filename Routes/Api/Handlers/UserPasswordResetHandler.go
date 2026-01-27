package Handlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Services"
	"Polybub/Data/Validators"
	"Polybub/Jsend"
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
		Jsend.Error(req.Context(), w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func postUserPasswordReset(w http.ResponseWriter, req *http.Request) {
	givenEmail := req.Header.Get("email")
	if givenEmail == "" {
		Jsend.Error(req.Context(), w, "Invalid email.", http.StatusBadRequest)
		return
	}

	userId, err := Services.GetIdByEmail(givenEmail)
	if err == nil {
		// If there is a real user:
		err2 := Services.AddResetKeyThenDeleteOthers(userId, givenEmail)
		if err2 != nil {
			// We don't show errors to avoid giving away user identities
		}
	}

	Jsend.Success(req.Context(), w, nil)
	OAuth2.DeleteTokenAndRedirect(w, "/forgot-password?confirm")
}

func putUserPasswordReset(w http.ResponseWriter, req *http.Request) {
	queryId := req.Header.Get("id")
	givenUserId, err := strconv.Atoi(queryId)
	if err != nil {
		Jsend.Error(req.Context(), w, "Invalid request.", http.StatusBadRequest)
		return
	}

	givenKey := req.Header.Get("key")
	newPassword := req.Header.Get("newPassword")
	checkPassword := req.Header.Get("checkPassword")

	err = Validators.UserPassword(newPassword, checkPassword)
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error(), http.StatusBadRequest)
		return
	}

	actualKey, err := Services.GetResetKey(int32(givenUserId))
	if err != nil {
		Jsend.Error(req.Context(), w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	if givenKey != actualKey {
		Jsend.Error(req.Context(), w, "Invalid or expired key.", http.StatusBadRequest)
		return
	}

	err = Services.DeleteAllResetKeys(int32(givenUserId))
	if err != nil {
		Jsend.Error(req.Context(), w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	err = Services.UpdatePasswordAndSalt(int32(givenUserId), newPassword)
	if err != nil {
		Jsend.Error(req.Context(), w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	Jsend.Success(req.Context(), w, nil)
	OAuth2.DeleteTokenAndRedirect(w, "/login")
}
