package ComponentHandlers

import (
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/Wrappers"
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
	data := ""
	Wrappers.TemplateHandler(w, req, "comp.forgot-password", data)
}

func getForgotPasswordConfirm(w http.ResponseWriter, req *http.Request) {
	data := ""
	Wrappers.TemplateHandler(w, req, "comp.forgot-password-confirm", data)
}

func getForgotPasswordAttempt(w http.ResponseWriter, req *http.Request) {
	data := ""
	Wrappers.TemplateHandler(w, req, "comp.forgot-password-attempt", data)
}
