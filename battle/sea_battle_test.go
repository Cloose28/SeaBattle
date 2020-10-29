package battle

import "testing"

func TestSeaBattleGame_Clear(t *testing.T) {
	battle := CreateSeaBattle()
	battle.Clear()
}

func TestSeaBattleGame_CreateGame(t *testing.T) {
	battle := CreateSeaBattle()
	err := battle.CreateGame(1)
	if err != nil {
		t.Errorf("can't create new game %v", err)
	}
}

func TestSeaBattleGame_WrongCreateGame(t *testing.T) {
	battle := CreateSeaBattle()
	err := battle.CreateGame(minSize - 1)
	if err != ErrWrongSize {
		t.Errorf("not rise error on wrong data %v", err)
	}
}

func TestSeaBattleGame_WrongMaxCreateGame(t *testing.T) {
	battle := CreateSeaBattle()
	err := battle.CreateGame(maxSize + 1)
	if err != ErrWrongSize {
		t.Errorf("not rise error on wrong data %v", err)
	}
}

func TestGetCoords(t *testing.T) {
	const coord = "1A"
	x, y, err := getCoords(coord)
	if err != nil {
		t.Errorf("can't get coord %s, err %v", coord, err)
	}
	if x != 0 || y != 0 {
		t.Errorf("got wrong coord %d, %d expected 0, 0 by string %s", x, y, coord)
	}
}

func TestGetWrongCoords(t *testing.T) {
	const coord = "DFSF"
	_, _, err := getCoords(coord)
	if err != ErrWrongCoord {
		t.Errorf("got coords for wrong data %s, err %v", coord, err)
	}
}

func TestValidCoords(t *testing.T) {
	const x, y = 0, 0
	battle := CreateSeaBattle()
	const battleSize = 1
	battle.CreateGame(battleSize)
	result := battle.isValidCoords(x, y)
	if !result {
		t.Errorf("coords %d %d not valid for battle size %d", x, y, battleSize)
	}
}

func TestNotValidXCoords(t *testing.T) {
	const x, y = 1, 0
	battle := CreateSeaBattle()
	const battleSize = 1
	battle.CreateGame(battleSize)
	result := battle.isValidCoords(x, y)
	if result {
		t.Errorf("coords %d %d not valid for battle size %d", x, y, battleSize)
	}
}

func TestNotValidYCoords(t *testing.T) {
	const x, y = 0, 1
	battle := CreateSeaBattle()
	const battleSize = 1
	battle.CreateGame(battleSize)
	result := battle.isValidCoords(x, y)
	if result {
		t.Errorf("coords %d %d not valid for battle size %d", x, y, battleSize)
	}
}

func TestAddOneShipOnMatrix(t *testing.T) {
	battle := createBattleWithSize(2)
	const x, y, xEnd, yEnd = 0, 0, 1, 1
	const expectedLives = 2 * (xEnd - x + yEnd - y)
	battle.setShipOnMatrix(x, y, xEnd, yEnd)
	expectedShipId := battle.matrix[0][0]
	if _, ok := battle.ships[expectedShipId.Int()]; !ok {
		t.Errorf("ship was not added to game")
	}
	if battle.ships[expectedShipId.Int()].liveCells != expectedLives {
		t.Errorf("ship has wrong lives count")
	}
	if !isShipOnMap(x, y, xEnd, yEnd, expectedShipId.Int(), battle) {
		t.Errorf("ship was not added to map")
	}
}

func isShipOnMap(x, y, xEnd, yEnd, shipId int, game *SeaBattleGame) bool {
	for i := x; i <= xEnd; i++ {
		for j := y; j <= yEnd; j++ {
			if game.matrix[i][j].Int() != shipId {
				return false
			}
		}
	}
	return true
}

func createBattleWithSize(size int) *SeaBattleGame {
	battle := CreateSeaBattle()
	battle.CreateGame(size)
	return battle
}

