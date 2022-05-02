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
}

func NewApp(shutdown chan os.Signal) *App {
	return &App {
		ContextMux: httptreemux.NewContextMux(),
		shutdown: shutdown,
	}	 
}

func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

func (a *App) Handle(method string, group string, path string, handler Handler) {

	h := func (w http.ResponseWriter, r *http.Request) {
		// PRE PROCESSING

		if err := handler(r.Context(), w, r); err != nil {
			// ERROR Handling
			return
		}

		// POST PROCESSING
	}

	finalPath := path
	if group != "" {
		finalPath = "/" + group + finalPath
	}

	a.ContextMux.Handle(method, finalPath, h)
}
