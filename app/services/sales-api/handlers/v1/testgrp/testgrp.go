// Package testgrp maintains the group of handlers for testing.
package testgrp

import (
	"context"
	"net/http"

	"github.com/sourava/service/foundation/web"
	"go.uber.org/zap"
)

// Handlers manages the set of check endpoints.
type Handlers struct {
	Build string
	Log   *zap.SugaredLogger
}

func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	} {
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
