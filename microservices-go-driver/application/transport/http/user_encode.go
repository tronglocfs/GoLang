package http

import (
	"context"
	"encoding/json"
	"net/http"

	infra "github.com/microservices/infrastructure"
)

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	e, ok := response.(errorer)
	if ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	errcheck := json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})

	if errcheck != nil {
		return
	}
}

func codeFrom(err error) int {
	switch err {
	case infra.ErrNotFound:
		return http.StatusNotFound
	case infra.ErrAlreadyExists:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
