package driving

import "github.com/rs/zerolog/log"

// State maintains driver state as well as the concept of time along a route
type State struct {
	deliveryClock   int
	currentLocation Stop
	sortedStops     *DriverRoute
}

func newRouteState() *State {
	s := new(State)
	s.sortedStops = blankDriverRoute()
	s.currentLocation = Stop{}
	s.deliveryClock = 0
	return s
}

// AddStop adds route stops at a higher level and lets the logic be
// handled at the route list level
// TODO: make thread safe
func (s *State) AddStop(stops []Stop) ([]InternalStop, int) {
	added := make([]InternalStop, len(stops))
	for i := range stops {
		addContext := InternalStop{Actual: stops[i], Offset: i}
		log.Info().Msgf("addcontext %v", addContext)
		added[i] = s.sortedStops.AddSortedStop(addContext)
	}
	return added, 0
}
