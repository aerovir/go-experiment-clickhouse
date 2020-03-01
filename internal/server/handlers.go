package server

import "net/http"

func (s *Server) YourOwnHandler(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]string)

	// ToDo: Add your requests to ClickHouse

	s.sendJSON(w, http.StatusOK, result)
}

func (s *Server) PingHandler(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]string)

	if err := s.DB.Ping(); err != nil {
		result["status"] = "failed"
		result["error"] = err.Error()
		s.sendJSON(w, http.StatusInternalServerError, result)
		return
	}

	result["status"] = "OK"
	s.sendJSON(w, http.StatusOK, result)
}
