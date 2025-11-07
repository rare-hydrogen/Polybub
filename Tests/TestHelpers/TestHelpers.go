package TestHelpers

import (
	"Polybub/Utilities"

	"github.com/google/uuid"
)

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
