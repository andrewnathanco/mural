import { makePersisted } from "@solid-primitives/storage";
import { createContext, useContext } from "solid-js";
import { SetStoreFunction, createStore } from "solid-js/store";
import { Game } from "../model/game";
import { get_todays_game } from "../service";
import { GameDifficulty } from "../presentation/board/difficulty-selector";
import { User } from "../model/user";

const UserContext = createContext<[User, SetStoreFunction<User>]>([
  {} as User,
  () => {},
]);

export function UserProvider(props: any) {
  let value = makePersisted(
    createStore<User>({ difficulty: GameDifficulty.hard }),
    {
      name: "mural_user",
    }
  );

  return (
    <UserContext.Provider value={value}>{props.children}</UserContext.Provider>
  );
}

export function useUser() {
  return useContext(UserContext);
}
