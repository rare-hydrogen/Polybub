package main

import (
	"net/http"

	"github.com/rs/cors"

	"Polybub/Routes"
)

func main() {
	baseUrl := "http://localhost"

	corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedOrigins: []string{baseUrl},
	})

	mux := Routes.AddRoutes()
	handler := corsHandler.Handler(mux)

	http.ListenAndServe(":8090", handler)
}
