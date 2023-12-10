import { Movie } from "../movie/model";
import { GameDifficulty } from "./presentation/board/difficulty-selector";

export interface GameTile {
  row: number;
  col: number;
  key: number;
  penalty: number;
  flipped: boolean;
}

export enum GameStatus {
  won,
  lost,
  init,
  started,
}

export enum BoardState {
  current,
  flipped,
  game,
}

export enum GameTheme {
  _2020 = "2020",
  _2010 = "2010",
  _2000 = "2000",
  _1990 = "1990",
  _1980 = "1980",
  _1970 = "1970",
  random = "Random",
}

export enum AvailableThemes {
  _2020 = "2020",
  _2010 = "2010",
  _2000 = "2000",
  _1990 = "1990",
  _1980 = "1980",
  _1970 = "1970",
}

export interface Hints {
  year: boolean;
  genres: boolean;
  description: boolean;
}

export interface Game {
  game_key: number;
  flipped: number[];
  score: number | "❎";
  selected_tile?: number;
  difficulty: GameDifficulty;
  status: GameStatus;
  correct_option: Movie;
  easy_mode_options: Movie[];
  selected_option?: Movie;
  theme: GameTheme;
  board_state: BoardState;
  hints: Hints;
}
