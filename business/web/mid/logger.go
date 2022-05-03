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
			contextValues, err := web.GetValues(ctx)
			if err != nil {
				return err
			}

			log.Infow("request started", "traceid", contextValues.TraceID, "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr)

			// Call the next handler.
			err = handler(ctx, w, r)

			log.Infow("request completed", "traceid", contextValues.TraceID, "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr, "statuscode", contextValues.StatusCode, "since", time.Since(contextValues.Now))

			// Return the error so it can be handled further up the chain.
			return err
		}
	
		return h
	}

	return m
}