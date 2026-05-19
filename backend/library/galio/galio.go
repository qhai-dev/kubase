package galio

import (
	"context"
	"net/http"
	"sync"

	"google.golang.org/grpc"
)

type App struct {
	conf   any
	ctx    context.Context
	cancel context.CancelFunc

	gs *grpc.Server
	rs *http.Server

	mu sync.Mutex
}

func New() *App {
	ctx, cancel := context.WithCancel(context.Background())

	return &App{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (app *App) Run() {

}
