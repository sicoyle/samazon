package server

import "net/http"

// Server will hold the router info
type Server struct {
	Router *http.ServeMux
}

// NewServer creates the server type
func NewServer(restPort int) *Server {
	s := &Server{
		Router: http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
