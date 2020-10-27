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
	err := battle.CreateGame(-5)
	if err != ErrWrongSize {
		t.Errorf("not rise error on wrong data %v", err)
	}
}


