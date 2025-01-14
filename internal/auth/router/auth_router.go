package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/theskinnycoder/streamlens/internal/auth/handler"
)

func NewAuthRouter(handler handler.AuthHandler) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.SetHeader("Content-Type", "application/json"))
	router.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	router.Use(middleware.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS"))
	router.Use(middleware.SetHeader("Access-Control-Allow-Credentials", "true"))
	router.Use(middleware.StripSlashes)

	router.Post("/auth/v1/register", handler.Register)
	router.Post("/auth/v1/login", handler.Login)
	router.Post("/auth/v1/logout", handler.Logout)

	return router
}
