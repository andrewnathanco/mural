package db

import "mural/config"

type IDAL interface {
	PingDatabase() error

	// meta functions
	GetMeta() (MuralMeta, error)
	UpsertMeta(MuralMeta) error
	GetMuralForUser(string, config.MuralConfig) (Mural, error)

	// game functions
	UpsertGame(Game) error
	GetCurrentGame(config.MuralConfig) (Game, error)

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

	// optoin
	SetNewCorrectOption(config.MuralConfig) (Option, error)
	SetNewEasyModeOptions(config.MuralConfig) ([]Option, error)
	// TODO: Add unit test
	GetCorrectOption() (Option, error)
	// TODO: Add unit test
	GetEasyModeOptions() ([]Option, error)

	// user methods
	UpsertUser(User) error
	GetUserByUserKey(user_key string) (User, error)
}
