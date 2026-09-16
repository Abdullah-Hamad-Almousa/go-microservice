package app

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{router: loadRoutes()}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe() //We use %w to wrap our error with another error around it
	if err != nil {
		return fmt.Errorf("Failed to start server: %w", err)
	} //To make 2 errors print

	return nil
}
