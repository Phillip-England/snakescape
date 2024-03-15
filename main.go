package main

import (
	"fmt"
	"net/http"
	"os"
	"snake-scape/internal/middleware"
	"snake-scape/internal/view"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		middleware.Chain(w, r, view.Home, middleware.Init, middleware.Log)
	})


	fmt.Println(fmt.Sprintf("server is running on port %s", os.Getenv("PORT")))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	if err != nil {
		fmt.Println(err)
	}

}