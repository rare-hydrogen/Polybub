package Handlers

import (
	"Polybub/Data/Models"
	"Polybub/Data/Services"
	"Polybub/Data/Validators"
	"Polybub/Routes/Jsend"
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
		Jsend.Error(req.Context(), w, "not allowed", http.StatusMethodNotAllowed)
	}
}

func getFooBar(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Components/Templates/foobar.html"
	data, err := Services.ReadLatestFooBar()
	if err != nil {
		data = Models.FooBar{
			Id:       1,
			Name:     "Foonius Barius",
			Type:     "Type 1",
			Amount:   100,
			Currency: "USD",
		}
	}

	htmlText, err := GlobalWrapper.GetSafeHtml(path, data)
	if err != nil {
		Jsend.Error(req.Context(), w, "error reading template", http.StatusBadRequest)
		return
	}
	Jsend.Ui(req.Context(), w, htmlText)
}

func postFooBar(w http.ResponseWriter, req *http.Request) {
	buf, err := io.ReadAll(req.Body)
	if err != nil {
		Jsend.Error(req.Context(), w, "Something went wrong!", http.StatusInternalServerError)
		return
	}

	var dto Models.FooBar
	err = json.Unmarshal(buf, &dto)
	if err != nil {
		Jsend.Error(req.Context(), w, "missing or invalid data", http.StatusBadRequest)
		return
	}
	err = Validators.FooBar(dto)
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error(), http.StatusBadRequest)
		return
	}

	obj, err := Services.CreateFooBar(dto)
	if err != nil {
		Jsend.Error(req.Context(), w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	Jsend.Success(req.Context(), w, obj)
}
