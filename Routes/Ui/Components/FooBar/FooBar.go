package FooBar

import (
	"Polybub/Data/Models"
	"Polybub/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"fmt"
	"net/http"

	"github.com/Rhymond/go-money"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		get(w, req)
	} else {
		Jsend.Error(w, "not allowed", http.StatusMethodNotAllowed)
	}
}

func get(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Components/FooBar/foobar.html"
	data := Models.FooBar{
		Id:       1,
		Name:     "Foonius Barius",
		Type:     "Type 1",
		Amount:   100,
		Currency: *money.GetCurrency("USD"),
	}

	htmlText, err := GlobalWrapper.GetSafeHtml(path, data)
	if err != nil {
		Jsend.Error(w, "error reading template", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, htmlText)
}
