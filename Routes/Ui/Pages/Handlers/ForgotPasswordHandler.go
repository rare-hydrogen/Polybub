package Handlers

import (
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/TemplateEmbeds"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"net/http"
)

func ForgotPasswordHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if req.URL.Query().Has("confirm") {
			getForgotPasswordConfirm(w, req)
		} else if req.URL.Query().Has("key") {
			getForgotPasswordAttempt(w, req)
		} else {
			getForgotPassword(w, req)
		}
	default:
		Jsend.MethodNotAllowed(w, req)
		return
	}
}

func getForgotPassword(w http.ResponseWriter, req *http.Request) {
	b, _ := TemplateEmbeds.PageEmbeds.ReadFile("PageEmbeds/forgot-password.html")
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

func getForgotPasswordConfirm(w http.ResponseWriter, req *http.Request) {
	b, _ := TemplateEmbeds.PageEmbeds.ReadFile("PageEmbeds/forgot-password-confirm.html")
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

func getForgotPasswordAttempt(w http.ResponseWriter, req *http.Request) {
	b, _ := TemplateEmbeds.PageEmbeds.ReadFile("PageEmbeds/forgot-password-attempt.html")
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
