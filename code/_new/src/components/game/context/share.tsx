import { makePersisted } from "@solid-primitives/storage";
import { createContext, useContext } from "solid-js";
import { SetStoreFunction, createStore } from "solid-js/store";
import { Game } from "../model";
import { get_todays_game } from "../service";

const ShareContext = createContext<[Game, SetStoreFunction<Game>]>([
  {} as Game,
  () => {},
]);

export function ShareProvider(props: any) {
  let value = makePersisted(createStore<Game>(get_todays_game()), {
    name: "mural_share-game",
  });

  return (
    <ShareContext.Provider value={value}>
      {props.children}
    </ShareContext.Provider>
  );
}

export function useShare() {
  return useContext(ShareContext);
}
