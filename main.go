package main

import (
	"net/http"

	"github.com/rs/cors"

	"Polybub/Routes"
	"Polybub/Utilities"
)

func main() {
	Utilities.GlobalConfig = Utilities.GetConfig()

	corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedOrigins: []string{Utilities.GetCurrentEnv(Utilities.GlobalConfig)},
	})

	mux := Routes.AddRoutes()
	handler := corsHandler.Handler(mux)

	http.ListenAndServe(":8090", handler)
}
