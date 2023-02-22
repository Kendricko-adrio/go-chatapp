package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/kendricko-adrio/go-ws/entity"
	"github.com/kendricko-adrio/go-ws/session"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := CheckCookie(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&entity.Response{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", user)
		next.ServeHTTP(w, r)
	})
}

func CheckCookie(w http.ResponseWriter, r *http.Request) (entity.User, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return entity.User{}, errors.New("cookie not found")
	}
	sessionValue := cookie.Value
	if sessionValue == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return entity.User{}, errors.New("session value is empty")
	}
	userSession, ok := session.GetSession(sessionValue)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return entity.User{}, errors.New("user not found")
	}
	if userSession.Expired.Before(time.Now()) {
		session.DeleteUserSession(sessionValue)
		w.WriteHeader(http.StatusUnauthorized)
		return entity.User{}, errors.New("session expired")
	}

	return userSession.User, nil
}
