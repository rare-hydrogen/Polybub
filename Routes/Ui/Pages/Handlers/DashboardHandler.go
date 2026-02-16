package Handlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/TemplateEmbeds"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"net/http"
)

// TODO: Fix this
type dddd struct {
	Name string
}

func DashboardHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getDashboard(w, req)
	default:
		Jsend.MethodNotAllowed(w, req)
		return
	}
}

func getDashboard(w http.ResponseWriter, req *http.Request) {
	b, _ := TemplateEmbeds.PageEmbeds.ReadFile("PageEmbeds/dashboard.html")
	tokenString, err := OAuth2.GetTokenStringFromHeader(req)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	claims, err := OAuth2.GetClaimsFromTokenString(tokenString)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	data := dddd{
		Name: claims.Name,
	}

	body, err := GlobalWrapper.GetSafeHtml(b, data)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	wrappedBody, err := GlobalWrapper.GetWrappedTemplate(body)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	Jsend.Ui(req.Context(), w, wrappedBody)
}
