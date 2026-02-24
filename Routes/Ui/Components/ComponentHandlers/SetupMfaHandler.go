package ComponentHandlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Auth/Totp"
	"Polybub/Data/Services"
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/Wrappers"
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

func retrieveVariantMfaData(req *http.Request) (variantMfaData, error) {
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
	v, err := retrieveVariantMfaData(req)
	if err != nil {
		Jsend.Error(w, req, err, "invalid cookies", http.StatusBadRequest)
		return
	}

	Wrappers.TemplateHandler(w, req, "comp.setup-mfa", v)
}

func getSetupMfaForm(w http.ResponseWriter, req *http.Request) {
	v, err := retrieveVariantMfaData(req)
	if err != nil {
		Jsend.Error(w, req, err, "invalid cookies", http.StatusBadRequest)
		return
	}

	Wrappers.TemplateHandler(w, req, "comp.setup-mfa-form", v)
}
