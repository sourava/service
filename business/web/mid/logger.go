package mid

import (
	"context"
	"net/http"
	"time"

	"github.com/sourava/service/foundation/web"
	"go.uber.org/zap"
)

func Logger(log *zap.SugaredLogger) web.Middleware {
	m := func (handler web.Handler) web.Handler {
		h := func (ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			traceID := 00000
			statuscode := http.StatusOK
			now := time.Now()

			log.Info("request started", "traceID", traceID, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)
			
			err := handler(ctx, w, r);
	
			log.Info("request completed", "traceID", traceID, "method", r.Method, "path", r.URL.Path,  "remoteaddr", r.RemoteAddr, "statuscode", statuscode, "since", time.Since(now))
	
			return err
		}
	
		return h
	}

	return m
}