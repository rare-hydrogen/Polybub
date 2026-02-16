package TestHelpers

import (
	"Polybub/Data"
	"Polybub/Data/Enums/UserGroups"
	"Polybub/Routes/Jsend"
	"Polybub/Utilities"
	"Polybub/Utilities/Logger"
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

var BasicRequestDetails = Logger.RequestDetails{
	UserId:      1,
	UserName:    "username",
	UserGroupId: UserGroups.InternalUsers,
	Verb:        "",
	Endpoint:    "",
	QueryString: "",
	Body:        "",
}

func FindRoot() string {
	full, err := os.Getwd()
	parent, _, ok := strings.Cut(full, "Polybub/")

	if err != nil || !ok {
		println(full)
		println(parent)
		panic("failed to locate root")
	}

	return parent + "Polybub/"
}

func UniqueTestConfig() Utilities.Config {
	db := "file:testdb" + uuid.NewString() + "?mode=memory&cache=shared"

	return Utilities.Config{
		Env:        "development",
		Connection: db,
		CertPath:   "./.certs/",
		Pepper:     "+1ItkRehw/2xPXW0jd8a040QLnROEoZKYFtD4hN2c5U=", // fake
		Port:       "8080",
		Domain:     "polybub",
		TopDomain:  ".com",
		ApiTitle:   "polybub swagger title",
		ApiVersion: "1.0.0",
		CookieName: "polybub-jwt",
	}
}

func ApplySchema() {
	db := Data.GetConnection()

	root := FindRoot()
	schema, err := os.ReadFile(root + "Data/Schema/schema.sql")
	if err != nil {
		panic("failed to apply schema")
	}

	err = db.Exec(string(schema)).Error
	if err != nil {
		panic("failed to apply schema")
	}
}

func TestHandler(w http.ResponseWriter, req *http.Request) {
	Jsend.Success(req.Context(), w, nil, http.StatusOK)
}

func TestReqContext(rdv Logger.RequestDetails) context.Context {
	return context.WithValue(context.Background(), "requestDetails", rdv)
}
