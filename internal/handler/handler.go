package handler

import (
	"net/http"
	"path/filepath"
	"snake-scape/internal/middleware"
	"snake-scape/internal/template"
)

func ServeFavicon(w http.ResponseWriter, r *http.Request) {
	filePath := "favicon.ico"
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):]
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func Home(ctx *middleware.CustomContext, w http.ResponseWriter, r *http.Request) {
	template.Home().Render(ctx, w)
}