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

export interface Game {
  flipped: number[];
  score: number;
  selected?: number;
  difficulty: GameDifficulty;
  status: GameStatus;
}
