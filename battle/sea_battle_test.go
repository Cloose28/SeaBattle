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
		return
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
	const coordinates = "1A 2B"
	const expectedLives = 2 * (xEnd - x + yEnd - y)
	err := battle.InitShips(coordinates)
	if err != nil {
		t.Errorf("can't init ships %v", err)
		return
	}
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

func TestAddTwoShipOnMatrix(t *testing.T) {
	battle := createBattleWithSize(3)
	const x, y, xEnd, yEnd = 0, 0, 1, 0
	const coordinates = "1A 2A,1C 1C"
	const expectedLives1 = 2 * (xEnd - x + yEnd - y)
	err := battle.InitShips(coordinates)
	if err != nil {
		t.Errorf("can't init ships %v", err)
		return
	}
	expectedShipId := battle.matrix[0][0]
	if _, ok := battle.ships[expectedShipId.Int()]; !ok {
		t.Errorf("ship was not added to game")
	}
	if battle.ships[expectedShipId.Int()].liveCells != expectedLives1 {
		t.Errorf("ship has wrong lives count")
	}
		if !isShipOnMap(x, y, xEnd, yEnd, expectedShipId.Int(), battle) {
		t.Errorf("ship was not added to map")
	}
}

func TestAddWrongXYShipOnMatrix(t *testing.T) {
	battle := createBattleWithSize(2)
	const x, y, xEnd, yEnd = 1, 1, 0, 0
	battle.setShipOnMatrix(x, y, xEnd, yEnd)
	isEmptyMap := len(battle.ships) == 0
	if !isEmptyMap {
		t.Errorf("ship was added to game with wrong coords")
	}
	if !isShipOnMap(xEnd, yEnd, x, y, Empty.Int(), battle) {
		t.Errorf("ship was added to map with wrong coords")
	}
}

func TestShipDestroying(t *testing.T) {
	battle := createBattleWithSize(2)
	const x, y, xEnd, yEnd = 0, 0, 0, 0
	const coord = "1A"
	battle.setShipOnMatrix(x, y, xEnd, yEnd)
	expectedResult := NewDestroyedShot(true)
	shot, err := battle.Shot(coord)
	if err != nil {
		t.Errorf("error on ship shooting %v", err)
		return
	}
	if shot != expectedResult {
		t.Errorf("wrong shot result")
	}
}

func TestShipKnocking(t *testing.T) {
	battle := createBattleWithSize(2)
	const x, y, xEnd, yEnd = 0, 0, 0, 1
	const coord = "1A"
	battle.setShipOnMatrix(x, y, xEnd, yEnd)
	expectedResult := NewKnockedShot()
	shot, err := battle.Shot(coord)
	if err != nil {
		t.Errorf("error on ship shooting %v", err)
		return
	}
	if shot != expectedResult {
		t.Errorf("wrong shot result")
	}
}

func TestShotEmpty(t *testing.T) {
	battle := createBattleWithSize(2)
	const coord = "1A"
	expectedResult := NewEmptyShot()
	shot, err := battle.Shot(coord)
	if err != nil {
		t.Errorf("error on empty cell shooting %v", err)
		return
	}
	if shot != expectedResult {
		t.Errorf("wrong shot result")
	}
}

func TestShotCellTwice(t *testing.T) {
	battle := createBattleWithSize(2)
	const coord = "1A"
	_, err := battle.Shot(coord)
	if err != nil {
		t.Errorf("error on empty cell shooting %v", err)
		return
	}
	_, err = battle.Shot(coord)
	if err != ErrCellShoted {
		t.Errorf("cell was shooted twice %v", err)
	}
}

func TestStatInit(t *testing.T) {
	battle := createBattleWithSize(2)
	stat := battle.GetStat()
	expectedStat := Stats{
		ShipCount: 0,
		Destroyed: 0,
		Knocked:   0,
		ShotCount: 0,
	}
	if stat != expectedStat {
		t.Errorf("wrong initial stat %v", stat)
	}
}

func TestStatShipCount(t *testing.T) {
	battle := createBattleWithSize(3)
	battle.InitShips("1A 1A,2B 2B")
	stat := battle.GetStat()
	expectedStat := Stats{
		ShipCount: 2,
		Destroyed: 0,
		Knocked:   0,
		ShotCount: 0,
	}
	if stat != expectedStat {
		t.Errorf("wrong stat with ships %v", stat)
	}
}

func TestStatShipDestroying(t *testing.T) {
	battle := createBattleWithSize(3)
	battle.InitShips("1A 1A,2B 2B")
	battle.Shot("1A")
	stat := battle.GetStat()
	expectedStat := Stats{
		ShipCount: 2,
		Destroyed: 1,
		Knocked:   0,
		ShotCount: 1,
	}
	if stat != expectedStat {
		t.Errorf("wrong stat with ships %v", stat)
	}
}

func TestStatShipKnocked(t *testing.T) {
	battle := createBattleWithSize(3)
	battle.InitShips("1A 2A,3B 3B")
	battle.Shot("1A")
	stat := battle.GetStat()
	expectedStat := Stats{
		ShipCount: 2,
		Destroyed: 0,
		Knocked:   1,
		ShotCount: 1,
	}
	if stat != expectedStat {
		t.Errorf("wrong stat with ships %v", stat)
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
