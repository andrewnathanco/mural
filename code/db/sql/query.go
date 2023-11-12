package sql

// game queries
const (
	createGameQuery = `
	create table if not exists games (
		game_key integer primary key,
		theme text,
		played_on timestamp,
		game_status text
	);`

	upsertGameQuery = `
		insert into games (game_key, theme, played_on, game_status)
		values (:game_key, :theme, :played_on, :game_status)
		on conflict (game_key) do update set 
			game_status = excluded.game_status
		;
	`

	getGameByStatus = `
		select * from games
		where game_status = ?
	`
)