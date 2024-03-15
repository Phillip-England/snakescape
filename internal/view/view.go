package view

import (
	"net/http"
	"snake-scape/internal/component"
	"snake-scape/internal/middleware"
)

func Home(ctx *middleware.CustomContext, w http.ResponseWriter, r *http.Request) {
	component.Hello("World!").Render(r.Context(), w)
}