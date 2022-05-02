// Package testgrp maintains the group of handlers for testing.
package testgrp

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// Handlers manages the set of check endpoints.
type Handlers struct {
	Build string
	Log   *zap.SugaredLogger
}

func (h Handlers) Test(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	} {
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)

	statusCode := http.StatusOK

	h.Log.Infow("readiness", "statusCode", statusCode, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)
}
