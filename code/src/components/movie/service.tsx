import { AvailableThemes, GameTheme } from "../game/model/game";
import { Movie } from "./model";
import movies from "./movies.json";
import seedrandom from "seedrandom";

function shuffle<T>(array: T[]) {
  for (let i = array.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [array[i], array[j]] = [array[j], array[i]];
  }

  return array;
}

export function get_correct_option_by_theme_and_key(
  theme: GameTheme,
  game_key: number
): Movie {
  if (theme != GameTheme.random) {
    const today_movies = movies[theme];
    let sorted_movies = today_movies.toSorted(
      (a, b) => b.vote_count - a.vote_count
    );

    var rng = seedrandom(game_key.toString());
    const rand_index = Math.floor(rng() * movies[theme].length);
    const movie = sorted_movies[rand_index];
    return movie;
  } else {
    const avail_themes = Object.values(AvailableThemes);
    var rng = seedrandom(game_key.toString());
    const rand_index = Math.floor(rng() * avail_themes.length);
    const rand_theme = avail_themes[rand_index] as AvailableThemes;
    const today_movies = movies[rand_theme];
    let sorted_movies = today_movies.sort(
      (a, b) => b.vote_count - a.vote_count
    );

    rng = seedrandom(game_key.toString());
    const rand_movie_index = Math.floor(rng() * movies[theme].length);

    const movie = sorted_movies[rand_movie_index];
    return movie;
  }
}

export function get_easy_mode_options_by_theme(
  theme: GameTheme,
  correct_option: Movie
): Movie[] {
  let movie_options: Movie[] = [];
  if (theme != GameTheme.random) {
    movie_options = movies[theme];
  } else {
    const avail_themes = Object.values(AvailableThemes);
    const rand_index = Math.floor(Math.random() * avail_themes.length);
    const rand_theme = avail_themes[rand_index] as AvailableThemes;
    movie_options = movies[rand_theme];
  }

  const rand_movie_options = shuffle(movie_options);
  return shuffle([...rand_movie_options.slice(0, 3), correct_option]);
}

export function query_option(query: string) {
  const all_movies: Movie[] = [];

  for (const decade in movies) {
    if (movies.hasOwnProperty(decade)) {
      all_movies.push(...movies[decade as AvailableThemes]);
    }
  }

  const filtered_movies = all_movies.filter(
    (movie) =>
      movie.title.toLowerCase().includes(query.toLowerCase()) ||
      movie.title.toLowerCase() == query.toLowerCase()
  );

  filtered_movies.sort((a, b) => {
    return b.vote_count - a.vote_count || a.title.length - b.title.length;
  });

  return filtered_movies;
}

export function get_movie_from_id(id: number) {
  const all_movies: Movie[] = [];

  for (const decade in movies) {
    if (movies.hasOwnProperty(decade)) {
      all_movies.push(...movies[decade as AvailableThemes]);
    }
  }

  return all_movies.find((movie) => movie.id == id) ?? all_movies[0];
}
