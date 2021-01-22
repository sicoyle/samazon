package driving

type SubRoute struct {
	NextStop    Stop
	MilesAway   int
	EstDelivery int
}

type Stop struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type StopRequest struct {
	Route []Stop `json:"route"`
}

type InternalStop struct {
	Actual Stop `json:"stop"`
	Offset int  `json:"offset"`
}
