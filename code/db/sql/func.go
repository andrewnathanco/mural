package sql

import (
	"database/sql"
	"encoding/json"
	"mural/model"
	"os"
)

func createFileIfNotExists(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// File does not exist, so create it
		_, err = os.Create(filename)
		return err
	}
	return nil
}

func insertSession(
	user_key string,
	session *model.Session,
	dal *SQLiteDAL,
) (error) {
	game_sessions_marshalled, err := json.Marshal(session)
	if err != nil {
		return err
	}


	_, err = dal.DB.Exec(upsertGameSession, user_key, string(game_sessions_marshalled))
	if err != nil {
		return err
	}

	return nil
}

func cacheAnswers(
	answers []model.Answer,
	dal *SQLiteDAL,
) (error) {
	for _, answer := range answers {
		answer_marshalled, err := json.Marshal(answer)
		if err != nil {
			return err
		}

		_, err = dal.DB.Exec(insertAnswers, answer.ID, string(answer_marshalled))
		if err != nil {
			return err
		}
	}

	return nil
}

func redlistAnswer(
	answer model.Answer,
	current_game model.Game,
	dal *SQLiteDAL,
) (error) {
	_, err := dal.DB.Exec(redlistAnswerQuery, answer.ID, current_game.Date, current_game.GameKey)
	if err != nil {
		return err
	}

	return nil
}

func setNewCurrentGame(
	game model.Game,
	dal *SQLiteDAL,
) (error) {
	_, err := dal.DB.Exec(closeCurrentGame)
	if err != nil {
		return err
	}


	game_marshalled, err := json.Marshal(game)
	if err != nil {
		return err
	}

	_, err = dal.DB.Exec(setNewGame, game.GameKey, game_marshalled)
	if err != nil {
		return err
	}

	return nil
}



func resetSessions(
	dal *SQLiteDAL,
) (error) {
	_, err := dal.DB.Exec(resetGameSessions)
	if err != nil {
		return err
	}

	return nil
}

func getSession(
	user_key string,
	dal *SQLiteDAL,
) (*model.Session, error) {
	var game_str string
	var game model.Session
	row := dal.DB.QueryRow(selectGameSession, user_key)
	err := row.Scan(&game_str)
	if err != nil  {
		return nil, err
	}

	err = json.Unmarshal([]byte(game_str), &game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}


func getCurrentGameInfo(
	dal *SQLiteDAL,
) (*model.Game, error) {
	var game_str string
	var game model.Game
	row := dal.DB.QueryRow(currentGameQuery)
	err := row.Scan(&game_str)
	if err != nil  {
		return nil, err
	}

	err = json.Unmarshal([]byte(game_str), &game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func setCurrentMoviePageFromDB(
	new_movie_page int,
	dal *SQLiteDAL,
) (error) {
	_, err := dal.DB.Exec(setCurrentMoviePageFromDBQuery, new_movie_page)
	if err != nil {
		return err
	}

	return nil
}

func getCurrentMoviePageFromDB(
	dal *SQLiteDAL,
) (*int, error) {
	var current_movie_page int
	row := dal.DB.QueryRow(currentMoviePageFromDBQuery)
	err := row.Scan(&current_movie_page)
	if err != nil  {
		return nil, err
	}

	return &current_movie_page, nil
}


func setupMuralSchema(
	db *sql.DB,
) error {
	_, err := db.Exec(createGameSessionsTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createStatsTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createRedListTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createAnswersTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createGameSessionsTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createCurrentGameTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createMetaTable)
	if err != nil {
		return err
	}

	return nil
}

func getRandomAnswers(dal *SQLiteDAL) ([]model.Answer, error) {
	// first lets get back the answer

	var answer_data string
	var correct_answer model.Answer
	var answers []model.Answer

	row := dal.DB.QueryRow(getRandomCorrectAnswerQuery)
	err := row.Scan(&answer_data)
	if err != nil  {
		return nil, err
	}

	err = json.Unmarshal([]byte(answer_data), &correct_answer)
	if err != nil {
		return nil, err
	}

	correct_answer.IsCorrect = true
	answers = append(answers, correct_answer)

	// now lets get back the rest
	rows, err := dal.DB.Query(getOtherRandomAnswersQuery, correct_answer.ID)
	if err != nil  {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var answer_data string
		var answer model.Answer
		err := rows.Scan(&answer_data)
		if err != nil  {
			return nil, err
		}

		err = json.Unmarshal([]byte(answer_data), &answer)
		if err != nil {
			return nil, err
		}

		answers = append(answers, answer)
	}

	return answers, nil
}

func newSessionForUser(user_key string, dal *SQLiteDAL) (*model.Session, error) {
	session := model.NewSession(
		user_key,
	)

	return &session, nil
}

func setupMetadata(dal *SQLiteDAL) error {
	_, err := dal.DB.Exec(setupMetada)
	if err != nil {
		return err
	}

	return err
}