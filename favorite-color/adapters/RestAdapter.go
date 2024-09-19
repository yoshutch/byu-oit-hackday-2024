package adapters

import (
	"encoding/json"
	"net/http"
)

type RestAdapter struct {
	mux *http.ServeMux
}

func NewRestAdapter(mux *http.ServeMux) (*RestAdapter, error) {
	return &RestAdapter{mux}, nil
}

type ErrorDTO struct {
	Message string `json:"message"`
}

func (RestAdapter) InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	jsonBytes, _ := json.Marshal(ErrorDTO{Message: "Internal Server Error"})
	_, _ = w.Write(jsonBytes)
}

func (RestAdapter) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	jsonBytes, _ := json.Marshal(ErrorDTO{Message: "Not Found"})
	_, _ = w.Write(jsonBytes)
}

func (a RestAdapter) HandleRoutes() {
	a.mux.HandleFunc("GET /api/foo", func(w http.ResponseWriter, r *http.Request) {
		type Foo struct {
			Foo string `json:"foo"`
		}
		// get data from service classes
		jsonBytes, err := json.Marshal(Foo{Foo: "bar"})
		if err != nil {
			a.InternalServerErrorHandler(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonBytes)
	})

	// return JSON not found for any undefined routes
	a.mux.HandleFunc("/api/", a.NotFoundHandler)
}
