package ComponentHandlers

import (
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/Wrappers"
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
	data := ""
	Wrappers.TemplateHandler(w, req, "comp.validate-mfa", data)
}
