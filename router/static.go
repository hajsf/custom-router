package module1

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"example.com/assets"
	"example.com/router"
)

func init() {
	router.R.RegisterRoute(http.MethodGet, "/file", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hi from module1")
		if router.R.Debug {
			http.ServeFile(w, req, "assets/static/file.txt")
		} else {
			fsys, err := fs.Sub(assets.Static, "static")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			data, err := fs.ReadFile(assets.Static, "static/file.txt")
			if err != nil {
				fmt.Println("file not found:", err)
			} else {
				fmt.Println("file found:", string(data))
			}
			fileServer := http.FileServer(http.FS(fsys))
			req.URL.Path = "file.txt"
			fileServer.ServeHTTP(w, req)
		}

	})
}

func init() {
	router.R.RegisterRoute(http.MethodGet, "/list", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hi from module1")
		if router.R.Debug {
			files, err := os.ReadDir("assets/static")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			for _, file := range files {
				fmt.Fprintln(w, file.Name())
			}
		} else {
			fsys, err := fs.Sub(assets.Static, "static")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() {
					fmt.Fprintln(w, path)
				}
				return nil
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})

	router.R.RegisterRoute(http.MethodGet, "/static/", func(w http.ResponseWriter, req *http.Request) {
		file := strings.TrimPrefix(req.URL.Path, "/static/")
		fmt.Println("Serving static file:", file)
		fmt.Println("static folder:", assets.StaticFilesDir)
		if router.R.Debug {
			http.ServeFile(w, req, assets.StaticFilesDir+"/"+file) //"assets/static/"
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
