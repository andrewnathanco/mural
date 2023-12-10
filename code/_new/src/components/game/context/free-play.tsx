import { makePersisted } from "@solid-primitives/storage";
import { createContext, useContext } from "solid-js";
import { SetStoreFunction, createStore } from "solid-js/store";
import { Game } from "../model";
import { get_todays_game } from "../service";

const FreePlayGameContext = createContext<[Game, SetStoreFunction<Game>]>([
  {} as Game,
  () => {},
]);

export function FreePlayGameProvider(props: any) {
  let value = makePersisted(createStore<Game>(get_todays_game()), {
    name: "mural_free-play-game",
  });

  return (
    <FreePlayGameContext.Provider value={value}>
      {props.children}
    </FreePlayGameContext.Provider>
  );
}

export function useFreePlay() {
  return useContext(FreePlayGameContext);
}
