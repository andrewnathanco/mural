-- drop table mural_meta;
create table if not exists mural_meta (
    system_key integer primary key,
    last_tmdb_movie_page integer not null
);
-- drop table games;
create table if not exists games (
    game_key integer primary key,
    option_order integer,
    theme text,
    played_on timestamp,
    game_status text
);
-- drop table "sessions";
create table if not exists sessions (
    session_key integer primary key,
    user_key text unique,
    option_key integer,
    session_status string
);
-- drop table session_tiles;
create table if not exists session_tiles (
    tile_key integer,
    session_key integer,
    tile_status text,
    primary key (tile_key, session_key)
);
-- drop table tiles;
create table if not exists tiles (
    tile_key integer primary key,
    row_number integer,
    col_number integer,
    penalty integer,
    constraint unique_row_col unique (row_number, col_number)
);
-- drop table movies;
create table if not exists movies (
    movie_key integer primary key,
    id integer unique,
    title text,
    original_title text,
    release_date text,
    -- you can use text for date in sqlite
    overview text,
    vote_average real,
    vote_count integer,
    popularity real,
    adult integer,
    -- using integer to represent boolean values (0 for false, 1 for true)
    video integer,
    -- using integer to represent boolean values (0 for false, 1 for true)
    backdrop_path text,
    poster_path text
);
-- drop table users;
create table if not exists users (user_key text unique, game_type text);
-- drop table "options";
create table if not exists options (
    option_key integer primary key,
    movie_key integer,
    game_key integer,
    option_status text
);
-- drop table game_stats
create table if not exists game_stats (
    user_key text,
    game_key integer,
    game_type text,
    session_status text,
    score integer,
    constraint unique_user_game unique (user_key, game_key)
);