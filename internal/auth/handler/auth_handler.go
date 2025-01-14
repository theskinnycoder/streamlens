package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/theskinnycoder/streamlens/internal/auth/dto"
	"github.com/theskinnycoder/streamlens/internal/auth/service"
	"github.com/theskinnycoder/streamlens/internal/cookies"
)

type AuthHandler struct {
	authService   *service.AuthService
	cookieService *cookies.CookieService
}

func NewAuthHandler(service *service.AuthService, cookieService *cookies.CookieService) *AuthHandler {
	return &AuthHandler{authService: service, cookieService: cookieService}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// 1. Decode request body
	var body dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Register user
	response, err := h.authService.Register(r.Context(), body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Set refresh token cookie
	h.cookieService.SetCookie(w, "refresh_token", response.RefreshToken, time.Now().Add(time.Hour*24*7))

	// 4. Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.AuthResponseWithoutRefreshToken{
		Message: response.Message,
		User:    response.User,
		Token:   response.AccessToken,
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// 1. Decode request body
	var body dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Login user
	response, err := h.authService.Login(r.Context(), body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Set refresh token cookie
	h.cookieService.SetCookie(w, "refresh_token", response.RefreshToken, time.Now().Add(time.Hour*24*7))

	// 4. Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.AuthResponseWithoutRefreshToken{
		Message: response.Message,
		User:    response.User,
		Token:   response.AccessToken,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// 1. Delete refresh token cookie
	h.cookieService.DeleteCookie(w, "refresh_token")

	// 2. Return response
	w.WriteHeader(http.StatusOK)
}
