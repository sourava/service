package web

import (
	"context"
	"net/http"
	"os"
	"syscall"

	"github.com/dimfeld/httptreemux/v5"
)

type Handler func (ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*httptreemux.ContextMux
	shutdown chan os.Signal
	mw []Middleware
}

func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {
	return &App {
		ContextMux: httptreemux.NewContextMux(),
		shutdown: shutdown,
		mw: mw,
	}	 
}

func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

func (a *App) Handle(method string, group string, path string, handler Handler, mw ...Middleware) {

	// App specific middlewares
	handler = wrapMiddleware(a.mw, handler)

	// Handler specific middlewares
	handler = wrapMiddleware(mw, handler)

	h := func (w http.ResponseWriter, r *http.Request) {
		// INJECT CODE

		if err := handler(r.Context(), w, r); err != nil {
			// ERROR Handling
			return
		}

		// INJECT CODE
	}

	finalPath := path
	if group != "" {
		finalPath = "/" + group + finalPath
	}

	a.ContextMux.Handle(method, finalPath, h)
}
