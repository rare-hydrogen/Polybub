package ForgotPassword

import (
	"Polybub/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		path := "Routes/Ui/Pages/ForgotPassword/forgot-password.html"
		data := ""
		body, err := GlobalWrapper.GetSafeHtml(path, data)
		if err != nil {
			Jsend.Error(w, "Error reading template", http.StatusInternalServerError)
			return
		}

		wrappedBody, err := GlobalWrapper.GetWrappedTemplate(body)
		if err != nil {
			Jsend.Error(w, "Error wrapping template", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, wrappedBody)
	}
}
