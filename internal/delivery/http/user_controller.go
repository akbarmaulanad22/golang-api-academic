package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tugasakhir/internal/delivery/http/middleware"
	"tugasakhir/internal/model"
	"tugasakhir/internal/usecase"

	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	Usecase *usecase.UserUseCase
}

func NewUserController(log *logrus.Logger, usecase *usecase.UserUseCase) *UserController {

	return &UserController{
		Log:     log,
		Usecase: usecase,
	}

}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	request := new(model.UserRegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.WithError(err).Warnf("Failed to get body request")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.WebResponse[*model.UserResponse]{Data: nil, Errors: "Body request not valid"})
		return
	}

	userResponse, err := c.Usecase.Create(r.Context(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to register user")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.WebResponse[*model.UserResponse]{Data: nil, Errors: "Failed to register"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.WebResponse[*model.UserResponse]{Data: userResponse})

}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {

	request := new(model.UserLoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		c.Log.WithError(err).Warnf("Failed to get body request")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.WebResponse[*model.UserResponse]{Data: nil, Errors: "Body request not valid"})
		return
	}

	userResponse, err := c.Usecase.Login(r.Context(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to login user")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.WebResponse[*model.UserResponse]{Data: nil, Errors: fmt.Sprintf("Failed to Login: %+v", err)})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.WebResponse[*model.UserResponse]{Data: userResponse})
}

func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	auth := middleware.GetUser(r)

	request := &model.UserLogoutRequest{
		Username: auth.Username,
	}

	response, err := c.Usecase.Logout(r.Context(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to logout user")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.WebResponse[bool]{Data: false, Errors: fmt.Sprintf("Failed to Logout: %+v", err)})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.WebResponse[bool]{Data: response})

}
