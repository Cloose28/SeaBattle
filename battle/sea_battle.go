package battle

import (
	"errors"
	"strings"
)

const maxSize = 26
const minSize = 1

var ErrWrongSize = errors.New("matrix got wrong size")
var ErrWrongCoord = errors.New("wrong coord")
var ErrCellShot = errors.New("cell was shot before")

type SeaBattleGame struct {
	size   int
	matrix [][]cellStatus
	ships  map[int]*ship
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
	s.initMatrix(n)
	s.initStats()
	s.initShips()

	return nil
}

func (s *SeaBattleGame) InitShips(coordinates string) (err error) {
	split := strings.Split(coordinates, ",")
	defer func() {
		if err != nil {
			s.clearMatrix()
		}
	}()
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
	s.stats.setShipCount(len(s.ships))
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
	if lives > 0 {
		newShip := &ship{liveCells: lives}
		s.ships[nextShipId] = newShip
	}
}

func (s *SeaBattleGame) Shot(coord string) (ShotResult, error) {
	x, y, err := getCoords(coord)
	if err != nil {
		return ShotResult{}, err
	}

	if !s.isValidCoords(x, y) {
		return ShotResult{}, ErrWrongCoord
	}

	switch s.matrix[x][y] {
	case Empty:
		s.matrix[x][y] = Shot
		s.stats.addShot()
		return NewEmptyShot(), nil
	case Shot:
		return NewEmptyShot(), ErrCellShot
	default:
		shipId := s.matrix[x][y].Int()
		ship := s.ships[shipId]
		ship.shipShot()
		s.matrix[x][y] = Shot
		if ship.isAlive() {
			s.stats.addKnocked()
			return NewKnockedShot(), nil
		} else {
			delete(s.ships, shipId)
			s.stats.addDestroyed()
			isEnd := len(s.ships) == 0
			return NewDestroyedShot(isEnd), nil
		}
	}
}

func (s *SeaBattleGame) GetStat() Stats {
	return *s.stats
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
	s.clearMatrix()
	s.initStats()
}

func (s *SeaBattleGame) initStats() {
	s.stats = &Stats{}
}

func (s *SeaBattleGame) initShips() {
	s.ships = make(map[int]*ship)
}

func (s *SeaBattleGame) initMatrix(n int) {
	s.matrix = make([][]cellStatus, n)
	for i := 0; i < n; i++ {
		s.matrix[i] = make([]cellStatus, n)
	}
}

func (s *SeaBattleGame) clearMatrix() {
	if len(s.matrix) != s.size {
		return
	}
	for i := 0; i < s.size; i++ {
		for j := 0; j < s.size; j++ {
			s.matrix[i][j] = Empty
		}
	}
}
