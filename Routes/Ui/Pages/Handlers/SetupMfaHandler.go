package Handlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Auth/Totp"
	"Polybub/Data/Services"
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"encoding/base64"
	"html/template"
	"net/http"
)

type variantMfaData struct {
	Image template.URL
	Key   string
}

func SetupMfaHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		if req.URL.Query().Has("form") {
			getSetupMfaForm(w, req)
		} else {
			getSetupMfa(w, req)
		}
	}
}

func getVariantMfaData(req *http.Request) (variantMfaData, error) {
	// Get claims
	tokenString, err := OAuth2.GetTokenStringFromHeader(req)
	if err != nil {
		return variantMfaData{}, err
	}
	claims, err := OAuth2.GetClaimsFromTokenString(tokenString)
	if err != nil {
		return variantMfaData{}, err
	}

	// Use claims
	user, err := Services.UnsafeReadSingleUser(req.Context(), claims.Subject)
	if err != nil {
		return variantMfaData{}, err
	}
	imgBuf, key, err := Totp.BeginTotp(user)
	if err != nil {
		return variantMfaData{}, err
	}

	encoded := base64.StdEncoding.EncodeToString(imgBuf.Bytes())
	dataURI := "data:image/png;base64," + encoded

	return variantMfaData{
		Image: template.URL(dataURI),
		Key:   key,
	}, nil
}

func getSetupMfa(w http.ResponseWriter, req *http.Request) {
	v, err := getVariantMfaData(req)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error reading cookies", http.StatusBadRequest)
		return
	}

	path := "Routes/Ui/Pages/Templates/setup-mfa.html"
	body, err := GlobalWrapper.GetSafeHtml(path, v)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error reading template", http.StatusInternalServerError)
		return
	}

	parsedHtml, err := GlobalWrapper.GetWrappedTemplate(body)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error wrapping template", http.StatusInternalServerError)
		return
	}

	Jsend.Ui(req.Context(), w, parsedHtml)
}

func getSetupMfaForm(w http.ResponseWriter, req *http.Request) {
	v, err := getVariantMfaData(req)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error reading cookies", http.StatusBadRequest)
		return
	}

	path := "Routes/Ui/Pages/Templates/send-code-to-check.html"
	body, err := GlobalWrapper.GetSafeHtml(path, v)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error reading template", http.StatusInternalServerError)
		return
	}

	Jsend.Ui(req.Context(), w, body)
}
