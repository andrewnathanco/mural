package db

import (
	"fmt"
	"mural/model"

	"github.com/bluele/gcache"
)

var (
	DAL IDAL

	// errors
	ErrGettingBoardFromDB = fmt.Errorf("could not get board from db")
	ErrBoardNotFound = fmt.Errorf("no board set")
	ErrCastingBoard = fmt.Errorf("board not in correct format")
)

type IDAL interface {
	GetCurrentGame() (*model.Game, error)
	SetCurrentGame(model.Game)  error
} 

type MemoryDAL struct {
	Cache gcache.Cache
}

func NewMemoryDAL() MemoryDAL {
	gc := gcache.New(20).LRU().Build()

	return MemoryDAL{
		Cache: gc,
	}
}

func (dal MemoryDAL) GetCurrentGame() (*model.Game, error) {
	current_game_interface, err := dal.Cache.Get("current-game")
	if err != nil {
		return nil, ErrGettingBoardFromDB
	}

	// cast interface
	current_game, ok := current_game_interface.(model.Game)
	if !ok {
		return nil, ErrCastingBoard
	}

	return &current_game, nil
}

func (dal MemoryDAL) SetCurrentGame(current_game model.Game) error {
	return dal.Cache.Set("current-game", current_game)
}