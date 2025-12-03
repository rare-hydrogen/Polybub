package ForgotPassword

import (
	"Polybub/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		if req.URL.Query().Has("confirm") {
			getConfirm(w, req)
		} else if req.URL.Query().Has("key") {
			getAttempt(w, req)
		} else {
			get(w, req)
		}
	}
}

func get(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Pages/ForgotPassword/forgot-password.html"
	data := ""
	body, err := GlobalWrapper.GetSafeHtml(path, data)
	if err != nil {
		Jsend.Error(w, "Error reading template", http.StatusInternalServerError)
		return
	}

	wrappedBody, err := GlobalWrapper.GetWrappedTemplate(body)
	if err != nil {
		Jsend.Error(w, "Error wrapping template", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, wrappedBody)
}

func getConfirm(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Pages/ForgotPassword/forgot-password-confirm.html"
	data := ""
	body, err := GlobalWrapper.GetSafeHtml(path, data)
	if err != nil {
		Jsend.Error(w, "Error reading template", http.StatusInternalServerError)
		return
	}

	wrappedBody, err := GlobalWrapper.GetWrappedTemplate(body)
	if err != nil {
		Jsend.Error(w, "Error wrapping template", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, wrappedBody)
}

func getAttempt(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Pages/ForgotPassword/forgot-password-attempt.html"
	data := ""
	body, err := GlobalWrapper.GetSafeHtml(path, data)
	if err != nil {
		Jsend.Error(w, "Error reading template", http.StatusInternalServerError)
		return
	}

	wrappedBody, err := GlobalWrapper.GetWrappedTemplate(body)
	if err != nil {
		Jsend.Error(w, "Error wrapping template", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, wrappedBody)
}
