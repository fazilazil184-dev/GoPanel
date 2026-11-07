package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fazilazil184-dev/GoPanel/views"
)

type AuthController struct {
	DB       *sql.DB
	renderer *views.TemplateRenderer
}

func NewAuthController(db *sql.DB, renderer *views.TemplateRenderer) *AuthController {
	return &AuthController{
		DB:       db,
		renderer: renderer,
	}
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	if err := ac.renderer.Render(w, "login.html", nil); err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	data := map[string]any{
		"Title": "Register",
		"Error": "",
	}

	views.RenderTemplate(w, data, "register.html", "views/auth")
}
