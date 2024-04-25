package app

import (
	"github.com/Hunter-Hancock/dbproject/db"
	"github.com/Hunter-Hancock/dbproject/handler"
	mw "github.com/Hunter-Hancock/dbproject/middleware"
	"github.com/Hunter-Hancock/dbproject/session"
	"github.com/gorilla/sessions"
)

type Application struct {
	MyHandler      *handler.Handler
	AuthHandler    *handler.AuthHandler
	AuthStore      db.AuthStore
	FoodStore      db.FoodStore
	FoodHandler    *handler.FoodHandler
	CartHandler    *handler.CartHandler
	CartStore      db.CartStore
	OrderHandler   *handler.OrderHandler
	OrderStore     db.OrderStore
	Middleware     mw.MiddleWare
	SessionManager *session.SessionManager
}

func NewApplication() (*Application, error) {
	sql, err := db.Open()
	if err != nil {
		return nil, err
	}

	store := sessions.NewCookieStore([]byte("SuperSecret"))

	sessionManager := session.NewSessionManager(store)

	authStore := db.NewAuthStore(sql)
	authHandler := handler.NewAuthHandler(authStore, sessionManager)

	h := handler.NewHandler()

	cartStore := db.NewSQLCartStore(sql)
	cartHandler := handler.NewCartHandler(cartStore)

	foodStore := db.NewSQLFoodStore(sql)
	foodHandler := handler.NewFoodHandler(foodStore)

	orderStore := db.NewSQLOrderStore(sql)
	orderHandler := handler.NewOrderHandler(orderStore, cartStore)

	app := &Application{
		MyHandler:    h,
		AuthHandler:  authHandler,
		AuthStore:    authStore,
		FoodStore:    foodStore,
		FoodHandler:  foodHandler,
		CartHandler:  cartHandler,
		CartStore:    cartStore,
		OrderHandler: orderHandler,
		OrderStore:   orderStore,
		Middleware:   mw.MiddleWare{AuthStore: authStore, Session: sessionManager},
	}

	return app, nil
}
