package routes

import (
	"mural/app"
	"mural/db"
	"mural/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SelectOption(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	movie_key := c.QueryParam("option")
	movie_key_int, err := strconv.ParseInt(movie_key, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "need to define an option")
	}

	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)
	option, err := mural_service.DAL.GetOptionByMovie(int(movie_key_int))
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get the option")
	}

	sess, err := mural_service.DAL.GetSessionForUser(user_key)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get session")
	}

	sess.Option = option
	err = mural_service.DAL.UpsertSession(sess)
	if err != nil {
		return c.String(http.StatusBadRequest, "couldn't save tile")
	}

	mural_ses, err := mural_service.DAL.GetMuralForUser(
		user_key,
		mural_service.Config,
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
	}

	if mural_ses.User.GameType == db.EASY_MODE {
		return c.Render(http.StatusOK, "answers.html", mural_ses)
	} else {
		return c.Render(http.StatusOK, "answer-input.html", mural_ses)
	}

}
