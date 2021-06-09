package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	str := fmt.Sprintf("%v", response)

	if strings.Contains(str, "no documents") {
		EncodeErrorNotFound(ctx, errors.New("not found"), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	return json.NewEncoder(w).Encode(response)
}

func EncodeErrorNotFound(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"err": err.Error(),
	})
}
