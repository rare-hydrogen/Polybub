package Ui

import (
	"net/http"

	"Polybub/Routes/Ui/Components"
	"Polybub/Routes/Ui/Pages"
)

func AddUiRoutes(mux *http.ServeMux) {
	Pages.AddPageRoutes(mux)
	Components.AddComponentRoutes(mux)
}
