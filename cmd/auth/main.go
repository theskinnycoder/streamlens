package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/theskinnycoder/streamlens/internal/auth/handler"
	"github.com/theskinnycoder/streamlens/internal/auth/router"
	"github.com/theskinnycoder/streamlens/internal/auth/service"
	"github.com/theskinnycoder/streamlens/internal/cookies"
	"github.com/theskinnycoder/streamlens/internal/db/postgres"
	"github.com/theskinnycoder/streamlens/internal/hashing"
	"github.com/theskinnycoder/streamlens/internal/jwt"
	"github.com/theskinnycoder/streamlens/internal/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()

	portString := os.Getenv("AUTH_SERVICE_PORT")
	if portString == "" {
		portString = "8080"
	}

	conn, err := postgres.NewConnection(ctx)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer conn.Close(ctx)

	hashingService := hashing.NewHashingService(12)
	jwtService := jwt.NewJWTService(os.Getenv("JWT_SECRET"))
	cookieService := cookies.NewCookieService(os.Getenv("COOKIE_SECRET"))
	repo := repository.New(conn)
	authService := service.NewAuthService(repo, hashingService, jwtService)
	authHandler := handler.NewAuthHandler(authService, cookieService)

	v1AuthRouter := router.NewAuthRouter(*authHandler)

	log.Printf("Starting server on :%s", portString)
	http.ListenAndServe(":"+portString, v1AuthRouter)
}
