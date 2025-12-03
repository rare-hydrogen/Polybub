package UserPasswordResetApi

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Services"
	"Polybub/Data/Validators"
	"Polybub/Jsend"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		requestReset(w, req)
	case http.MethodPut:
		attemptReset(w, req)
	default:
		Jsend.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func requestReset(w http.ResponseWriter, req *http.Request) {
	givenEmail := req.Header.Get("email")
	if givenEmail == "" {
		Jsend.Error(w, "Invalid email.", http.StatusBadRequest)
		return
	}

	userId, err := Services.GetIdByEmail(givenEmail)
	if err == nil {
		// If there is a real user:
		err2 := Services.AddResetKeyThenDeleteOthers(userId)
		if err2 != nil {
			// We don't show errors to avoid giving away user identities
		}
	}

	Jsend.Success(w, nil)
	OAuth2.DeleteTokenAndRedirect(w, "/forgot-password?confirm")
}

func attemptReset(w http.ResponseWriter, req *http.Request) {
	queryId := req.Header.Get("id")
	givenUserId, err := strconv.Atoi(queryId)
	if err != nil {
		Jsend.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}

	givenKey := req.Header.Get("key")
	newPassword := req.Header.Get("newPassword")
	checkPassword := req.Header.Get("checkPassword")

	err = Validators.UserPassword(newPassword, checkPassword)
	if err != nil {
		Jsend.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actualKey, err := Services.GetResetKey(int32(givenUserId))
	if err != nil {
		Jsend.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	if givenKey != actualKey {
		Jsend.Error(w, "Invalid or expired key.", http.StatusBadRequest)
		return
	}

	err = Services.DeleteAllResetKeys(int32(givenUserId))
	if err != nil {
		Jsend.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	err = Services.UpdatePasswordAndSalt(int32(givenUserId), newPassword)
	if err != nil {
		Jsend.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	Jsend.Success(w, nil)
	OAuth2.DeleteTokenAndRedirect(w, "/login")
}
