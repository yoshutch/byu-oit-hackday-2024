package adapters

import (
	"byu.edu/hackday-profile/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type RestAdapter struct {
	mux            *http.ServeMux
	profileService *services.ProfileService
}

func NewRestAdapter(mux *http.ServeMux, service *services.ProfileService) (*RestAdapter, error) {
	return &RestAdapter{mux, service}, nil
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
	a.mux.HandleFunc("GET /api/profile/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			a.InternalServerErrorHandler(w, r)
			return
		}
		profile, err := a.profileService.LoadProfile(id)
		if err != nil {
			a.InternalServerErrorHandler(w, r)
			return
		}
		if profile == nil {
			a.NotFoundHandler(w, r)
			return
		}
		// get data from service classes
		jsonBytes, err := json.Marshal(profile)
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
