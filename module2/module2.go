package module2

import (
	"net/http"

	"example.com/router"
)

func init() {
	router.R.RegisterRoute(http.MethodGet, "/signin", func(w http.ResponseWriter, req *http.Request) {
		router.R.ServeTemplate(w, nil, "signin.html")
	})

	router.R.RegisterRoute(http.MethodPost, "/signin", func(w http.ResponseWriter, req *http.Request) {
		router.R.ServeTemplate(w, nil, "welcome.html")
	})
}
