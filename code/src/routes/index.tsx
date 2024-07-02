import {
  Component,
  Show,
  Suspense,
  createEffect,
  createSignal,
} from "solid-js";
import InfoButton from "../components/buttons/info-button";
import GameArea from "../components/game/presentation/game-area";
import {
  get_current_number_played,
  get_game_key,
  get_todays_game,
} from "../components/game/service";
import InfoDialog from "../components/dialog/info/info-dialog";
import HintDialog from "../components/dialog/hint/hint-dialog";
import { InfoDialogProvider } from "../components/dialog/info/context";
import { HintDialogProvider } from "../components/dialog/hint/context";
import { GameProvider, useGame } from "../components/game/context/game";
import { UserProvider } from "../components/game/context/game-difficulty";
import { Meta } from "@solidjs/meta";


function IndexBody() {
  const [game, set_game] = useGame();
  const [number_played, set_number_played] = createSignal<number | undefined>(
    undefined
  );

  const [version, _] = createSignal(
    import.meta.env.VITE_VERSION ?? "v0.1.0"
  );

  createEffect(() => {
    if (
      game.game_key != get_game_key() ||
      game.correct_option.id != get_todays_game().correct_option.id
    ) {
      localStorage.removeItem("mural_game");
      set_game(get_todays_game());
    }
  });

  createEffect(() => {
    get_current_number_played().then((x) => {
      set_number_played(x);
    });
  });

  return (
    <>
      <Meta
        name="viewport"
        content="width=device-width, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=0"
      />
      <InfoDialogProvider>
          <HintDialogProvider>
            <div class="flex flex-col items-center justify-center">
              <div class="flex flex-col items-center space-y-4 p-4">
                <div class="flex flex-col space-y-4 w-full">
                  <div class="rounded-lg p-4 border-2 border-river-bed-700 flex space-x-1 justify-center">
                    <div>Play other daily games</div>
                    <a
                      class="text-contessa-600 underline"
                      href="https://ancgames.com"
                    >
                      here.
                    </a>
                  </div>
                  <div class="text-5xl flex space-x-2 items-center">
                    <div>Mural #{game.game_key}</div>
                    <div
                      id="game-version"
                      class="font-semibold w-min h-min text-gray-600 text-xs border-2 px-1 border-river-bed-700 rounded-lg"
                    >
                      {version()}
                    </div>
                  </div>
                  <div class="flex justify-between">
                    <div class="flex flex-col space-y-1 items-start">
                      <div id="game-theme" class="text-contessa-500 text-4xl">
                        {game.theme}
                      </div>
                      <div class="text-md">Today's Theme</div>
                    </div>
                    {number_played() ? (
                      <div class="flex flex-col space-y-1 items-start">
                        <div
                          id="games-played"
                          class="text-contessa-500 text-4xl"
                        >
                          {number_played()}
                        </div>
                        <div class="text-md">Have Played</div>
                      </div>
                    ) : (
                      <></>
                    )}
                  </div>

                  <div class="flex flex-col space-y-1">
                    <div class="w-full">
                      <InfoButton />
                    </div>
                  </div>

                  <div class="h-0.5 w-full rounded-full bg-river-bed-600"></div>

                  <div class="text-3xl flex space-x-2 items-center flex-col">
                    <GameArea />
                  </div>
                </div>
              </div>
              <InfoDialog />
              <HintDialog />
            </div>
          </HintDialogProvider>
      </InfoDialogProvider>
    </>
  );
};

export default function App() {
  return (
    <GameProvider>
      <UserProvider>
        <IndexBody />
      </UserProvider>
    </GameProvider>
  );
}
