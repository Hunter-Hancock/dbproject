package routes

import (
	"fmt"
	"net/http"

	"github.com/Hunter-Hancock/dbproject/app"
	"github.com/Hunter-Hancock/dbproject/view"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	app, err := app.NewApplication()
	if err != nil {
		fmt.Println(err)
	}

	r.Use(app.Middleware.RequireUser)
	r.Get("/", app.TestHandler.Home)
	r.Post("/click", app.TestHandler.Click)

	r.Get("/products", app.FoodHandler.HandleGetAll)

	r.Route("/api", func(r chi.Router) {
		r.Get("/categories", app.FoodHandler.HandleGetCategories)
	})

	fileServer := http.FileServer(http.FS(view.Files))
	r.Handle("/*", fileServer)

	return r
}
