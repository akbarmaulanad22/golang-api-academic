package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tugasakhir/internal/delivery/http/middleware"
	"tugasakhir/internal/helper"
	"tugasakhir/internal/model"
	"tugasakhir/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type EnrollmentController struct {
	Log     *logrus.Logger
	UseCase *usecase.EnrollmentUseCase
}

func NewEnrollmentController(useCase *usecase.EnrollmentUseCase, logger *logrus.Logger) *EnrollmentController {
	return &EnrollmentController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *EnrollmentController) ListByStudentUserID(w http.ResponseWriter, r *http.Request) {

	auth := middleware.GetUser(r)

	request := &model.ListEnrollmentRequest{
		UserID: auth.ID,
	}

	// Panggil UseCase
	response, err := c.UseCase.ListByStudentUserID(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to get enrollment: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.EnrollmentResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *EnrollmentController) Create(w http.ResponseWriter, r *http.Request) {

	// Parse request body
	var request model.CreateEnrollmentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Panggil UseCase
	response, err := c.UseCase.Create(r.Context(), &request)
	if err != nil {
		c.Log.Printf("Failed to get course: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.EnrollmentAdminResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *EnrollmentController) List(w http.ResponseWriter, r *http.Request) {

	// Panggil UseCase
	response, err := c.UseCase.List(r.Context())
	if err != nil {
		c.Log.Printf("Failed to get course: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.EnrollmentAdminResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *EnrollmentController) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Log.Printf("Failed to parse id: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Parse request body
	var request model.UpdateEnrollmentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	request.ID = uint(idInt)

	// Panggil UseCase
	response, err := c.UseCase.Update(r.Context(), &request)
	if err != nil {
		c.Log.Printf("Failed to get course: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.EnrollmentAdminResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *EnrollmentController) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Log.Printf("Failed to parse id: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	var request model.DeleteEnrollmentRequest
	request.ID = uint(idInt)

	// Panggil UseCase
	if err := c.UseCase.Delete(r.Context(), &request); err != nil {
		c.Log.Printf("Failed to get course: %v", err)
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
