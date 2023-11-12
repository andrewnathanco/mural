package sql

// game queries
const (
	createGameTable = `
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

// session info
const (
	createSessionTable = `
		create table if not exists sessions (
			session_key integer primary key,
			user_key text unique,
			selected_option_key integer,
			session_status string
		);
	`

	upsertSession = `
		insert into sessions (user_key, selected_option_key, session_status)
		values (:user_key, :selected_option_key, :session_status)
		on conflict (user_key) do update set 
			session_status = excluded.session_status,
			selected_option_key = excluded.selected_option_key
		;
	`

	getSessionByUser = `
		select * from sessions
		where user_key = ?
	;`

	getNumberOfSessionsPlayed = `
		select count(*) from sessions
		where session_status = ?
	; `
)

// tiles
const (
	createTilesTables = `
		create table if not exists tiles (
			tile_key integer primary key,
			row_number integer,
			col_number integer,
			penalty integer
		);
		
		create table if not exists session_tile (
			tile_key integer,
			session_key integer,
			tile_status text
		);
	`

	insertTilesQuery = `
		insert into tiles (tile_key, row_number, col_number, penalty)
		values (:tile_key, :row_number, :col_number, :penalty);
	`
)
