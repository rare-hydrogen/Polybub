package Handlers

import (
	"Polybub/Data/Models"
	"Polybub/Data/Services"
	"Polybub/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		getCreateUser(w, req)
	}
	if req.Method == "POST" {
		postCreateUser(w, req)
	}
}

func getCreateUser(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Pages/Templates/create-user.html"
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

func postCreateUser(w http.ResponseWriter, req *http.Request) {
	var dto Models.User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&dto)
	if err != nil {
		Jsend.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d, err := Services.CreateUser(dto)
	if err != nil {
		Jsend.Error(w, err.Error())
		return
	}

	Services.UpdatePasswordAndSalt(d.Id, dto.Password)

	Jsend.Success(w, d)
}
