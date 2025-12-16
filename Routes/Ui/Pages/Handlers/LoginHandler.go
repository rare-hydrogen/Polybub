package Handlers

import (
	"Polybub/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getLogin(w, req)
	}
}

func getLogin(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Pages/Templates/login.html"
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
