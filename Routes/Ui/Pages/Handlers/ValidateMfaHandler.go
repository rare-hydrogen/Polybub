package Handlers

import (
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"net/http"
)

func ValidateMfaHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getValidateMfa(w, req)
	}
}

func getValidateMfa(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Pages/Templates/validate-mfa.html"
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
