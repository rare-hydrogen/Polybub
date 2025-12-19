package Handlers

import (
	"Polybub/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"fmt"
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
		Jsend.Error(w, "Error reading template", http.StatusInternalServerError)
		return
	}

	wrappedBody, err := GlobalWrapper.GetPublicTemplate(body)
	if err != nil {
		Jsend.Error(w, "Error wrapping template", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, wrappedBody)
}
