package router

import (
	"context"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"regexp"

	"example.com/assets"
)

var R *Router

func init() {
	R = New(false) // Set to true to serve files from local directory
}

type route struct {
	pattern *regexp.Regexp
	handler http.HandlerFunc
}

type Router struct {
	Debug  bool
	routes map[string][]route
}

type key int

const paramsKey key = 0

func New(debug bool) *Router {
	return &Router{
		Debug:  debug,
		routes: make(map[string][]route),
	}
}

func (r *Router) RegisterRoute(method, pattern string, handler http.HandlerFunc) {
	r.routes[method] = append(r.routes[method], route{regexp.MustCompile(pattern), handler})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if routes, ok := r.routes[req.Method]; ok {
		for _, route := range routes {
			if match := route.pattern.FindStringSubmatch(req.URL.Path); match != nil {
				ctx := context.WithValue(req.Context(), paramsKey, match[1:])
				req = req.WithContext(ctx)
				route.handler(w, req)
				return
			}
		}
	}
	r.ServeStatics(w, req)
}

func (r *Router) ServeStatics(w http.ResponseWriter, req *http.Request) {
	/*	file := strings.TrimPrefix(req.URL.Path, "/static/")
		fmt.Println("Serving static file:", file)
		if r.Debug {
			http.ServeFile(w, req, "assets/static/"+file)
		} else {
			fsys, err := fs.Sub(assets.Static, "static")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fileServer := http.FileServer(http.FS(fsys))
			req.URL.Path = file
			fileServer.ServeHTTP(w, req)
		}

	*/var fileSystem http.FileSystem

	if r.Debug {
		fileSystem = http.Dir("assets/static")
	} else {
		fsys, err := fs.Sub(assets.Static, ".")
		if err != nil {
			log.Fatal(err)
		}
		fileSystem = http.FS(fsys)
	}

	http.FileServer(fileSystem).ServeHTTP(w, req)

}

func (r *Router) ServeTemplate(w http.ResponseWriter, data interface{}, files ...string) {
	var tmpl *template.Template
	prefixedLocal := make([]string, len(files))
	prefixedEmbed := make([]string, len(files))
	for i, file := range files {
		prefixedLocal[i] = "assets/template/" + file
		prefixedEmbed[i] = "template/" + file
	}
	var err error

	if r.Debug {
		tmpl, err = template.ParseFiles(prefixedLocal...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		fsys, err := fs.Sub(assets.Template, ".")
		if err != nil {
			log.Fatal(err)
		}

		tmpl, err = template.ParseFS(fsys, prefixedEmbed...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetParam(req *http.Request, index int) string {
	params := req.Context().Value(paramsKey).([]string)
	return params[index]
}

func (r *Router) Index(files ...string) {
	r.RegisterRoute(http.MethodGet, "/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hi from root")
		// http.ServeFile(w, req, "assets/static/layout.html")
		r.ServeTemplate(w, nil, files...)
	})
}
