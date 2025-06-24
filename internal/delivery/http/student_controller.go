package http

import (
	"encoding/json"
	"net/http"
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
