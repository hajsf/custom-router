package main

import (
	"fmt"
	_ "module1"
	"net/http"
	"router"
)

func init() {
	fmt.Println("Register Index and static files")
	router.R.RegisterStaticRoute()
	router.R.Index("layout.html", "welcome.html")
}

func main() {
	http.ListenAndServe(":8080", &router.R)
}
