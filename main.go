package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"example.com/assets"
	_ "example.com/module1"
	_ "example.com/module2"
	"example.com/router"
)

func init() {
	router.R.Index("layout.html", "welcome.html")
}

func main() {
	exe, err := os.Executable()
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
	http.ListenAndServe(":8080", router.R)
}
