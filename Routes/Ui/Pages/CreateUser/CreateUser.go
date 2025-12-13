package CreateUser

import (
	"Polybub/Data/Models"
	"Polybub/Data/Services"
	"Polybub/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"encoding/json"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		get(w, req)
	}
	if req.Method == "POST" {
		post(w, req)
	}
}

func get(w http.ResponseWriter, req *http.Request) {
	path := "Routes/Ui/Pages/CreateUser/create-user.html"
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

func post(w http.ResponseWriter, req *http.Request) {
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
