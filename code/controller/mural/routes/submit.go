package routes

import (
	"mural/api"
	"mural/app"
	"mural/db"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AnswerType string

const GIVE_UP = "GIVE_UP"

func Submit(c echo.Context) error {
	option := c.QueryParam("type")

	user_key := middleware.GetUserKeyFromContext(c)
	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)
	mural_ses, err := mural_service.DAL.GetMuralForUser(
		user_key,
		mural_service.Config,
	)

	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
	}

	if option == GIVE_UP {
		mural_ses.Session.SessionStatus = db.SESSION_LOST
		mural_ses.Session.CurrentScore = nil
	} else {
		if mural_ses.Session.OptionKey == mural_ses.Game.CorrectOption.OptionKey {
			mural_ses.Session.SessionStatus = db.SESSION_WON
		} else {
			mural_ses.Session.SessionStatus = db.SESSION_LOST
			mural_ses.Session.CurrentScore = nil
		}
	}

	stat := db.GameStat{
		Game:          mural_ses.Game,
		GameType:      mural_ses.User.GameType,
		UserKey:       user_key,
		SessionStatus: mural_ses.Session.SessionStatus,
		Score:         mural_ses.Session.CurrentScore,
	}

	err = mural_service.DAL.UpsertGameStat(stat)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not upsert stat")
	}

	err = mural_service.DAL.UpsertSession(mural_ses.Session)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not upsert session")
	}

	// do analytics stuff
	mural_service.AnalyticsContoller.RegisterEvent(api.EVENT_SUBMIT, c.Request())
	mural_ses, err = mural_service.DAL.GetMuralForUser(
		user_key,
		mural_service.Config,
	)

	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
	}

	return c.Render(http.StatusOK, "game-board.html", mural_ses)
}
