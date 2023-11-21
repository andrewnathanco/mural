package db

import "mural/config"

type IDAL interface {
	PingDatabase() error

	// meta functions
	GetMeta() (MuralMeta, error)
	UpsertMeta(MuralMeta) error
	GetMuralForUser(string, config.MuralConfig, bool) (Mural, error)

	// game functions
	UpsertGame(Game) error
	GetOrCreateNewGame(config.MuralConfig) (Game, error)
	GetCurrentGame() (Game, error)

	// get session
	UpsertSession(Session) error
	GetSessionForUser(string) (Session, error)
	GetNumberOfSessionsPlayed() (int, error)
	DeleteSessions() error
	// TODO: Add unit test
	GetScoreForUser(config.MuralConfig, string) (int, error)

	// tiles
	PopulateTiles(int) error
	SelectTileForUser(SessionTile) error
	SaveTileStatusForUser(SessionTile) error
	GetTile(int, int) (Tile, error)
	GetSessionTileForUser(int, int, string) (SessionTile, error)
	// TODO: Add unit test
	GetBoardForUser(config.MuralConfig, string) ([][]SessionTile, error)

	// movies
	SaveMovies([]Movie) error
	// TODO: test this
	GetMovieByMovieKey(int) (Movie, error)
	GetMovieByMovieID(int) (Movie, error)

	// option
	SetNewCorrectOption(config.MuralConfig, *Movie) (Option, error)
	SetNewEasyModeOptions(config.MuralConfig) ([]Option, error)
	// TODO: Add unit test
	GetCorrectOption() (Option, error)
	// TODO: Add unit test
	GetEasyModeOptions() ([]Option, error)
	// TODO: test this
	GetOptionByMovie(int) (Option, error)
	// TODO: test this
	GetOptionByKey(int64) (Option, error)
	// TODO test this
	GetOptionsByQuery(string) ([]Option, error)

	// user methods
	UpsertUser(User) error
	GetUserByUserKey(string) (User, error)
	CheckForUser(string) (bool, error)

	// stats
	UpsertGameStat(GameStat) error
	GetStreaksForUser(string) (int, int, error)
	GetWeeklyStatsForUser(string) (map[string]map[string]DayStat, error)
}
