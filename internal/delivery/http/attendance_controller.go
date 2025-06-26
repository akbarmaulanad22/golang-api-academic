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

type AttendanceController struct {
	Log     *logrus.Logger
	UseCase *usecase.AttendanceUseCase
}

func NewAttendanceController(useCase *usecase.AttendanceUseCase, logger *logrus.Logger) *AttendanceController {
	return &AttendanceController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *AttendanceController) AttendStudent(w http.ResponseWriter, r *http.Request) {

	auth := middleware.GetUser(r)

	// Set header content-type sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Menutup body setelah selesai dibaca
	defer r.Body.Close()

	// Parsing request body
	var request *model.AttendanceCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Printf("Failed to parse request body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Buat request untuk use case
	request.UserId = auth.ID

	// Panggil UseCase.Logout
	response, err := c.UseCase.AttendStudent(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to attend user: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.AttendanceResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func (c *AttendanceController) AttendLecturer(w http.ResponseWriter, r *http.Request) {

	auth := middleware.GetUser(r)

	// Set header content-type sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Menutup body setelah selesai dibaca
	defer r.Body.Close()

	// Parsing request body
	var request *model.AttendanceCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Printf("Failed to parse request body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Buat request untuk use case
	request.UserId = auth.ID

	// Panggil UseCase.Logout
	response, err := c.UseCase.AttendLecturer(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to attend user: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.AttendanceResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func (c *AttendanceController) ListByUserID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	npm := vars["npm"]

	npmInt, err := strconv.Atoi(npm)
	if err != nil {
		c.Log.Printf("Failed to parse npm: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	request := &model.ListAttendanceRequest{
		Npm: uint(npmInt),
	}

	// Panggil UseCase.Logout
	response, err := c.UseCase.ListByUserID(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to attend user: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.AttendanceResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func (c *AttendanceController) ListByCourseCodeAndNpm(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	npm := vars["npm"]

	npmInt, err := strconv.Atoi(npm)
	if err != nil {
		c.Log.Printf("Failed to parse npm: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	courseCode := vars["courseCode"]

	request := &model.ListInLecturerAttendanceRequest{
		Npm:        uint(npmInt),
		CourseCode: courseCode,
	}

	// Panggil UseCase.Logout
	response, err := c.UseCase.ListByCourseCodeAndNpm(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to attend user: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[[]model.AttendanceResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func (c *AttendanceController) Update(w http.ResponseWriter, r *http.Request) {

	// Set header content-type sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Menutup body setelah selesai dibaca
	defer r.Body.Close()

	// Parsing request body
	var request *model.AttendanceUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Printf("Failed to parse request body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Log.Printf("Failed to parse id: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	request.ID = uint(idInt)
	// Panggil UseCase.Logout
	response, err := c.UseCase.Update(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to attend user: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.AttendanceResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
