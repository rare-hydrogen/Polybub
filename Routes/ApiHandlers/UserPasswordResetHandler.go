package ApiHandlers

import (
	"Polybub/Data/Services"
	"Polybub/Jsend"
	"net/http"
	"strconv"
)

func UserPasswordResetHandler(w http.ResponseWriter, req *http.Request) {
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
	var username = "ts"

	userId, err := Services.GetIdByUsername(username)
	if err != nil {
		Jsend.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	err = Services.AddResetKeyThenDeleteOthers(userId)
	if err != nil {
		Jsend.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	Jsend.Success(w, nil)
}

func attemptReset(w http.ResponseWriter, req *http.Request) {
	queryId := req.URL.Query().Get("id")
	givenUserId, err := strconv.Atoi(queryId)
	if err != nil {
		Jsend.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}
	givenKey := req.URL.Query().Get("key")
	newPassword := "" // TODO: Replace this

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
}
