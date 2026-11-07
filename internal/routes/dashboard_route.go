package routes

import (
	"database/sql"
	"net/http"

	"github.com/fazilazil184-dev/GoPanel/internal/controllers"
)

func RegisterDashboardRoutes(db *sql.DB) http.Handler {
	dashboard := http.NewServeMux()

	dashboardController := controllers.NewDashboardController(db)

	dashboard.HandleFunc("/dashboard", dashboardController.Dashboard)

	return dashboard
}
