package battle

// Stats provides current game parameters
type Stats struct {
	ShipCount int
	Destroyed int
	Knocked   int
	ShotCount int
}

func (s *Stats) setShipCount(count int) {
	s.ShipCount = count
}

func (s *Stats) addShot() {
	s.ShotCount++
}

func (s *Stats) addDestroyed() {
	s.Destroyed++
}

func (s *Stats) addKnocked() {
	s.Knocked++
}
