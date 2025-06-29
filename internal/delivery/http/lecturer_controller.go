package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tugasakhir/internal/helper"
	"tugasakhir/internal/model"
	"tugasakhir/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type LecturerController struct {
	Log     *logrus.Logger
	UseCase *usecase.LecturerUseCase
}

func NewLecturerController(useCase *usecase.LecturerUseCase, logger *logrus.Logger) *LecturerController {
	return &LecturerController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *LecturerController) Create(w http.ResponseWriter, r *http.Request) {

	// Parse request body
	var request model.CreateLecturerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Panggil UseCase
	response, err := c.UseCase.Create(r.Context(), &request)
	if err != nil {
		c.Log.Printf("Failed to get lecturer: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.LecturerResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *LecturerController) List(w http.ResponseWriter, r *http.Request) {

	// Panggil UseCase
	response, err := c.UseCase.List(r.Context())
	if err != nil {
		c.Log.Printf("Failed to get lecturer: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.LecturerResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *LecturerController) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nidn := vars["nidn"]

	nidnInt, err := strconv.Atoi(nidn)
	if err != nil {
		c.Log.Printf("Failed to parse nidn: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Parse request body
	var request model.UpdateLecturerRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//
	request.NIDN = uint(nidnInt)

	// Panggil UseCase
	response, err := c.UseCase.Update(r.Context(), &request)
	if err != nil {
		c.Log.Printf("Failed to get lecturer: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.LecturerResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *LecturerController) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nidn := vars["nidn"]

	nidnInt, err := strconv.Atoi(nidn)
	if err != nil {
		c.Log.Printf("Failed to parse nidn: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	var request model.DeleteLecturerRequest
	request.NIDN = uint(nidnInt)

	// Panggil UseCase
	if err := c.UseCase.Delete(r.Context(), &request); err != nil {
		c.Log.Printf("Failed to get lecturer: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[bool]{Data: true}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
