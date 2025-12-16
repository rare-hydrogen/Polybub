package Components

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Routes/Ui/Components/Handlers"
	"Polybub/Utilities/Permissions"
	"net/http"
)

func AddComponentRoutes(mux *http.ServeMux) {
	OAuth2.JwtPermit(mux, "/comp/foobar", Handlers.FooBarHandler, Permissions.DASHBOARD_R, nil)
}
