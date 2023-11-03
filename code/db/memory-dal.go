package db

import (
	"fmt"
	"mural/controller/mural/cons"
	"mural/controller/mural/service"
	"mural/middleware"
	"mural/model"

	"github.com/bluele/gcache"
)


type MemoryDAL struct {
	Cache gcache.Cache
}

func NewMemoryDAL() MemoryDAL {
	gc := gcache.New(20).LRU().Build()

	return MemoryDAL{
		Cache: gc,
	}
}

func (dal MemoryDAL) InitGames () error {
	current_games := map[string]model.Game{}


	for i := 0; i < 1_000_000; i++ {
		board := service.NewGameBoard(cons.BoardSize)
		correct_movie, answers := service.NewAnswers()

		game_key := middleware.HashString(fmt.Sprintf("random_key_%d", i))
		current_game := model.Game{
			GameKey: game_key,
			CurrentScore: cons.InititalScore,
			Board: *board,
			TodayAnswer: correct_movie,
			Answers: answers,
			GameState: model.GAME_INIT,
		}

		current_games[game_key] = current_game
	}

	err := dal.Cache.Set("all_games", current_games)
	if err != nil {
		return ErrSettingCurrentGames
	}
	return nil
}


func (dal MemoryDAL) GetCurrentGames (
)(map[string]model.Game, error) {
	current_games_interface, err := dal.Cache.Get("all_games")
	if err != nil {
		return nil, ErrGettingBoardFromDB
	}


	// cast interface
	current_games, ok := current_games_interface.(map[string]model.Game)
	if !ok {
		return nil, ErrCastingBoard
	}

	return current_games, nil
}

func (dal MemoryDAL) GetCurrentGame(
	game_key string,
) (*model.Game, error) {
	current_games, err := dal.GetCurrentGames()
	if err != nil {
		return nil, err

	}

	current_game, ok := current_games[game_key]
	if ! ok {
		board := service.NewGameBoard(cons.BoardSize)
		correct_movie, answers := service.NewAnswers()

		current_game = model.Game{
			GameKey: game_key,
			CurrentScore: cons.InititalScore,
			Board: *board,
			TodayAnswer: correct_movie,
			Answers: answers,
			GameState: model.GAME_INIT,
		}
	}


	return &current_game, nil
}

func (dal MemoryDAL) SetCurrentGame(current_game model.Game) error {
	current_games, err := dal.GetCurrentGames()
	if err != nil {
		return err
	}

	current_games[current_game.GameKey] = current_game

	return dal.Cache.Set("all_games", current_games)
}