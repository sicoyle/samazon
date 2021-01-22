package server

import (
	"net/http"

	"github.com/sicoyle/samazon/driving"
)

// Server will hold the router info
type Server struct {
	Router       *http.ServeMux
	DriversRoute *driving.DeliveryVan
	Simulation   []driving.Stop
}

// NewServer creates the server type
func NewServer(restPort int) *Server {
	s := &Server{
		Router:       http.NewServeMux(),
		DriversRoute: driving.NewDeliveryVan(),
		Simulation:   nil,
	}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
