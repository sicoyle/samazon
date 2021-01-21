package server

import "net/http"

func (s *Server) routes() {
	s.Router.Handle("/status", s.handleStatus())
}

// handleStatus returns driver route status information
func (s *Server) handleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse client parameters if any and perform checks on data

		// Do logic to gather response data

		w.Write([]byte("success"))
		w.WriteHeader(http.StatusOK)
	}
}
