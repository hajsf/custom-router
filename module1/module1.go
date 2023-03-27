package module1

import (
	"net/http"

	"router"
)

func init() {
	router.R.RegisterRoute(http.MethodGet, "/signin", func(w http.ResponseWriter, req *http.Request) {
		router.R.ServeTemplate(w, nil, "signin.html")
	})

	router.R.RegisterRoute(http.MethodPost, "/signin", func(w http.ResponseWriter, req *http.Request) {
		router.R.ServeTemplate(w, nil, "layout.html", "welcome.html")
	})
}
