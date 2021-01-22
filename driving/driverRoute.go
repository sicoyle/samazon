package driving

import "github.com/rs/zerolog/log"

// DriverRoute maintains the route stop information and
// its time context within the route simulation
type DriverRoute struct {
	Data []*InternalStop
}

// blankDriverRoute creates a new route list
func blankDriverRoute() *DriverRoute {
	return &DriverRoute{Data: make([]*InternalStop, 0)}
}

//AddSortedStop adds stops to route in sorted order with external context
// TODO: sort route based on its location
func (r *DriverRoute) AddSortedStop(s InternalStop) InternalStop {
	log.Info().Msgf("here is s %v", s)
	// TODO: Add logic to maintain sorted order
	r.Data = append(r.Data, nil)
	copy(r.Data[s.Offset+1:], r.Data[s.Offset:])
	r.Data[s.Offset] = &s
	log.Info().Msgf("r.Data %v", *r.Data[s.Offset])

	return s
}

// TODO: Clear route

// TODO: Clear route stops we delivered at already
