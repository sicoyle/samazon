package driving

import "github.com/rs/zerolog/log"

// DeliveryVan was created to maintain state at a high level
// DeliveryVan can be considered the heart of the route simulation
// and drive the time ticker for making stops along the route
// TODO: capture more of a sense of driver state and introduce time to the simulation
type DeliveryVan struct {
	state *State
	clock int
}

func NewDeliveryVan() *DeliveryVan {
	d := new(DeliveryVan)
	d.state = newRouteState()
	d.clock = 0
	return d
}

func (d *DeliveryVan) AddSimulatedRoute(add []Stop) (added []InternalStop, clock int) {
	log.Info().Msgf("stops to add: %v", add)
	return d.state.AddStop(add)
}
