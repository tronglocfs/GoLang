package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/microservices/application/endpoints"
)

func EncodeCreateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	v, err := json.Marshal(&response)

	if err != nil {
		EncodeErrorInternal(ctx, err, w)
		return nil
	}

	var res endpoints.CreateUserResponse

	err = json.Unmarshal(v, &res)

	if err != nil {
		EncodeErrorInternal(ctx, err, w)
		return nil
	}

	if res.Err != "" {
		EncodeErrorBadRequest(ctx, errors.New(res.Err), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	return json.NewEncoder(w).Encode(response)
}

func EncodeGetUserByIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	v, err := json.Marshal(&response)

	if err != nil {
		EncodeErrorInternal(ctx, err, w)
		return nil
	}

	var res endpoints.GetUserByIdResponse

	err = json.Unmarshal(v, &res)

	if err != nil {
		EncodeErrorInternal(ctx, err, w)
		return nil
	}

	if res.Err != "" {
		EncodeErrorNotFound(ctx, errors.New(res.Err), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	return json.NewEncoder(w).Encode(response)
}

func EncodeDeleteUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	v, err := json.Marshal(&response)

	if err != nil {
		EncodeErrorInternal(ctx, err, w)
		return nil
	}

	var res endpoints.DeleteUserResponse

	err = json.Unmarshal(v, &res)

	if err != nil {
		EncodeErrorInternal(ctx, err, w)
		return nil
	}

	if res.Err != "" {
		EncodeErrorNotFound(ctx, errors.New(res.Err), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	return json.NewEncoder(w).Encode(response)
}

func EncodeUpdateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	v, err := json.Marshal(&response)

	if err != nil {
		EncodeErrorInternal(ctx, err, w)
		return nil
	}

	var res endpoints.UpdateUserResponse

	err = json.Unmarshal(v, &res)

	if err != nil {
		EncodeErrorInternal(ctx, err, w)
		return nil
	}

	if res.Err != "" {
		EncodeErrorBadRequest(ctx, errors.New(res.Err), w)
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

func EncodeErrorBadRequest(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"err": err.Error(),
	})
}

func EncodeErrorInternal(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"err": err.Error(),
	})
}
