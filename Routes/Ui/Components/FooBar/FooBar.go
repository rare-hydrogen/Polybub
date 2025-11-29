package FooBar

import (
	"Polybub/Data/Models"
	"Polybub/Data/Services"
	"Polybub/Data/Validators"
	"Polybub/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Rhymond/go-money"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		get(w, req)
	case "POST":
		post(w, req)
	default:
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

func post(w http.ResponseWriter, req *http.Request) {
	buf, err := io.ReadAll(req.Body)
	if err != nil {
		Jsend.Error(w, "Something went wrong!", http.StatusInternalServerError)
		return
	}

	var dto Models.FooBar
	err = json.Unmarshal(buf, &dto)
	if err != nil {
		Jsend.Error(w, "missing or invalid data", http.StatusBadRequest)
		return
	}
	err = Validators.FooBar(dto)
	if err != nil {
		Jsend.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	obj, err := Services.CreateFooBar(dto)
	if err != nil {
		Jsend.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	Jsend.Success(w, obj)
}
