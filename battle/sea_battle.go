package battle

import (
	"errors"
	"strings"
)

const maxSize = 26
const minSize = 1

var ErrWrongSize = errors.New("matrix got wrong size")
var ErrWrongCoord = errors.New("wrong coord")


type SeaBattleGame struct {
	size   int
	matrix [][]cellStatus
	ships map[int]*ship
	stats  *Stats
}

func CreateSeaBattle() *SeaBattleGame {
	return &SeaBattleGame{}
}

func (s *SeaBattleGame) CreateGame(n int) error {
	if n < minSize || n > maxSize {
		return ErrWrongSize
	}
	s.size = n
	s.matrix = make([][]cellStatus, n)
	for i := 0; i < n; i++ {
		s.matrix[i] = make([]cellStatus, n)
	}
	s.stats = &Stats{}
	s.ships = make(map[int]*ship)

	return nil
}

func (s *SeaBattleGame) InitShips(coordinates string) error {
	split := strings.Split(coordinates, ",")
	for _, value := range split {
		coords := strings.Split(value, " ")
		if len(coords) != 2 {
			return ErrWrongCoord
		}
		x, y, err := getCoords(coords[0])
		if err != nil {
			return err
		}
		xEnd, yEnd, err := getCoords(coords[1])
		if err != nil {
			return err
		}
		if !s.isValidCoords(x, y) || !s.isValidCoords(xEnd, yEnd) {
			return ErrWrongCoord
		}
		s.setShipOnMatrix(x, y, xEnd, yEnd)
	}
	return nil
}

func (s *SeaBattleGame) setShipOnMatrix(x, y, xEnd, yEnd int) {
	nextShipId := len(s.ships) + 1
	lives := 0
	for i := x; i <= xEnd; i++ {
		for j := y; j <= yEnd; j++ {
			s.matrix[i][j] = cellStatus(nextShipId)
			lives++
		}
	}
	newShip := &ship{liveCells:lives}
	s.ships[nextShipId] = newShip
}

func (s *SeaBattleGame) Shot(coord string) (ShotResult, error) {
	x, y, err := getCoords(coord)
	if err != nil {
		return ShotResult{}, err
	}

	if !s.isValidCoords(x, y) {
		return ShotResult{}, ErrWrongCoord
	}



	return ShotResult{}, nil
}

func getCoords(coord string) (int, int, error) {
	if len(coord) != 2 {
		return 0, 0, ErrWrongCoord
	}
	return int(coord[0] - '1'), int(coord[1] - 'A'), nil
}

func (s *SeaBattleGame) isValidCoords(x, y int) bool {
	return x >= 0 && x < s.size && y >= 0 && y < s.size
}

func (s *SeaBattleGame) Clear() {
	s.size = 0
	s.stats = &Stats{}
}
