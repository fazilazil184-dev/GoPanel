package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fazilazil184-dev/GoPanel/internal/controllers"
	"github.com/fazilazil184-dev/GoPanel/views"
)

func RegisterAuthRoutes(db *sql.DB) http.Handler {
	router := http.NewServeMux()

	renderer, err := views.NewTemplateRenderer("views/auth")
	if err != nil {
		log.Fatalf("Failed to initialize templates: %v", err)
	}

	authController := controllers.NewAuthController(db, renderer)

	router.HandleFunc("/login", authController.Login)
	router.HandleFunc("/register", authController.Register)

	return router
}
