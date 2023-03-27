package assets

import (
	"embed"
	"path/filepath"
	"runtime"
)

//go:embed static/*
var Static embed.FS

//go:embed template/*
var Template embed.FS

// StaticFilesDir is the path to the directory containing the static files.
var StaticFilesDir = filepath.Join(".", "assets/static/")

// FilePath returns the absolute path to a file in the static module.
func FilePath(name string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), name)
}
