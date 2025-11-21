package Components

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Routes/Ui/Components/FooBar"
	"Polybub/Utilities/Permissions"
	"net/http"
)

func AddComponentRoutes(mux *http.ServeMux) {
	OAuth2.JwtPermit(mux, "/comp/foobar", FooBar.Handler, Permissions.DASHBOARD_R, nil)
}
