package http

import (
	"encoding/json"
	"net/http"
	"tugasakhir/internal/delivery/http/middleware"
	"tugasakhir/internal/helper"
	"tugasakhir/internal/model"
	"tugasakhir/internal/usecase"

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
