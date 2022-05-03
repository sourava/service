package web

import (
	"context"
	"errors"
	"time"
)

type ctxKey int

const key ctxKey = 1

type Values struct {
	TraceID string
	Now time.Time
	StatusCode int
}

func GetValues(ctx context.Context) (*Values, error) {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return nil, errors.New("web value is missing from context")
	}
	return v, nil
}

func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return "00000-00000-00000"
	}
	return v.TraceID
}

func SetStatusCode(ctx context.Context, statusCode int) error {
	v, ok := ctx.Value(key).(*Values)
	if !ok {
		return errors.New("web value is missing from context")
	}
	v.StatusCode = statusCode
	return nil
}


