package handlers

import (
	"fmt"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "Stranger"
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// HandleRequests returns handlers
func HandleRequests() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homePageHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/ready", readinessHandler)

	return mux
}
