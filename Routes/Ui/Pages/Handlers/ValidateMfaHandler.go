package Handlers

import (
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/TemplateEmbeds"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"net/http"
)

func ValidateMfaHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getValidateMfa(w, req)
	default:
		Jsend.MethodNotAllowed(w, req)
		return
	}
}

func getValidateMfa(w http.ResponseWriter, req *http.Request) {
	b, _ := TemplateEmbeds.PageEmbeds.ReadFile("PageEmbeds/validate-mfa.html")
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
