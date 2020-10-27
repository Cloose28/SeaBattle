package battle

import "errors"

const maxSize = 26
const minSize = 1

var ErrWrongSize = errors.New("matrix got wrong size")

type SeaBattleGame struct {
	size   int
	matrix [][]cellStatus
	stats  *Stats
}

func CreateSeaBattle() *SeaBattleGame {
	return &SeaBattleGame{}
}

type cellStatus struct {
	cell int8
}

func (s *SeaBattleGame) CreateGame(n int) error {
	if n < minSize || n > maxSize {
		return ErrWrongSize
	}
	s.size = n
	s.matrix = make([][]cellStatus, n)
	s.stats = &Stats{}

	return nil
}

func (s *SeaBattleGame) Clear() {
	s.size = 0
	s.stats = &Stats{}
}
