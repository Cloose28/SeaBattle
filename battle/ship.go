package battle

type ship struct {
	liveCells int
}

func (s *ship) shipShoted() {
	s.liveCells--
}

func (s *ship) isAlive() bool {
	return s.liveCells > 0
}