package Handlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Auth/Totp"
	"Polybub/Data/Services"
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/TemplateEmbeds"
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
	switch req.Method {
	case "GET":
		if req.URL.Query().Has("form") {
			getSetupMfaForm(w, req)
		} else {
			getSetupMfa(w, req)
		}
	default:
		Jsend.MethodNotAllowed(w, req)
		return
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
		Jsend.Error(req.Context(), w, err, "invalid cookies", http.StatusBadRequest)
		return
	}

	b, _ := TemplateEmbeds.PageEmbeds.ReadFile("PageEmbeds/setup-mfa.html")
	body, err := GlobalWrapper.GetSafeHtml(b, v)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	parsedHtml, err := GlobalWrapper.GetWrappedTemplate(body)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	Jsend.Ui(req.Context(), w, parsedHtml)
}

func getSetupMfaForm(w http.ResponseWriter, req *http.Request) {
	v, err := getVariantMfaData(req)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	b, _ := TemplateEmbeds.PageEmbeds.ReadFile("PageEmbeds/send-code-to-check.html")
	body, err := GlobalWrapper.GetSafeHtml(b, v)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	Jsend.Ui(req.Context(), w, body)
}
