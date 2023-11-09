package worker

import (
	"database/sql"
	"fmt"
	"log/slog"
	"math/rand"
	"mural/controller/mural/service"
	"mural/db"
	"mural/model"
	"time"
)

type MuralWorker struct {}

func NewMuralWorker() MuralWorker {
	return MuralWorker{}
}


func (mw MuralWorker) ResetGameSessions() {
	// we need to wait till we've actually created the dal to do this
	// only edge case where this matters is if we restart the server exactly at midnight EST
	slog.Info("Resetting game sessions.")
	err := db.DAL.ResetGameSessions()
	if err != nil {
		slog.Error(fmt.Errorf("could not reset game sessions: %w", err).Error())
		return
	}
}


func getSQLDecade() string {
	decade := service.GetCurrentDecade()
	currentDay := time.Now().Weekday()

	decade_sql := ""
	if currentDay == time.Sunday {
		decade_sql += "%"
	} else {
		decade_sql += replaceLastCharacter(decade, '%')
	}
	
	return decade_sql
}

func (mw MuralWorker) SetupNewGame() {
	slog.Info("Setting up game")

	// first we need to get the current game info
	current_game, curr_game_err := db.DAL.GetCurrentGameInfo()
	if curr_game_err != sql.ErrNoRows {
		if curr_game_err  != nil  {
			slog.Error(fmt.Errorf("could not get game info: %w", curr_game_err).Error())
			return
		}
	}

	// get new answers from 
	answers, err := db.DAL.GetRandomAnswers(getSQLDecade())
	if err != nil {
		slog.Error(fmt.Errorf("could not get random answers: %w", err).Error())
		return
	}
	
	var correct_answer model.Answer
	for _, answer := range answers {
		if answer.IsCorrect {
			correct_answer = answer
		}
	}

	randomizeAnswers(answers)

	var game_key int
	if curr_game_err == sql.ErrNoRows {
		game_key = 1
	} else {
		game_key = current_game.GameKey + 1
	}

	current_date := time.Now().Format("2022/10/10")
	new_game := model.Game{
		Date: current_date,
		Answers: answers,
		CorrectAnswer: correct_answer,
		GameKey: game_key,
		IsCurrent: true,
	}

	// set new game
	err = db.DAL.SetNewCurrentGame(new_game)
	if err != nil {
		slog.Error(fmt.Errorf("could not set current game: %w", err).Error())
		return
	}

	// now lets redlist the answer
	err = db.DAL.RedlistAnswer(correct_answer)
	if err != nil {
		slog.Error(fmt.Errorf("could not redlist answer: %w", err).Error())
		return
	}
}

func randomizeAnswers(a []model.Answer) {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}
func replaceLastCharacter(inputString string, newChar rune) string {
	if len(inputString) == 0 {
		return inputString // Return the original string if it's empty
	}

	// Convert the string to a rune slice to work with individual characters
	strRunes := []rune(inputString)

	// Update the last character
	strRunes[len(strRunes)-1] = newChar

	// Convert the rune slice back to a string
	return string(strRunes)
}
