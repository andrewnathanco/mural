import { Movie } from "../movie/model";
import {
  get_correct_option_by_theme_and_key,
  get_easy_mode_options_by_theme,
} from "../movie/service";
import {
  AvailableThemes,
  BoardState,
  Game,
  GameStatus,
  GameTheme,
} from "./model/game";

export function get_digits(key: number): { row: number; col: number } {
  if (key < 0 || key > 99 || isNaN(key) || !Number.isInteger(key)) {
    return { row: 0, col: 0 };
  }

  const tens = Math.floor(key / 10);
  const ones = key % 10;

  return { row: tens, col: ones };
}

export function get_penalty_from_key(key: number) {
  const { row, col } = get_digits(key);
  return Math.min(row, col, 10 - row - 1, 10 - col - 1) + 1;
}

export function get_game_key() {
  const now: Date = new Date();
  const specificDate: Date = new Date(2023, 10, 26, 5, 0, 0);
  const duration: number =
    (now.getTime() - specificDate.getTime()) / (1000 * 60 * 60 * 24);

  // return Math.floor(duration);
  return 15;
}

export function get_theme_for_day(): GameTheme {
  const estOffset = -5 * 60 * 60 * 1000; // 5 hours behind UTC for Eastern Standard Time (EST)
  const now: Date = new Date(new Date().getTime() + estOffset);

  const dayOfWeek = now.getUTCDay(); // 0 for Sunday, 1 for Monday, ..., 6 for Saturday

  switch (dayOfWeek) {
    case 1:
      return GameTheme._2020;
    case 2:
      return GameTheme._2010;
    case 3:
      return GameTheme._2000;
    case 4:
      return GameTheme._1990;
    case 5:
      return GameTheme._1980;
    case 6:
      return GameTheme._1970;
    default:
      return GameTheme.random;
  }
}

export function get_countdown_till_next_game(): string {
  // Get the current local time
  const now = new Date();

  // midnight EST in UTC
  const midnight = new Date(
    now.getFullYear(),
    now.getMonth(),
    now.getDate(),
    5,
    0,
    0,
    0
  );

  // Calculate the duration until midnight EST
  const durationUntilMidnight = midnight.getTime() - now.getTime();

  const hours = Math.floor(durationUntilMidnight / (1000 * 60 * 60)) + 24;
  const minutes =
    Math.floor((durationUntilMidnight % (1000 * 60 * 60)) / (1000 * 60)) + 60;
  const seconds = Math.floor((durationUntilMidnight % (1000 * 60)) / 1000) + 60;

  return `${padZero(hours)}:${padZero(minutes)}:${padZero(seconds)}`;
}

function padZero(num: number): string {
  return num < 10 ? `0${num}` : `${num}`;
}

export function get_todays_game(): Game {
  const today_game_key = get_game_key();
  const theme = get_theme_for_day();
  const correct_option = get_correct_option_by_theme_and_key(
    theme,
    today_game_key
  );

  return {
    game_key: today_game_key,
    board_state: BoardState.current,
    flipped: [],
    score: 100,
    status: GameStatus.init,
    theme,
    correct_option,
    selected_tile: undefined,
    selected_option: undefined,
    hints: { year: false, genres: false, description: false },
    easy_mode_options: get_easy_mode_options_by_theme(theme, correct_option),
  };
}

export async function get_current_number_played() {
  const db_url = import.meta.env.VITE_DB_URL;
  const response = await fetch(`${db_url}/GET/mural_games-played`);
  const games_played_res = await response.json();
  const num_played = parseInt(games_played_res["GET"]);
  return Number.isNaN(num_played) ? undefined : num_played;
}

export async function update_number_of_games_played() {
  const db_url = import.meta.env.VITE_DB_URL;
  const num_played = await get_current_number_played();
  if (!num_played) {
    await fetch(`${db_url}/SET/mural_games-played/1`);
  } else {
    await fetch(`${db_url}/SET/mural_games-played/${num_played + 1}`);
  }
}
