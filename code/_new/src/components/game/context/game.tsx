import { makePersisted } from "@solid-primitives/storage";
import { createContext, useContext } from "solid-js";
import { SetStoreFunction, createStore } from "solid-js/store";
import { Game } from "../model";
import { get_todays_game } from "../service";

const GameContext = createContext<[Game, SetStoreFunction<Game>]>([
  {} as Game,
  () => {},
]);

export function GameProvider(props: any) {
  let value = makePersisted(createStore<Game>(get_todays_game()), {
    name: "mural_game",
  });

  return (
    <GameContext.Provider value={value}>{props.children}</GameContext.Provider>
  );
}

export function useGame() {
  return useContext(GameContext);
}
const FreePlayGameContext = createContext<[Game, SetStoreFunction<Game>]>([
  {} as Game,
  () => {},
]);
