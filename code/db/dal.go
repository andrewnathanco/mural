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

	// tiles
	PopulateTiles(int) error
	SaveTileStatusForUser(SessionTile) error
	GetTile(int, int) (Tile, error)
	GetSessionTileForUser(int, int, string) (SessionTile, error)

	// movies
	SaveMovies([]Movie) error

	// optoin
	SetNewCorrectOption(config.MuralConfig) (Option, error)
	SetNewEasyModeOptions(config.MuralConfig) ([]Option, error)
}
