package routes

import (
	"database/sql"
	"fmt"
	"log/slog"
	"math/rand"
	"mural/app"
	"mural/config"
	"mural/db"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func ManualOverride(c echo.Context) error {
	override_key := c.QueryParam("override_key")
	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)

	if override_key != mural_service.Config.OverrideKey {
		return c.String(http.StatusForbidden, "incorrect or missing override key")
	}

	// step 1:
	slog.Info("Manually Overriding game")
	movie_id := c.QueryParam("movie_id")
	movie_id_int, err := strconv.ParseInt(movie_id, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "couldn't convert")
	}
	if override_key != mural_service.Config.OverrideKey {
		return c.String(http.StatusBadRequest, "incorrect movie key")
	}

	movie, err := mural_service.DAL.GetMovieByMovieID(int(movie_id_int))
	if err != nil {
		slog.Error(err.Error())
		if err == sql.ErrNoRows {
			return c.String(http.StatusBadRequest, "not a valid movie")
		} else {
			return c.String(http.StatusBadRequest, "invalid movie key")
		}
	}

	// now that we have a movie lets start doing the rest
	err = mural_service.DAL.DeleteSessions()
	if err != nil {
		slog.Error(err.Error())
		return c.String(http.StatusInternalServerError, "couldn't delete sessions")
	}

	// for now lets do this automatically, eventually we can look into doing this in the request
	mural_service.Config.TodayTheme = config.GetTodayThemeDefault()

	last_game, err := mural_service.DAL.GetOrCreateNewGame(mural_service.Config)
	if err != nil {
		slog.Error(err.Error())
		return c.String(http.StatusInternalServerError, "couldn't get or create new game")
	}

	// start generating options
	_, err = mural_service.DAL.SetNewCorrectOption(mural_service.Config, &movie)
	if err != nil {
		slog.Error(err.Error())
		return c.String(http.StatusInternalServerError, "couldn't set new correct option")
	}

	_, err = mural_service.DAL.SetNewEasyModeOptions(mural_service.Config)
	if err != nil {
		slog.Error(err.Error())
		return c.String(http.StatusInternalServerError, "couldn't set new easy mode options")
	}

	// start building the new one
	new_game := db.Game{
		GameKey:     last_game.GameKey,
		OptionOrder: rand.Intn(4),
		PlayedOn:    time.Now(),
		Theme:       mural_service.Config.TodayTheme,
		GameStatus:  db.GAME_CURRENT,
	}

	err = mural_service.DAL.UpsertGame(new_game)
	if err != nil {
		return c.String(http.StatusInternalServerError, "couldn't set game")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Update game: %d, with movie: %s", new_game.GameKey, movie.Title))
}
