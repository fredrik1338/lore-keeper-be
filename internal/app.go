package internal

import (
	"context"
)

type App struct {
	server Server
}

//TODO use context to handle graceful shutdown
func NewApp(ctx context.Context) App {

	return App{
		server: newServer(),
	}
}

func (app App) Start() {
	app.server.Start()
}
