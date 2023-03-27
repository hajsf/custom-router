package router

func init() {
	/*	exe, err := os.Executable()
		if err != nil {
			panic(err)
		}
		dir := filepath.Dir(exe)
		fmt.Println("The app is running at:", dir)

		_, file, _, ok := runtime.Caller(0)
		if !ok {
			panic("failed to get caller information")
		}
		root := filepath.Dir(file)
		fmt.Println("The root path of the code is:", root)
		fmt.Println("Below all the embeded files and folders")
		fs.WalkDir(assets.Static, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				fmt.Println(path)
			}
			return nil
		})
	*/
}

func init() {
	/*	R.RegisterRoute(http.MethodGet, "/file", func(w http.ResponseWriter, req *http.Request) {
				fmt.Println("Hi from module1")
				if R.Debug {
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
			R.RegisterRoute(http.MethodGet, "/list", func(w http.ResponseWriter, req *http.Request) {
				fmt.Println("Hi from module1")
				if R.Debug {
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

				R.RegisterRoute(http.MethodGet, "/static/", func(w http.ResponseWriter, req *http.Request) {
				file := strings.TrimPrefix(req.URL.Path, "/static/")
				fmt.Println("Serving static file:", file)
				fmt.Println("static folder:", assets.StaticFilesDir)
				if R.Debug {
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
			}) */
}
