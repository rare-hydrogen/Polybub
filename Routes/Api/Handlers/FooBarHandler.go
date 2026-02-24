package Handlers

import (
	"Polybub/Data/Models"
	"Polybub/Data/Services"
	"Polybub/Routes/Jsend"
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
	default:
		Jsend.MethodNotAllowed(w, req)
		return
	}
}

func getSingleFooBar(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)
	if err != nil {
		Jsend.Error(w, req, err, "missing id", http.StatusBadRequest)
		return
	}

	d, err := Services.ReadSingleFooBar(req.Context(), int32(id))
	if err != nil {
		Jsend.Error(w, req, err, "not found", http.StatusNotFound)
		return
	}

	Jsend.Success(w, req, d)
}

func getManyFooBar(w http.ResponseWriter, req *http.Request) {
	d, err := Services.ReadManyFooBar(req.Context())
	if err != nil {
		Jsend.InternalServerError(w, req, err)
		return
	}

	Jsend.Success(w, req, d)
}

func postFooBar(w http.ResponseWriter, req *http.Request) {
	var dto Models.FooBar
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		Jsend.Error(w, req, err, "invalid request", http.StatusBadRequest)
		return
	}

	d, err := Services.CreateFooBar(req.Context(), dto)
	if err != nil {
		Jsend.Error(w, req, err, "create failed", http.StatusBadRequest)
		return
	}

	Jsend.Success(w, req, d)
}

func patchFooBar(w http.ResponseWriter, req *http.Request) {
	var dto Models.FooBar
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		Jsend.Error(w, req, err, "update failed", http.StatusBadRequest)
		return
	}

	d, err := Services.UpdateFooBar(req.Context(), dto)
	if err != nil {
		Jsend.Error(w, req, err, "update failed", http.StatusBadRequest)
		return
	}

	Jsend.Success(w, req, d)
}

func deleteFooBar(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseInt(req.URL.Query().Get("id"), 10, 64)
	if err != nil {
		Jsend.Error(w, req, err, "missing id", http.StatusBadRequest)
		return
	}

	err = Services.SoftDeleteFooBar(req.Context(), int32(id))
	if err != nil {
		Jsend.Error(w, req, err, "delete failed", http.StatusBadRequest)
		return
	}

	Jsend.Success(w, req, nil)
}
