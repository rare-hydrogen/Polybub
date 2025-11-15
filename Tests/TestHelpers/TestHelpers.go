package TestHelpers

import (
	"Polybub/Data"
	"Polybub/Utilities"
	"os"
	"strings"

	"github.com/google/uuid"
)

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
