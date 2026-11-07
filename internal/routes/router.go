package routes

import (
	"database/sql"
	"net/http"
)

func SetupRouter(db *sql.DB) *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("/", home)

	router.Handle("/auth/", http.StripPrefix("/auth", RegisterAuthRoutes(db)))
	router.Handle("/dashboard", RegisterDashboardRoutes(db))

	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World\n"))
}
