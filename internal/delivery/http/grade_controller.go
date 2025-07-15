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

type GradeController struct {
	Log     *logrus.Logger
	UseCase *usecase.GradeUseCase
}

func NewGradeController(useCase *usecase.GradeUseCase, logger *logrus.Logger) *GradeController {
	return &GradeController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *GradeController) ListByStudentUserID(w http.ResponseWriter, r *http.Request) {

	// get user by context in middleware
	auth := middleware.GetUser(r)

	request := &model.ListGradeRequest{
		UserID: auth.ID,
	}

	// Panggil UseCase
	response, err := c.UseCase.ListByStudentUserID(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to get grades: %v", err)
		http.Error(w, "Failed to get grades", http.StatusInternalServerError)
		return
	}

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.GradeResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *GradeController) ListByNpmAndCourseCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	npm := vars["npm"]
	npmInt, err := strconv.Atoi(npm)
	if err != nil {
		c.Log.Printf("Failed to parse npm: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	courseCode := vars["courseCode"]

	request := &model.ListInLecturerGradeRequest{
		Npm:        uint(npmInt),
		CourseCode: courseCode,
	}

	// Panggil UseCase
	response, err := c.UseCase.ListByNpmAndCourseCode(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to get grades: %v", err)
		http.Error(w, "Failed to get grades", http.StatusInternalServerError)
		return
	}

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.GradeInLecturerResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func (c *GradeController) Create(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	npm := vars["npm"]
	npmInt, err := strconv.Atoi(npm)
	if err != nil {
		c.Log.Printf("Failed to parse npm: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	courseCode := vars["courseCode"]

	// Parse request body
	var request model.CreateGradeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	request.Npm = uint(npmInt)
	request.CourseCode = courseCode

	// Panggil UseCase
	response, err := c.UseCase.Create(r.Context(), &request)
	if err != nil {
		c.Log.Printf("Failed to get grade: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.GradeInLecturerResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *GradeController) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Log.Printf("Failed to parse id: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Parse request body
	var request model.UpdateGradeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//
	request.ID = uint(idInt)

	// Panggil UseCase
	response, err := c.UseCase.Update(r.Context(), &request)
	if err != nil {
		c.Log.Printf("Failed to get grade: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.GradeInLecturerResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *GradeController) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Log.Printf("Failed to parse id: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	var request model.DeleteGradeRequest
	request.ID = uint(idInt)

	// Panggil UseCase
	if err := c.UseCase.Delete(r.Context(), &request); err != nil {
		c.Log.Printf("Failed to get grade: %v", err)
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
