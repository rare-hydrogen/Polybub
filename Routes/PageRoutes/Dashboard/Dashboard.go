package Dashboard

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Jsend"
	"Polybub/Routes/GlobalWrapper"
	"fmt"
	"net/http"
)

// TODO: Fix this
type dddd struct {
	Name string
}

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		path := "Routes/PageRoutes/Dashboard/dashboard.html"
		tokenString, err := OAuth2.GetTokenStringFromHeader(req)
		if err != nil {
			Jsend.Error(w, "Error reading token", http.StatusInternalServerError)
			return
		}

		claims, err := OAuth2.GetClaimsFromTokenString(tokenString)
		if err != nil {
			Jsend.Error(w, "Error reading token", http.StatusInternalServerError)
			return
		}

		data := dddd{
			Name: claims.Name,
		}

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
