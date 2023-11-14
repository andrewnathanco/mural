package sql

// meta
const (
	createMetaTable = `
		create table if not exists mural_meta (
			system_key int primary key,
			last_tmdb_movie_page integer not null
		)
	`

	upsertMeta = `
		insert into mural_meta (system_key, last_tmdb_movie_page)
		values (:system_key, :last_tmdb_movie_page)
		on conflict (system_key) do update set 
			last_tmdb_movie_page = excluded.last_tmdb_movie_page
		;
	`

	getMeta = `
		select * from mural_meta
	`
)

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

	getLastGame = `
		select * from games order by game_key desc
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
	deleteSessions = `
		delete from sessions
; `
)

// tiles
const (
	createTilesTables = `
		create table if not exists tiles (
			tile_key integer primary key,
			row_number integer,
			col_number integer,
			penalty integer,

			constraint unique_row_col unique (row_number, col_number)
		);
		
		create table if not exists session_tiles (
			tile_key integer,
			session_key integer,
			tile_status text,

			primary key (tile_key, session_key)
		);
	`

	insertTilesQuery = `
		insert into tiles (row_number, col_number, penalty)
		values (:row_number, :col_number, :penalty)
		on conflict (tile_key) do update set 
			penalty = excluded.penalty
		on conflict (row_number, col_number) do update set 
			penalty = excluded.penalty
		;
	`

	upsertSessionTiles = `
		insert into session_tiles (tile_key, session_key, tile_status)
		values (:tile_key, :session_key, :tile_status)
		on conflict (tile_key, session_key) do update set 
			tile_status = excluded.tile_status
		;
	`

	getTile = `
	select 
		t.*
	from 
		tiles t
	where 
		t.row_number = ? and t.col_number = ?
`

	getSessionTileForUser = `
		select 
			s.*,
			t.*
		from 
			session_tiles s
		inner join 
			tiles t on t.tile_key = s.tile_key
		inner join 
			sessions sess on sess.session_key = s.session_key
		where 
			t.row_number = ? and t.col_number = ?
		and sess.user_key = ?
	`
)

// movies
const (
	createMovieTable = `
		create table if not exists movies (
			movie_key integer primary key,
			id integer unique,
			title text,
			original_title text,
			release_date text, -- you can use text for date in sqlite
			overview text,
			vote_average real,
			vote_count integer,
			popularity real,
			adult integer, -- using integer to represent boolean values (0 for false, 1 for true)
			video integer, -- using integer to represent boolean values (0 for false, 1 for true)
			backdrop_path text,
			poster_path text
		);
	`

	upsertMovie = `
		insert into movies (id, title, release_date, original_title, overview, vote_average, vote_count, popularity, adult, video, poster_path, backdrop_path)
		values (:id, :title, :release_date, :original_title,  :overview, :vote_average, :vote_count, :popularity, :adult, :video, :poster_path, :backdrop_path)
		on conflict(id) do nothing
	`

	getMovieByKey = `
		select * from movies where movie_key = ?
	`

	getRandomMovie = `
		select 
			* 
		from movies 
		left join "options" o on o.movie_key = movies.movie_key 
		where 
			((vote_count * 0.8 + popularity * 0.2) / 21) >= ?
		and option_key is null
		order by random()
	`
)

// optoins
const (
	createOptionTable = `
		create table if not exists options (
			option_key integer primary key,
			movie_key integer,
			game_key integer,
			option_status text
		);
	`

	upsertOption = `
		insert into options (movie_key, game_key, option_status)
		values (:movie_key, :game_key, :option_status)
		on conflict (option_key) do update set 
			option_status = excluded.option_status
	;
	`

	getCurrentCorrectOption = `
		select * 
			from option 
		where 
			option_status = ?
	`
)

// users
const (
	createUsersTable = `
		create table if not exist users (
			user_key integer primary key,
			name text,
			game_type text,
			total_score text,
			last_played text
		);
	`
)
