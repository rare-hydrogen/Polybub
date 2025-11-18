package UI

import (
	"Polybub/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"fmt"
	"net/http"
)

func ForgotPasswordHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		get(w, req)
	} else {
		Jsend.Error(w, "not allowed", http.StatusMethodNotAllowed)
	}
}

func get(w http.ResponseWriter, req *http.Request) {
	// Frontend
	path := "Routes/Components/Dashboards/ProviderDashboard/provider-dashboard.html"
	htmlText, err := GlobalWrapper.GetSafeHtml(path, nil)
	if err != nil {
		Jsend.Error(w, "error reading template", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, htmlText)
}
