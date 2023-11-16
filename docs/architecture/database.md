# Database

This is the database design for Mural. It started as essentially a data dump but has progressed into a normalized schema. It will likely keep evolving, but as it stands, this should be sufficient to include expanding to other game modes, themes or option types.

![database](/docs/assets/mural-db.png)

```sql
CREATE TABLE "games" (
  "game_key" integer PRIMARY KEY,
  "theme" text,
  "played_on" timestamp,
  "game_status" text
);

CREATE TABLE "sessions" (
  "session_key" integer PRIMARY KEY,
  "user_key" text UNIQUE,
  "selected_option_key" integer,
  "session_status" string
);

CREATE TABLE "session_tile" (
  "tile_key" integer,
  "session_key" integer,
  "tile_status" text
);

CREATE TABLE "tiles" (
  "tile_key" integer PRIMARY KEY,
  "row_number" integer,
  "col_number" integer,
  "penalty" integer
);

CREATE TABLE "movies" (
  "movie_key" SERIAL PRIMARY KEY,
  "id" int,
  "title" text,
  "original_title" text,
  "release_date" date,
  "overview" text,
  "vote_average" real,
  "vote_count" int,
  "popularity" real,
  "adult" boolean,
  "video" boolean,
  "backdrop_path" text,
  "poster_path" text
);

CREATE TABLE "options" (
  "option_key" SERIAL PRIMARY KEY,
  "reference_key" int,
  "game_key" int,
  "option_status" text
);

CREATE TABLE "users" (
  "user_key" integer PRIMARY KEY,
  "game_type" text,
  "last_played" text
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("session_key") REFERENCES "session_tile" ("session_key");

ALTER TABLE "tiles" ADD FOREIGN KEY ("tile_key") REFERENCES "session_tile" ("tile_key");

ALTER TABLE "movies" ADD FOREIGN KEY ("movie_key") REFERENCES "options" ("reference_key");

ALTER TABLE "games" ADD FOREIGN KEY ("game_key") REFERENCES "options" ("game_key");

ALTER TABLE "users" ADD FOREIGN KEY ("user_key") REFERENCES "sessions" ("user_key");

ALTER TABLE "options" ADD FOREIGN KEY ("option_key") REFERENCES "sessions" ("selected_option_key");
```
