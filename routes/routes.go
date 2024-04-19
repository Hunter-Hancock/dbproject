package routes

import (
	"fmt"
	"net/http"

	"github.com/Hunter-Hancock/dbproject/app"
	"github.com/Hunter-Hancock/dbproject/view"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	app, err := app.NewApplication()
	if err != nil {
		fmt.Println(err)
	}

	r.Use(middleware.Recoverer)
	// Cors options
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders:   []string{"Content-Type", "Set-Cookie", "Cookie"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(app.Middleware.RequireUser)
	r.Get("/", app.MyHandler.Home)

	r.Get("/signup", app.AuthHandler.SigupPage)

	r.Get("/products", app.FoodHandler.HandleGetAll)

	r.Route("/api", func(r chi.Router) {
		r.Get("/categories", app.FoodHandler.HandleGetCategories)
		r.Get("/products/{id}", app.FoodHandler.HandleGetProductsBySubID)

		r.Post("/signup", app.AuthHandler.Signup)
	})

	r.Get("/category/{name}", app.FoodHandler.HandleGetCategory)
	r.Get("/subcat/{id}", app.FoodHandler.HandleGetProductsBySubID)

	fileServer := http.FileServer(http.FS(view.Files))
	r.Handle("/*", fileServer)

	return r
}
