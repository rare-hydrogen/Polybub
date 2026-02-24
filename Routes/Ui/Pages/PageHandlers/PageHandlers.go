package PageHandlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/Wrappers"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	data := ""
	Wrappers.TemplateHandler(w, req, "page.login", data)
}

func DashboardHandler(w http.ResponseWriter, req *http.Request) {
	// Get the user's claims
	claims, err := OAuth2.GetClaimsFromRequest(w, req)
	if err != nil {
		Jsend.Error(w, req, err, "invalid request", http.StatusUnauthorized)
		return
	}

	// Assign data
	data := map[string]string{
		"Name": claims.Name,
	}

	// Feed into page
	Wrappers.TemplateHandler(w, req, "page.dashboard", data)
}

func SetupMfaHandler(w http.ResponseWriter, req *http.Request) {
	Wrappers.TemplateHandler(w, req, "page.setup-mfa", nil)
}

func ForgotPasswordHandler(w http.ResponseWriter, req *http.Request) {
	Wrappers.TemplateHandler(w, req, "page.forgot-password", nil)
}

func ValidateMfaHandler(w http.ResponseWriter, req *http.Request) {
	Wrappers.TemplateHandler(w, req, "page.validate-mfa", nil)
}

func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	Wrappers.TemplateHandler(w, req, "page.create-user", nil)
}
