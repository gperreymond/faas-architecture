package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Serve static assets directly.
	r.PathPrefix("/company/static/").Handler(http.StripPrefix("/company/static/", http.FileServer(http.Dir("public/dist/"))))
	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	r.PathPrefix("/company").HandlerFunc(IndexHandler("public/index.html"))
	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, r),
		Addr:         "0.0.0.0:8088",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// IndexHandler : handler index
func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}
