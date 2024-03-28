package main

import (
	"fmt"
	"net/http"
	"os"
	"snake-scape/internal/handler"
	"snake-scape/internal/middleware"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /favicon.ico", handler.ServeFavicon)
	mux.HandleFunc("GET /static/", handler.ServeStaticFiles)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		middleware.Chain(w, r, handler.Introduction)
	})

	mux.HandleFunc("GET /getting-started/technologies", func(w http.ResponseWriter, r *http.Request) {
		middleware.Chain(w, r, handler.Technologies)
	})

	mux.HandleFunc("GET /getting-started/game-configuration", func(w http.ResponseWriter, r *http.Request) {
		middleware.Chain(w, r, handler.GameConfiguration)
	})

	mux.HandleFunc("GET /getting-started/developer-environment", func(w http.ResponseWriter, r *http.Request) {
		middleware.Chain(w, r, handler.DeveloperEnvironment)
	})


	fmt.Println(fmt.Sprintf("server is running on port %s", os.Getenv("PORT")))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	if err != nil {
		fmt.Println(err)
	}

}