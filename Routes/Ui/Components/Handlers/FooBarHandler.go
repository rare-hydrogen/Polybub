package Handlers

import (
	"Polybub/Data/Models"
	"Polybub/Data/Services"
	"Polybub/Data/Validators"
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/TemplateEmbeds"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"encoding/json"
	"io"
	"net/http"
)

func FooBarHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getFooBar(w, req)
	case "POST":
		postFooBar(w, req)
	default:
		Jsend.MethodNotAllowed(w, req)
		return
	}
}

func getFooBar(w http.ResponseWriter, req *http.Request) {
	b, _ := TemplateEmbeds.ComponentEmbeds.ReadFile("ComponentEmbeds/foobar.html")
	data, err := Services.ReadLatestFooBar(req.Context())
	if err != nil {
		data = Models.FooBar{
			Id:       1,
			Name:     "Foonius Barius",
			Type:     "Type 1",
			Amount:   100,
			Currency: "USD",
		}
	}

	htmlText, err := GlobalWrapper.GetSafeHtml(b, data)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}
	Jsend.Ui(req.Context(), w, htmlText)
}

func postFooBar(w http.ResponseWriter, req *http.Request) {
	buf, err := io.ReadAll(req.Body)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	var dto Models.FooBar
	err = json.Unmarshal(buf, &dto)
	if err != nil {
		Jsend.Error(req.Context(), w, err, "missing or invalid data", http.StatusBadRequest)
		return
	}
	err = Validators.FooBar(dto)
	if err != nil {
		Jsend.Error(req.Context(), w, err, err.Error(), http.StatusBadRequest)
		return
	}

	obj, err := Services.CreateFooBar(req.Context(), dto)
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	Jsend.Success(req.Context(), w, obj)
}
