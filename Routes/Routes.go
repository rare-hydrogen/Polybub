package Routes

import (
	"Polybub/Routes/Api"
	"Polybub/Routes/Static"
	"Polybub/Routes/Ui"
	"net/http"
)

func AddRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	Api.AddApiRoutes(mux)
	Static.AddStaticRoutes(mux)
	Ui.AddUiRoutes(mux)

	return mux
}
