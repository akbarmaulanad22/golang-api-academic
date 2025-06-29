package http

import (
	"encoding/json"
	"net/http"
	"tugasakhir/internal/delivery/http/middleware"
	"tugasakhir/internal/helper"
	"tugasakhir/internal/model"
	"tugasakhir/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type CourseController struct {
	Log     *logrus.Logger
	UseCase *usecase.CourseUseCase
}

func NewCourseController(useCase *usecase.CourseUseCase, logger *logrus.Logger) *CourseController {
	return &CourseController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *CourseController) ListByLecturerUserID(w http.ResponseWriter, r *http.Request) {

	auth := middleware.GetUser(r)

	request := &model.ListCourseRequest{
		UserID: auth.ID,
	}

	// Panggil UseCase
	response, err := c.UseCase.ListByLecturerUserID(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to get course: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.CourseResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *CourseController) Create(w http.ResponseWriter, r *http.Request) {

	// Parse request body
	var request model.CreateCourseRequest
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
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.CourseResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *CourseController) List(w http.ResponseWriter, r *http.Request) {

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
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.CourseResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *CourseController) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]

	// Parse request body
	var request model.UpdateCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	request.Code = code

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
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.CourseResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *CourseController) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]

	var request model.DeleteCourseRequest
	request.Code = code

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
