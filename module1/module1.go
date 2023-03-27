package module1

import (
	"fmt"
	"net/http"

	"router"
)

func init() {
	router.R.RegisterRoute(http.MethodGet, "/signin", func(w http.ResponseWriter, req *http.Request) {
		router.R.ServeTemplate(w, nil, "layout.html", "signin.html")
	})

	router.R.RegisterRoute(http.MethodPost, "/signin", func(w http.ResponseWriter, req *http.Request) {
		// router.R.ServeTemplate(w, nil, "layout.html", "welcome.html")
		router.R.Forbidden(w, req)
	})

	router.R.RegisterRoute(http.MethodGet, "/hello/:name", func(w http.ResponseWriter, req *http.Request) {
		name := router.URLParam(req, "name")
		fmt.Fprintf(w, "Hello %s!", name)
	})
}
