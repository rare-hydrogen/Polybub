package Routes

import (
	"Polybub/Routes/ApiRoutes"
	"net/http"
)

func AddRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	ApiRoutes.AddApiRoutes(mux)

	return mux
}
