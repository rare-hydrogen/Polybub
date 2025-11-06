package Routes

import (
	"net/http"
)

func AddRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/foobar", Handler)

	return mux
}

func Handler(w http.ResponseWriter, req *http.Request) {
	println("blah blah blah")
}
