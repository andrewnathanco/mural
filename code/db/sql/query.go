package sql

const createGameSessionsTable string = `
	create table if not exists game_sessions (
	user_key string not null primary key,
	session string not null
);`

const createStatsTable string = `
	create table if not exists user_stats (
	user_key string not null primary key,
	current_streak int not null,
	longest_streak int not null,
	best_score int not null,
	games_played int not null,
	last_game int not null
);`

const createRedListTable string = `
	create table if not exists red_list (
	answer_key string not null primary key,
	used_date date not null,
	used_game int not null
);`

const createAnswersTable string = `
	create table if not exists answers (
	answer_key string not null primary key,
	answer_data string not null
);`

const createCurrentGameTable string = `
	create table if not exists games (
	game_key int not null primary key,
	is_current bool not null,
	game_data string not null
);`


const createMetaTable string = `
	create table if not exists mural_meta (
	current_game int not null primary key,
	current_movie_page int not null
);`

const createUserDataTable string = `
	create table if not exists user_data (
	user_key string not null primary key,
	hard_mode_enabled bool not null
);`


const upsertGameSession string = `
    insert or replace into game_sessions (user_key, session)
    values (?, ?)
`

const selectGameSession string = `
    select session from game_sessions
    where user_key = ?
`

const resetGameSessions string = `
    delete from game_sessions
`

const insertAnswers string = `
    insert or ignore into answers (answer_key, answer_data)
    values (?, ?)
`

const redlistAnswerQuery string = `
    insert or ignore into red_list (answer_key, used_date, used_game)
    values (?, ?, ?)
`

const currentGameQuery string = `
    select game_data from games
    where is_current = 1
`

const currentMoviePageFromDBQuery string = `
    select current_movie_page from mural_meta
`

const setCurrentMoviePageFromDBQuery string = `
	update mural_meta set current_movie_page = ?
`

const getRandomCorrectAnswerQuery string = `
SELECT answer_data
FROM answers
WHERE answer_key NOT IN (
    SELECT answer_key
    FROM red_list
)
AND (CAST(SUBSTR(json_extract(answer_data, '$.ReleaseDate'), 1, 4) AS INT) / 10) * 10 like ?
ORDER BY random()
LIMIT 1;
`

const getOtherRandomAnswersQuery string = `
select answer_data
from answers
where answer_key != ?
AND (CAST(SUBSTR(json_extract(answer_data, '$.ReleaseDate'), 1, 4) AS INT) / 10) * 10 like ?
order by random()
limit 3;
`

const closeCurrentGame string = `
    update games set is_current = 0
    where is_current = 1
`

const setNewGame string = `
	insert or replace into games (game_key, is_current, game_data)
	values (?, 1, ?)
`

const setupMetada string = `
	insert or ignore into mural_meta (current_game, current_movie_page)
	values (0, 0)
`

// create table if not exists user_stats (
// 	user_key string not null primary key,
// 	current_streak int not null,
// 	best_streak int not null,
// 	best_score int not null,
// 	last_day_played string not null
const setStatsForUserQuery string = `
	insert or replace into user_stats 
		(user_key, current_streak, longest_streak, best_score, games_played, last_game)
	values (?, ?, ?, ?, ?, ?)
`

const getStatsForUserQuery string = `
	select user_key, current_streak, longest_streak, best_score, games_played, last_game
	from user_stats
	where user_key = ?
`

// only have hard mode for now
const getUserDataForUserQuery string = `
	select hard_mode_enabled
	from user_data
	where user_key = ?
`

const insertUserData string = `
	insert or replace into user_data 
		(user_key, hard_mode_enabled)
	values (?, ?)
`


const getNumberOfSessionsQuery string = `
	select count(user_key)
	from game_sessions where session like '%SESSION_OVER%';
`


const getAnswersFromQuery string = `
select answer_data
from answers
where json_extract(answer_data, '$.Name') like ? || '%'
collate nocase
limit 20
`