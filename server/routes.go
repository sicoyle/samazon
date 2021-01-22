package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sicoyle/samazon/driving"
)

func (s *Server) routes() {
	s.Router.Handle("/status", s.handleStatus())
	s.Router.Handle("/sim", s.handleSim())
}

// handleStatus is my init test endpoint
func (s *Server) handleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse client parameters if any and perform checks on data

		// Do logic to gather response data

		w.Write([]byte("success"))
		w.WriteHeader(http.StatusOK)
	}
}

// handleRouteList returns list of stops along the drivers route
func (s *Server) handleSim() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse client parameters and perform checks on data
		var stopReq driving.StopRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&stopReq)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("failed - HandleSim() - %v", err)))
			w.WriteHeader(http.StatusInternalServerError)
			panic(err)
		}
		if stopReq.Route == nil {
			w.Write([]byte("failed - HandleSim() - invalid route simulation data"))
			w.WriteHeader(http.StatusBadRequest)
		}

		stops, clock := s.DriversRoute.AddSimulatedRoute(stopReq.Route)

		w.Write([]byte(fmt.Sprintf("success, %v %v", stops, clock)))
		w.WriteHeader(http.StatusOK)
	}
}
