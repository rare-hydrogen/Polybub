package Handlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"net/http"
)

// TODO: Fix this
type dddd struct {
	Name string
}

func DashboardHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getDashboard(w, req)
	}
}

func getDashboard(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Pages/Templates/dashboard.html"
	tokenString, err := OAuth2.GetTokenStringFromHeader(req)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error reading token", http.StatusInternalServerError)
		return
	}

	claims, err := OAuth2.GetClaimsFromTokenString(tokenString)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error reading token", http.StatusInternalServerError)
		return
	}

	data := dddd{
		Name: claims.Name,
	}

	body, err := GlobalWrapper.GetSafeHtml(path, data)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error reading template", http.StatusInternalServerError)
		return
	}

	wrappedBody, err := GlobalWrapper.GetWrappedTemplate(body)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error wrapping template", http.StatusInternalServerError)
		return
	}

	Jsend.Ui(req.Context(), w, wrappedBody)
}
