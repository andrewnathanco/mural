import { makePersisted } from "@solid-primitives/storage";
import { createContext, useContext } from "solid-js";
import { SetStoreFunction, createStore } from "solid-js/store";
import { Game, GameStatus } from "./model";
import { GameDifficulty } from "./presentation/board/difficulty-selector";

const GameContext = createContext<[Game, SetStoreFunction<Game>]>([
  {} as Game,
  () => {},
]);

export function GameProvider(props: any) {
  let value = makePersisted(
    createStore<Game>({
      flipped: [],
      score: 100,
      difficulty: GameDifficulty.easy,
      status: GameStatus.started,
    }),
    {
      name: "game",
    }
  );

  return (
    <GameContext.Provider value={value}>{props.children}</GameContext.Provider>
  );
}

export function useGame() {
  return useContext(GameContext);
}
