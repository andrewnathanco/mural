package sql

// meta
const (
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
	upsertGameQuery = `
		insert into games (game_key, option_order, theme, played_on, game_status)
		values (:game_key, :option_order, :theme, :played_on, :game_status)
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
	upsertSession = `
		insert into sessions (user_key, session_status, option_key)
		values (:user_key, :session_status, :option_key)
		on conflict (user_key) do update set 
			session_status = excluded.session_status,
			option_key = excluded.option_key
		;
	`

	getSessionByUser = `
		select * from sessions
		where user_key = ?
	;`

	getNumberOfSessionsPlayed = `
		select count(*) from sessions
		where session_status = ? or session_status = ?
	; `
	deleteSessions = `
		delete from sessions
; `
)

// tiles
const (
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
	updateOtherSelectedTiles = `
		update session_tiles 
		set tile_status = ?
		where tile_status = ?
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

	deleteSessionTiles = `
	delete from session_tiles; 
	`
)

// movies
const (
	upsertMovie = `
		insert into movies (id, title, release_date, original_title, overview, vote_average, vote_count, popularity, adult, video, poster_path, backdrop_path)
		values (:id, :title, :release_date, :original_title,  :overview, :vote_average, :vote_count, :popularity, :adult, :video, :poster_path, :backdrop_path)
		on conflict(id) do nothing
	`

	getRandomMovie = `
	select 
		movies.* 
	from movies 
		left join "options" o on o.movie_key = movies.movie_key 
	where 
		vote_count >= ?
	and substr(release_date, 1, 4) like ?
	and option_key is null
	order by random()
	limit ?
	`

	getMovieBykey = `
	select 
		 *
	from movies 
	where movie_key = ?
	`

	getMovieByID = `
	select 
		 *
	from movies 
	where id = ?
	`

	queryMovies = `
		select *
		from movies
		where title like ? || '%'
		collate nocase
		limit 20
	`
)

// optoins
const (
	upsertOption = `
		insert into options (movie_key, game_key, option_status)
		values (:movie_key, :game_key, :option_status)
	;
	`

	getOptionByStatus = `
		select * 
			from options
		inner join movies on movies.movie_key = options.movie_key
		where 
			option_status = ?
	`

	resetOptionByStatus = `
		update options 
		set option_status = ?
		where option_status = ?
	`
	getOptionByKey = `
		select * from options
		inner join movies on movies.movie_key = options.movie_key
		where option_key = ?
	;`
)

// user
const (
	upsertUser = `
		insert into users (user_key, game_type, display_name)
		values (:user_key, :game_type, :display_name)
		on conflict (user_key) do update set 
			game_type = excluded.game_type,
			display_name = excluded.display_name
	`
	getUserByKey = `
		select * from users
		where user_key = ?;`
)

// game stats
const (
	upsertGameStat = `
		insert into game_stats (user_key, game_key, game_type, session_status, score)
		values (:user_key, :game_key, :game_type, :session_status, :score)
		on conflict (user_key, game_key) do nothing
	`

	getTotalGamesPlayedByUser = `
		select count(*) from game_stats where user_key = ?
	`

	getAllGamesStatsForUser = `
		select * 
		from game_stats 
		inner join games on games.game_key = game_stats.game_key
		where user_key = ?
		order by game_stats.game_key desc
	`
)