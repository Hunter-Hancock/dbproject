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
	r.Get("/", app.MyHandler.Home)

	r.Get("/signup", app.AuthHandler.SigupPage)

	r.Get("/products", app.FoodHandler.HandleGetAll)

	r.Route("/api", func(r chi.Router) {
		r.Get("/categories", app.FoodHandler.HandleGetCategories)
		r.Get("/products/{id}", app.FoodHandler.HandleGetProductsBySubID)
	})

	r.Get("/category/{name}", app.FoodHandler.HandleGetCategory)
	r.Get("/subcat/{id}", app.FoodHandler.HandleGetProductsBySubID)

	fileServer := http.FileServer(http.FS(view.Files))
	r.Handle("/*", fileServer)

	return r
}
