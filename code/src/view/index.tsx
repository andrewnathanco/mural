import { createEffect } from "solid-js";
import InfoButton from "../components/buttons/info-button";
import GameArea from "../components/game/presentation/game-area";
import {
  get_game_key,
  get_random_game,
  get_todays_game,
} from "../components/game/service";
import InfoDialog from "../components/dialog/info/info-dialog";
import HintDialog from "../components/dialog/hint/hint-dialog";
import { ShareDialog } from "../components/dialog/share/share-dialog";
import { InfoDialogProvider } from "../components/dialog/info/context";
import { ShareDialogProvider } from "../components/dialog/share/context";
import { HintDialogProvider } from "../components/dialog/hint/context";
import { GameProvider, useGame } from "../components/game/context/game";
import { UserProvider } from "../components/game/context/game-difficulty";

export function IndexBody() {
  const [game, set_game] = useGame();

  createEffect(() => {
    if (game.game_key != get_game_key()) {
      localStorage.removeItem("mural_game");
      set_game(get_todays_game());
    }
  });

  return (
    <InfoDialogProvider>
      <ShareDialogProvider>
        <HintDialogProvider>
          <div class="flex flex-col items-center justify-center">
            <div class="flex flex-col items-center space-y-4">
              <div class="flex flex-col space-y-4 w-full">
                <div class="text-5xl flex space-x-2 items-center">
                  <div>Mural #{game.game_key}</div>
                  <div
                    id="game-version"
                    class="font-semibold w-min h-min text-gray-600 text-xs border-2 px-1 border-river-bed-700 rounded-lg"
                  >
                    {import.meta.env.VITE_VERSION ?? "v0.1.1"}
                  </div>
                </div>
                <div class="flex justify-between">
                  <div class="flex flex-col space-y-1 items-start">
                    <div id="game-theme" class="text-contessa-500 text-4xl">
                      {game.theme}
                    </div>
                    <div class="text-md">Today's Theme</div>
                  </div>
                </div>

                <div class="flex flex-col space-y-1">
                  <div class="w-full">
                    <InfoButton />
                  </div>
                  <button
                    onclick={() => {
                      set_game(get_random_game());
                    }}
                    id="info-button"
                    class=" w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700 "
                  >
                    Reset Game
                  </button>
                </div>

                <div class="h-0.5 w-full rounded-full bg-river-bed-600"></div>

                <div class="text-3xl flex space-x-2 items-center flex-col">
                  <GameArea />
                </div>
              </div>
            </div>
            <InfoDialog />
            <ShareDialog />
            <HintDialog />
          </div>
        </HintDialogProvider>
      </ShareDialogProvider>
    </InfoDialogProvider>
  );
}

export default function Index() {
  return (
    <GameProvider>
      <UserProvider>
        <IndexBody />
      </UserProvider>
    </GameProvider>
  );
}
