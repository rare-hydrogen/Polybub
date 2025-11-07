package main

import (
	"net/http"

	"github.com/rs/cors"

	"Polybub/Routes"
	"Polybub/Swagger"
	"Polybub/Utilities"
)

func main() {
	Utilities.GlobalConfig = Utilities.GetConfig()
	baseUrl := Utilities.GetBaseUrl(Utilities.GlobalConfig)

	corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedOrigins: []string{baseUrl},
	})

	mux := Routes.AddRoutes()
	handler := corsHandler.Handler(mux)

	if Utilities.GlobalConfig.Env == "development" {
		Swagger.Setup(Utilities.GlobalConfig, baseUrl, mux)
		http.ListenAndServe(":8090", handler)
	}

	if Utilities.GlobalConfig.Env == "production" {
		http.ListenAndServe(":8090", handler)
	}
}
