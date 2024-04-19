package app

import (
	"github.com/Hunter-Hancock/dbproject/db"
	"github.com/Hunter-Hancock/dbproject/handler"
	mw "github.com/Hunter-Hancock/dbproject/middleware"
)

type Application struct {
	MyHandler   *handler.Handler
	AuthHandler *handler.AuthHandler
	FoodStore   db.FoodStore
	FoodHandler *handler.FoodHandler
	Middleware  mw.MiddleWare
}

func NewApplication() (*Application, error) {
	sql, err := db.Open()
	if err != nil {
		return nil, err
	}

	authHandler := handler.NewAuthHandler()

	h := handler.NewHandler()

	foodStore := db.NewSQLFoodStore(sql)
	foodHandler := handler.NewFoodHandler(foodStore)

	app := &Application{
		MyHandler:   h,
		AuthHandler: authHandler,
		FoodStore:   foodStore,
		FoodHandler: foodHandler,
		Middleware:  mw.MiddleWare{},
	}

	return app, nil
}
