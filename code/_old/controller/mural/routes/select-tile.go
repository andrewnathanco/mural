package routes

import (
	"mural/app"
	"mural/db"
	"mural/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SelectTile(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	row_num := c.QueryParam("row_num")
	row_num_int, err := strconv.ParseInt(row_num, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "need to define in the row direction")
	}

	col_num := c.QueryParam("col_num")
	col_num_int, err := strconv.ParseInt(col_num, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "need to define in the col direction")
	}

	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)

	tile, err := mural_service.DAL.GetTile(int(row_num_int), int(col_num_int))
	if err != nil {
		return c.String(http.StatusBadRequest, "couldn't get tile")
	}

	sess, err := mural_service.DAL.GetSessionForUser(user_key)
	if err != nil {
		return c.String(http.StatusBadRequest, "couldn't get session")
	}

	session_tile := db.SessionTile{
		SessionKey:        sess.SessionKey,
		Tile:              tile,
		SessionTileStatus: db.TILE_SELECTED,
	}

	err = mural_service.DAL.SelectTileForUser(session_tile)
	if err != nil {
		return c.String(http.StatusBadRequest, "couldn't save tile")
	}

	mural_ses, err := mural_service.DAL.GetMuralForUser(
		user_key,
		mural_service.Config,
		false,
	)
	if err != nil {
		return c.String(http.StatusBadRequest, "couldn't get sess")
	}

	return c.Render(http.StatusOK, "game-board.html", mural_ses)
}
