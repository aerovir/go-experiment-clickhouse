package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	DB *sqlx.DB
}

func InitServer(db *sqlx.DB) *Server {
	server := &Server{db}

	return server
}

func (s *Server) Start() error {
	return http.ListenAndServe(":8080", s.getRouter())
}

func (s *Server) getRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggerMiddleware)
	r.HandleFunc("/", s.PingHandler)

	return r
}

func (s *Server) sendJSON(w http.ResponseWriter, status int, data interface{}) {
	encodedData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to encode JSON: %s", err)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(encodedData)))
	fmt.Fprint(w, string(encodedData))
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", time.Now().Format(time.RFC3339), r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
