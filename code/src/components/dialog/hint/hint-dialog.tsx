import { createEffect, createSignal } from "solid-js";
import { useHintDialog } from "./context";
import { get_countdown_till_next_game } from "../../game/service";
import { useGame } from "../../game/context/game";

export default function HintDialog() {
  const [is_open, { open, close }] = useHintDialog();
  const [game, set_game] = useGame();

  return (
    <div classList={{ hidden: !is_open(), block: is_open() }}>
      <div class="z-10 absolute top-0 left-0 right-0 bottom-0 justify-center items-center bg-black flex opacity-70"></div>
      <div class="z-20 p-4 border-2 border-river-bed-700 absolute top-0 left-0 right-0 md:w-128 md:mx-auto m-4 rounded-lg bg-desert-sand-100 shadow-lg flex flex-col space-y-2 justify-between overflow-auto">
        <div id="dialog-content" class="p-4 flex flex-col space-y-2 w-full">
          <div id="info-dialog">
            <div class="flex flex-col space-y-2">
              <div class="flex justify-between items-center">
                <div class="text-3xl">Hints</div>
                <button
                  onClick={() => {
                    close();
                  }}
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    fill="currentColor"
                    class="w-6 h-6"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M5.47 5.47a.75.75 0 011.06 0L12 10.94l5.47-5.47a.75.75 0 111.06 1.06L13.06 12l5.47 5.47a.75.75 0 11-1.06 1.06L12 13.06l-5.47 5.47a.75.75 0 01-1.06-1.06L10.94 12 5.47 6.53a.75.75 0 010-1.06z"
                      clip-rule="evenodd"
                    ></path>
                  </svg>
                </button>
              </div>
              <div class="h-0.5 w-full rounded-full bg-river-bed-600"></div>
              <div class="flex flex-col space-y-2">
                <div class="text-2xl">Year</div>
                <div class="text-contessa-500 text-lg">
                  {game.hints?.year ? (
                    new Date(game.correct_option.release_date).getFullYear()
                  ) : (
                    <button
                      onclick={() => {
                        if (typeof game.score == "number") {
                          set_game("score", game.score - 5);
                        }

                        set_game("hints", { ...game?.hints, year: true });
                      }}
                      id="info-button"
                      class=" w-full p-1 text-base text-desert-sand-100 rounded-md flex justify-center items-center border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700 "
                    >
                      Reveal Year (-5)
                    </button>
                  )}
                </div>
                <div>Genres</div>
                <div class="flex space-x-2">
                  {game.hints?.genres ? (
                    game.correct_option.genres.map((genre) => {
                      return (
                        <div class="rounded-full px-2 bg-contessa-500 text-el-salva-100">
                          {genre.name}
                        </div>
                      );
                    })
                  ) : (
                    <button
                      onclick={() => {
                        if (typeof game.score == "number") {
                          set_game("score", game.score - 10);
                        }

                        set_game("hints", { ...game?.hints, genres: true });
                      }}
                      id="info-button"
                      class=" w-full p-1 text-base text-desert-sand-100 rounded-md flex justify-center items-center border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700 "
                    >
                      Reveal Genres (-10)
                    </button>
                  )}
                </div>
                <div>Description</div>
                <div class="text-contessa-500 text-lg flex space-x-2">
                  {game.hints?.description ? (
                    game.correct_option.overview
                  ) : (
                    <button
                      onclick={() => {
                        if (typeof game.score == "number") {
                          set_game("score", game.score - 25);
                        }

                        set_game("hints", {
                          ...game?.hints,
                          description: true,
                        });
                      }}
                      id="info-button"
                      class=" w-full p-1 text-base text-desert-sand-100 rounded-md flex justify-center items-center border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700 "
                    >
                      Reveal Description (-25)
                    </button>
                  )}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
