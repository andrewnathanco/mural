package db

var (
	DAL IDAL
)

type IDAL interface {
	PingDatabase() error

	// game functions
	UpsertGame(Game) error
	GetCurrentGame() (*Game, error)

	// get session
	UpsertSession(Session) error
	GetSessionByUser(string) (*Session, error)
	GetNumberOfSessionsPlayed() (int, error)
}
