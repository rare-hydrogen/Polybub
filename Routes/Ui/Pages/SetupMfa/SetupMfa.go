package SetupMfa

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Auth/Totp"
	"Polybub/Data/Services"
	"Polybub/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
)

type variantMfaData struct {
	Image template.URL
	Key   string
}

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		if req.URL.Query().Has("form") {
			getForm(w, req)
		} else {
			get(w, req)
		}
	}
}

func get(w http.ResponseWriter, req *http.Request) {
	v, err := getData(req)
	if err != nil {
		Jsend.Error(w, "Error reading cookies", http.StatusBadRequest)
		return
	}

	path := "Routes/Ui/Pages/SetupMfa/setup-mfa.html"
	body, err := GlobalWrapper.GetSafeHtml(path, v)
	if err != nil {
		Jsend.Error(w, "Error reading template", http.StatusInternalServerError)
		return
	}

	parsedHtml, err := GlobalWrapper.GetWrappedTemplate(body)
	if err != nil {
		Jsend.Error(w, "Error wrapping template", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, parsedHtml)
}

func getData(req *http.Request) (variantMfaData, error) {
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
	user, err := Services.UnsafeReadSingleUser(claims.Subject)
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

func getForm(w http.ResponseWriter, req *http.Request) {
	v, err := getData(req)
	if err != nil {
		Jsend.Error(w, "Error reading cookies", http.StatusBadRequest)
		return
	}

	path := "Routes/Ui/Pages/SetupMfa/send-code-to-check.html"
	body, err := GlobalWrapper.GetSafeHtml(path, v)
	if err != nil {
		Jsend.Error(w, "Error reading template", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, body)
}
