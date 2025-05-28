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

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var request model.UserRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Panggil use case untuk membuat user
	response, err := c.UseCase.Create(r.Context(), &request)
	if err != nil {
		c.Log.Warnf("Failed to register user: %+v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err)) // helper untuk menentukan status code
		return
	}

	// Kirim response JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.UserResponse]{Data: response}); err != nil {
		c.Log.Warnf("Failed to write response: %+v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	// Set header content-type sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Menutup body setelah selesai dibaca
	defer r.Body.Close()

	// Parsing request body
	var request model.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.Printf("Failed to parse request body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Memanggil UseCase untuk login
	response, err := c.UseCase.Login(r.Context(), &request)
	if err != nil {
		c.Log.Printf("Failed to login user: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err)) // Gunakan helper status code
		return
	}

	// Mengirimkan response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[*model.UserResponse]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	auth := middleware.GetUser(r)

	// Buat request untuk use case
	request := &model.UserLogoutRequest{
		Username: auth.Username,
	}

	// Panggil UseCase.Logout
	response, err := c.UseCase.Logout(r.Context(), request)
	if err != nil {
		c.Log.Printf("Failed to logout user: %v", err)
		http.Error(w, err.Error(), helper.GetStatusCode(err))
		return
	}

	// Set header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Kirim response sukses
	if err := json.NewEncoder(w).Encode(model.WebResponse[bool]{Data: response}); err != nil {
		c.Log.Printf("Failed to write response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
