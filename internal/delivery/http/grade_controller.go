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
