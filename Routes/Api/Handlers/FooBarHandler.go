package Handlers

import (
	"Polybub/Data/Models"
	"Polybub/Data/Services"
	"Polybub/Jsend"
	"encoding/json"
	"net/http"
	"strconv"
)

func FooBarHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if req.URL.Query().Has("id") {
			getSingleFooBar(w, req)
		} else {
			getManyFooBar(w, req)
		}
	case http.MethodPost:
		postFooBar(w, req)
	case http.MethodPatch:
		patchFooBar(w, req)
	case http.MethodDelete:
		deleteFooBar(w, req)
	}
}

func getSingleFooBar(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error())
		return
	}

	d, err := Services.ReadSingleFooBar(int32(id))
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error())
		return
	}

	Jsend.Success(req.Context(), w, d)
}

func getManyFooBar(w http.ResponseWriter, req *http.Request) {
	d, err := Services.ReadManyFooBar()
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error())
		return
	}

	Jsend.Success(req.Context(), w, d)
}

func postFooBar(w http.ResponseWriter, req *http.Request) {
	var dto Models.FooBar
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error(), http.StatusBadRequest)
		return
	}

	d, err := Services.CreateFooBar(dto)
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error())
		return
	}

	Jsend.Success(req.Context(), w, d)
}

func patchFooBar(w http.ResponseWriter, req *http.Request) {
	var dto Models.FooBar
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error(), http.StatusBadRequest)
		return
	}

	d, err := Services.UpdateFooBar(dto)
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error(), http.StatusInternalServerError)
		return
	}

	Jsend.Success(req.Context(), w, d)
}

func deleteFooBar(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error())
		return
	}

	err = Services.SoftDeleteFooBar(int32(id))
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error())
		return
	}

	Jsend.Success(req.Context(), w, nil)
}
