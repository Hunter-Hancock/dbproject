package app

import (
	"github.com/Hunter-Hancock/dbproject/db"
	"github.com/Hunter-Hancock/dbproject/handler"
	mw "github.com/Hunter-Hancock/dbproject/middleware"
)

type Application struct {
	TestStore   db.TestStore
	TestHandler *handler.TestHandler
	Middleware  mw.MiddleWare
}

func NewApplication() (*Application, error) {
	sql, err := db.Open()
	if err != nil {
		return nil, err
	}

	testStore := db.NewSQLTestStore(sql)
	testHandler := handler.NewTestHandler(testStore)

	app := &Application{
		TestStore:   testStore,
		TestHandler: testHandler,
		Middleware:  mw.MiddleWare{},
	}

	return app, nil
}
