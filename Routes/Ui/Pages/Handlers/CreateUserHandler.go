package Handlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Models"
	"Polybub/Data/Services"
	"Polybub/Data/Validators"
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/Wrappers/GlobalWrapper"
	"Polybub/Utilities/Permissions"
	"encoding/json"
	"net/http"
)

type checkPasswordVariant struct {
	Models.User   // TODO: Nesting objects like this prevent validation rules from firing
	CheckPassword string
}

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
		Jsend.Error(req.Context(), w, "Error reading template", http.StatusInternalServerError)
		return
	}

	wrappedBody, err := GlobalWrapper.GetWrappedTemplate(body)
	if err != nil {
		Jsend.Error(req.Context(), w, "Error wrapping template", http.StatusInternalServerError)
		return
	}

	Jsend.Ui(req.Context(), w, wrappedBody, http.StatusOK)
}

func postCreateUser(w http.ResponseWriter, req *http.Request) {
	status, err := OAuth2.JwtPermitRequest(req, Permissions.USERS_CRUD, nil)
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error(), status)
		return
	}

	var dto checkPasswordVariant
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&dto)
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error(), http.StatusBadRequest)
		return
	}

	err = Validators.User(dto.User)
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error(), http.StatusBadRequest)
		return
	}

	if dto.Password != dto.CheckPassword {
		Jsend.Error(req.Context(), w, "passwords must match", http.StatusBadRequest)
		return
	}

	d, err := Services.CreateUser(dto.User)
	if err != nil {
		Jsend.Error(req.Context(), w, err.Error(), http.StatusBadRequest)
		return
	}

	err = Services.CreateDefaultPermissions(d.Id)
	if err != nil {
		Jsend.Error(req.Context(), w, "adding default permissions failed", http.StatusInternalServerError)
		return
	}

	err = Services.UpdatePasswordAndSalt(d.Id, dto.Password)
	if err != nil {
		Jsend.Error(req.Context(), w, "invalid password", http.StatusBadRequest)
		return
	}

	Jsend.Success(req.Context(), w, d)
}
