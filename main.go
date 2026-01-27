package main

import (
	"net/http"

	"github.com/rs/cors"

	"Polybub/Routes"
	"Polybub/Swagger"
	"Polybub/Utilities"
	"Polybub/Utilities/Logger/RequestMiddleware"
)

func main() {
	Utilities.GlobalConfig = Utilities.GetConfig()
	baseUrl := Utilities.GetBaseUrl(Utilities.GlobalConfig)

	corsHandler := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedOrigins:   []string{baseUrl},
		AllowCredentials: true,
		ExposedHeaders:   []string{"Authorization"},
		//Debug:          true,
	})

	mux := Routes.AddRoutes()
	log := RequestMiddleware.LogHandler(mux)
	handler := corsHandler.Handler(log)

	if Utilities.GlobalConfig.Env == "development" {
		Swagger.Setup(Utilities.GlobalConfig, baseUrl, mux)
		http.ListenAndServe(":8090", handler)
	}

	if Utilities.GlobalConfig.Env == "production" {
		if Utilities.GlobalConfig.IsSecure {
			// TODO: Handle HTTPS
			//certFile := "./certs/fullchain.pem"
			//keyFile := "./certs/myserver.key"
			//http.ListenAndServeTLS(":8090", certFile, keyFile, handler)
		} else {
			http.ListenAndServe(":8090", handler)
		}
	}
}
