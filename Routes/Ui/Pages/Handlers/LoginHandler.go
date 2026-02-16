package Handlers

import (
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/TemplateEmbeds"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getLogin(w, req)
	default:
		Jsend.MethodNotAllowed(w, req)
		return
	}
}

func getLogin(w http.ResponseWriter, req *http.Request) {
	b, err := TemplateEmbeds.PageEmbeds.ReadFile("PageEmbeds/login.html")
	data := ""
	body, err := GlobalWrapper.GetSafeHtml(b, data)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	wrappedBody, err := GlobalWrapper.GetPublicTemplate(body)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	Jsend.Ui(req.Context(), w, wrappedBody)
}
