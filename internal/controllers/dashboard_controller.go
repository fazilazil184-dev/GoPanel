package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fazilazil184-dev/GoPanel/views"
)

type DashboardController struct {
	DB *sql.DB
}

func NewDashboardController(db *sql.DB) *DashboardController {
	return &DashboardController{DB: db}
}

func (dc *DashboardController) Dashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	renderer, err := views.NewTemplateRenderer("views/dashboard")
	if err != nil {
		log.Fatalf("Failed to initialize templates: %v", err)
	}

	renderer.Render(w, "index.html", nil)
}
