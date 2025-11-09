package Foobar

import (
	"Polybub/Data/Models"
	"Polybub/Data/Services"
	"Polybub/Jsend"
	"encoding/json"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if req.URL.Query().Has("id") {
			getSingle(w, req)
		} else {
			getMany(w, req)
		}
	case http.MethodPost:
		post(w, req)
	case http.MethodPatch:
		patch(w, req)
	case http.MethodDelete:
		delete(w, req)
	}
}

func getSingle(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)
	if err != nil {
		Jsend.Error(w, err.Error())
		return
	}

	d, err := Services.ReadSingleFooBar(int32(id))
	if err != nil {
		Jsend.Error(w, err.Error())
		return
	}

	Jsend.Success(w, d)
}

func getMany(w http.ResponseWriter, req *http.Request) {
	d, err := Services.ReadManyFooBar()
	if err != nil {
		Jsend.Error(w, err.Error())
		return
	}

	Jsend.Success(w, d)
}

func post(w http.ResponseWriter, req *http.Request) {
	var dto Models.FooBar
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		Jsend.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d, err := Services.CreateFooBar(dto)
	if err != nil {
		Jsend.Error(w, err.Error())
		return
	}

	Jsend.Success(w, d)
}

func patch(w http.ResponseWriter, req *http.Request) {
	var dto Models.FooBar
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		Jsend.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d, err := Services.UpdateFooBar(dto)
	if err != nil {
		Jsend.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Jsend.Success(w, d)
}

func delete(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)
	if err != nil {
		Jsend.Error(w, err.Error())
		return
	}

	err = Services.SoftDeleteFooBar(int32(id))
	if err != nil {
		Jsend.Error(w, err.Error())
		return
	}

	Jsend.Success(w, nil)
}
