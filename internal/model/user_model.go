package model

import "time"

type UserResponse struct {
	Token     string    `json:"token,omitempty"`
	Username  string    `json:"username,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserRegisterRequest struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required,min=6"`
	RoleID   uint   `json:"role_id,omitempty" validate:"required"`
}
type UserLoginRequest struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required,min=6"`
}

type UserDeleteRequest struct {
	Username string `json:"username,omitempty" validate:"required"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}

type UserLogoutRequest struct {
	Username string `json:"username" validate:"required,max=100"`
}

type GetUserRequest struct {
	Username string `json:"username,omitempty" validate:"required"`
}
