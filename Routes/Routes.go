package Routes

import (
	"Polybub/Routes/ApiRoutes"
	"Polybub/Routes/PageRoutes"
	"Polybub/Routes/StaticRoutes"
	"net/http"
)

func AddRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	ApiRoutes.AddApiRoutes(mux)
	PageRoutes.AddPageRoutes(mux)
	StaticRoutes.AddStaticRoutes(mux)

	return mux
}
