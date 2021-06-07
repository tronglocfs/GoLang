package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
)

/*- Creating Interface
- Business Logic
- DateService provides operations on Date:
	+ Checking Date parameter format yyyy-MM-DD
	+ Confirm the microservice is successful
	+ Return current day */
type DateService interface {
	Validate(string) (bool, error)
	Status() (string, error)
	GetCurDay() (string, error)
}

// implement interface
type dateService struct{}

func (dateService) Validate(date string) (bool, error) {
	const layout = "2006-01-02"

	// Calling Parse() method with its parameters
	_, err := time.Parse(layout, date)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (dateService) Status() (string, error) {
	return "ok", nil
}

func (dateService) GetCurDay() (string, error) {
	date := time.Now().Format("2006-01-02")
	return date, nil
}

// requests and responses
type validateRequest struct {
	Date string `json:"date"`
}

type validateResponse struct {
	Validate bool   `json:"validate"`
	Err      string `json:"err,omitempty"`
}

type statusRequest struct {
}

type statusResponse struct {
	Status string `json:"status"`
	Err    string `json:"err,omitempty"`
}

type getCurDayRequest struct {
}

type getCurDayResponse struct {
	Date string `json:"date"`
	Err  string `json:"err,omitempty"`
}

// endpoints
func makeValidateEndpoint(svc DateService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(validateRequest)
		v, err := svc.Validate(req.Date)

		if err != nil {
			return validateResponse{v, err.Error()}, nil
		}

		return validateResponse{v, ""}, nil
	}
}

func makeStatusEndpoint(svc DateService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		_ = request.(statusRequest)
		v, err := svc.Status()

		if err != nil {
			return statusResponse{v, err.Error()}, nil
		}

		return statusResponse{v, ""}, nil
	}
}

func makeGetCurDay(svc DateService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		_ = request.(getCurDayRequest)
		v, err := svc.GetCurDay()

		if err != nil {
			return getCurDayResponse{v, err.Error()}, nil
		}

		return getCurDayResponse{v, ""}, nil
	}
}

// Transport

func main() {
	svc := dateService{}

	validateHandler := httptransport.NewServer(
		makeValidateEndpoint(svc),
		decodeValidateRequest,
		encodeResponse,
	)

	statusHandler := httptransport.NewServer(
		makeStatusEndpoint(svc),
		decodeStatusEndpoint,
		encodeResponse,
	)

	getCurDayHandler := httptransport.NewServer(
		makeGetCurDay(svc),
		decodeGetCurDayEndpoint,
		encodeResponse,
	)

	http.Handle("/validate", validateHandler)
	http.Handle("/status", statusHandler)
	http.Handle("/getCurDay", getCurDayHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func decodeValidateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request validateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeStatusEndpoint(_ context.Context, r *http.Request) (interface{}, error) {
	var request statusRequest

	return request, nil
}

func decodeGetCurDayEndpoint(_ context.Context, r *http.Request) (interface{}, error) {
	var request getCurDayRequest

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
