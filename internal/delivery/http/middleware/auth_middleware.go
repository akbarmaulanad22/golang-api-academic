package middleware

import (
	"context"
	"net/http"
	"tugasakhir/internal/model"
	"tugasakhir/internal/usecase"

	"github.com/gorilla/mux"
)

type contextKey string

const (
	AuthKey contextKey = "auth"
)

func NewAuth(userUseCase *usecase.UserUseCase) mux.MiddlewareFunc {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token := r.Header.Get("Authorization")
			if token == "" {
				token = "NOT_FOUND"
			}
			userUseCase.Log.Debugf("Authorization : %s", token)

			request := &model.VerifyUserRequest{Token: token}
			auth, err := userUseCase.Verify(r.Context(), request)
			if err != nil {
				userUseCase.Log.Warnf("Failed find user by token : %+v", err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userUseCase.Log.Debugf("User : %+v", auth.Username)

			// Masukkan auth ke context
			ctx := context.WithValue(r.Context(), AuthKey, auth)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func GetUser(r *http.Request) *model.Auth {
	if auth, ok := r.Context().Value(AuthKey).(*model.Auth); ok {
		return auth
	}
	return nil
}
