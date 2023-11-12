package db

var (
	DAL IDAL
)

type IDAL interface {
	PingDatabase() (error)
	UpsertGame(Game) (error)
	GetCurrentGame() (*Game, error)
} 