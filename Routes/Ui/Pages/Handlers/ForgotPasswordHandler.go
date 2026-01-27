package Handlers

import (
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"net/http"
)

func ForgotPasswordHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		if req.URL.Query().Has("confirm") {
			getForgotPasswordConfirm(w, req)
		} else if req.URL.Query().Has("key") {
			getForgotPasswordAttempt(w, req)
		} else {
			getForgotPassword(w, req)
		}
	}
}

func getForgotPassword(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Pages/Templates/forgot-password.html"
	data := ""
	body, err := GlobalWrapper.GetSafeHtml(path, data)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error reading template", http.StatusInternalServerError)
		return
	}

	wrappedBody, err := GlobalWrapper.GetPublicTemplate(body)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error wrapping template", http.StatusInternalServerError)
		return
	}

	Jsend.Ui(req.Context(), w, wrappedBody)
}

func getForgotPasswordConfirm(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Pages/Templates/forgot-password-confirm.html"
	data := ""
	body, err := GlobalWrapper.GetSafeHtml(path, data)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error reading template", http.StatusInternalServerError)
		return
	}

	wrappedBody, err := GlobalWrapper.GetPublicTemplate(body)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error wrapping template", http.StatusInternalServerError)
		return
	}

	Jsend.Ui(req.Context(), w, wrappedBody)
}

func getForgotPasswordAttempt(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Pages/Templates/forgot-password-attempt.html"
	data := ""
	body, err := GlobalWrapper.GetSafeHtml(path, data)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error reading template", http.StatusInternalServerError)
		return
	}

	wrappedBody, err := GlobalWrapper.GetPublicTemplate(body)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error wrapping template", http.StatusInternalServerError)
		return
	}

	Jsend.Ui(req.Context(), w, wrappedBody)
}
