package ComponentHandlers

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Models"
	"Polybub/Data/Services"
	"Polybub/Data/Validators"
	"Polybub/Routes/Jsend"
	"Polybub/Routes/Ui/Wrappers"
	"Polybub/Utilities/Permissions"
	"encoding/json"
	"net/http"
)

type checkPasswordVariant struct {
	Models.User   // TODO: Nesting objects like this prevent validation rules from firing
	CheckPassword string
}

func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getCreateUser(w, req)
	case "POST":
		postCreateUser(w, req)
	default:
		Jsend.MethodNotAllowed(w, req)
		return
	}
}

func getCreateUser(w http.ResponseWriter, req *http.Request) {
	data := ""
	Wrappers.TemplateHandler(w, req, "comp.create-user", data)
}

func postCreateUser(w http.ResponseWriter, req *http.Request) {
	status, err := OAuth2.JwtPermitRequest(req, Permissions.USERS_CRUD, nil)
	if err != nil {
		Jsend.Error(w, req, err, err.Error(), status)
		return
	}

	var dto checkPasswordVariant
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&dto)
	if err != nil {
		Jsend.Error(w, req, err, err.Error(), http.StatusBadRequest)
		return
	}

	err = Validators.User(dto.User)
	if err != nil {
		Jsend.Error(w, req, err, err.Error(), http.StatusBadRequest)
		return
	}

	if dto.Password != dto.CheckPassword {
		Jsend.Error(w, req, err, "passwords must match", http.StatusBadRequest)
		return
	}

	d, err := Services.CreateUser(req.Context(), dto.User)
	if err != nil {
		Jsend.Error(w, req, err, err.Error(), http.StatusBadRequest)
		return
	}

	err = Services.CreateDefaultPermissions(req.Context(), d.Id)
	if err != nil {
		Jsend.Error(w, req, err, "adding default permissions failed", http.StatusInternalServerError)
		return
	}

	err = Services.UpdatePasswordAndSalt(req.Context(), d.Id, dto.Password)
	if err != nil {
		Jsend.Error(w, req, err, "invalid password", http.StatusBadRequest)
		return
	}

	Jsend.Success(w, req, d)
}
