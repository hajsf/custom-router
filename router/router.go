package router

import (
	"config"
	"context"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"regexp"
	"strings"

	"assets"
)

var R Router

// call R.Init(config.Debug) from an init function in the same module where the R variable is defined to be sure
// that the R.Init(config.Debug) function is called before that init function.
func init() {
	fmt.Println("Initiate the router")
	R.Init(config.Debug) // Pass true or false to set the Debug field
}

type route struct {
	pattern *regexp.Regexp
	handler http.HandlerFunc
}

type Router struct {
	Debug            bool
	routes           map[string][]route
	NotFoundHandler  http.HandlerFunc
	ForbiddenHandler http.HandlerFunc
}

type key int

const paramsKey key = 0

func (r *Router) Init(debug bool) {
	r.Debug = debug
	r.routes = make(map[string][]route)
	r.NotFoundHandler = func(w http.ResponseWriter, req *http.Request) {
		r.ServeTemplate(w, nil, "layout.html", "404.html")
	}
	r.ForbiddenHandler = func(w http.ResponseWriter, req *http.Request) {
		r.ServeTemplate(w, nil, "layout.html", "403.html")
	}
}

func (r *Router) RegisterRoute(method, pattern string, handler http.HandlerFunc) {
	log.Printf("Adding route: %s %s", method, pattern)
	compiledPath := compilePath(pattern)
	fmt.Println("pattern compiled:", compiledPath)
	r.routes[method] = append(r.routes[method], route{compiledPath, handler})
	//r.routes[method] = append(r.routes[method], route{regexp.MustCompile(pattern), handler})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if routes, ok := r.routes[req.Method]; ok {
		for _, route := range routes {
			if match := route.pattern.FindStringSubmatch(req.URL.Path); match != nil {
				params := make(map[string]string)
				for i, name := range route.pattern.SubexpNames() {
					if i > 0 && i <= len(match) {
						params[name] = match[i]
					}
				}
				ctx := context.WithValue(req.Context(), paramsKey, params)
				req = req.WithContext(ctx)
				route.handler(w, req)
				return
			}
		}
	}

	if r.NotFoundHandler != nil {
		r.NotFoundHandler(w, req)
	} else {
		http.NotFound(w, req)
	}
}

//func (r *Router) NotFound(w http.ResponseWriter, req *http.Request) {
//	http.ServeFile(w, req, "path/to/404.html")
//}

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

func (r *Router) RegisterStaticRoute() {
	r.RegisterRoute(http.MethodGet, "/static/*file", func(w http.ResponseWriter, req *http.Request) {
		file := URLParam(req, "file")
		fmt.Println("Serving static file:", file)
		fmt.Println("static folder:", assets.StaticFilesDir)
		if r.Debug {
			http.ServeFile(w, req, assets.StaticFilesDir+"/"+file)
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
	})
}

func compilePath(path string) *regexp.Regexp {
	fmt.Println("path recieved:", path)
	var regex strings.Builder
	regex.WriteString(`^`)
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			regex.WriteString(`(?P<`)
			regex.WriteString(part[1:])
			regex.WriteString(`>[^/]+)`)
		} else if strings.HasPrefix(part, "*") {
			regex.WriteString(`(?P<`)
			regex.WriteString(part[1:])
			regex.WriteString(`>.+)`)
		} else {
			regex.WriteString(regexp.QuoteMeta(part))
		}
		regex.WriteString("/")
	}
	regexString := regex.String()[:regex.Len()-1]
	regexString += `$`
	return regexp.MustCompile(regexString)
}

func URLParam(r *http.Request, name string) string {
	fmt.Println("name:", name)
	ctx := r.Context()
	params, ok := ctx.Value(paramsKey).(map[string]string)
	if !ok {
		fmt.Println("iisue")
		return ""
	}
	value, ok := params[name]
	if !ok {
		fmt.Println("not ok")
		return ""
	}
	fmt.Println("value:", value)
	return value
}

func (r *Router) Forbidden(w http.ResponseWriter, req *http.Request) {
	if r.ForbiddenHandler != nil {
		r.ForbiddenHandler(w, req)
	} else {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	}
}
