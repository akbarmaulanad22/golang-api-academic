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

type StudentController struct {
	Log     *logrus.Logger
	UseCase *usecase.StudentUseCase
}

func NewStudentController(useCase *usecase.StudentUseCase, logger *logrus.Logger) *StudentController {
	return &StudentController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *StudentController) ListByCourseCode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	courseCode := vars["courseCode"]

	request := &model.ListStudentRequest{
		CourseCode: courseCode,
	}

	// Panggil UseCase
	response, err := c.UseCase.ListByCourseCode(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to get course: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.StudentResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *StudentController) Create(w http.ResponseWriter, r *http.Request) {

	// Parse request body
	var request model.CreateStudentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Panggil UseCase
	response, err := c.UseCase.Create(r.Context(), &request)
	if err != nil {
		c.Log.Printf("Failed to get student: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.StudentAdminResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *StudentController) List(w http.ResponseWriter, r *http.Request) {

	// Panggil UseCase
	response, err := c.UseCase.List(r.Context())
	if err != nil {
		c.Log.Printf("Failed to get student: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.StudentAdminResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *StudentController) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	npm := vars["npm"]

	npmInt, err := strconv.Atoi(npm)
	if err != nil {
		c.Log.Printf("Failed to parse npm: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Parse request body
	var request model.UpdateStudentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//
	request.Npm = uint(npmInt)

	// Panggil UseCase
	response, err := c.UseCase.Update(r.Context(), &request)
	if err != nil {
		c.Log.Printf("Failed to get student: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.StudentAdminResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *StudentController) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	npm := vars["npm"]

	npmInt, err := strconv.Atoi(npm)
	if err != nil {
		c.Log.Printf("Failed to parse npm: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	var request model.DeleteStudentRequest
	request.Npm = uint(npmInt)

	// Panggil UseCase
	if err := c.UseCase.Delete(r.Context(), &request); err != nil {
		c.Log.Printf("Failed to get student: %v", err)
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
