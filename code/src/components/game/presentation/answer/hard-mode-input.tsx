import { createSignal } from "solid-js";
import { query_option } from "../../../movie/service";
import { Movie } from "../../../movie/model";
import { GameStatus } from "../../model/game";
import CorrectOption from "./input/correct-option";
import WrongOption from "./input/wrong-option";
import { useGame } from "../../context/game";

export default function HardModeInput() {
  const [game, set_game] = useGame();
  const [options, set_options] = createSignal<Movie[]>([]);

  return (
    <div class="flex flex-col space-y-2">
      {game.status == GameStatus.started ? (
        <div class="relative">
          <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
            <svg
              class="w-4 h-4 text-river-bed-700"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 20 20"
            >
              <path
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"
              />
            </svg>
          </div>
          <input
            id="search-query"
            name="search-query"
            onfocusin={(e) => {
              e.target.value = "";
              set_game("selected_option", undefined);
            }}
            oninput={(e) => {
              e.preventDefault();
              if (e.target.value != "") {
                const options = query_option(e.target.value).slice(0, 10);
                set_options([...options]);
              } else {
                set_options([]);
              }
            }}
            placeholder="Enter a movie title..."
            value={
              game.selected_option
                ? `${game.selected_option?.title} (${new Date(
                    game.selected_option?.release_date
                  ).getFullYear()}) (${game.selected_option?.id})`
                : ""
            }
            class="block w-full px-4 py-4 ps-10 text-sm text-river-bed-700 bg-desert-sand-100 border-2 border-river-bed-700 placeholder:text-river-bed-700 rounded-full focus:ring-river-bed-700 focus:border-river-bed-700"
            required
          />
        </div>
      ) : (
        <></>
      )}
      {options().length > 1 ? (
        <div
          class="z-10 bg-desert-sand-100 border-2 border-river-bed-700 divide-y divide-desert-sand-100 rounded-lg shadow w-full overflow-scroll no-scrollbar no-scrollbar::-webkit-scrollbar max-h-44"
          id="answer-options"
        >
          <ul class="py-2 text-sm text-river-bed-700">
            {options().map((option) => {
              return (
                <li
                  class="block px-4 py-2 hover:bg-desert-sand-200 hover:cursor-pointer"
                  onclick={() => {
                    set_game("selected_option", option);
                    set_options([]);
                  }}
                >
                  {option.title} ({new Date(option.release_date).getFullYear()})
                </li>
              );
            })}
          </ul>
        </div>
      ) : (
        <></>
      )}
      {game.status == GameStatus.won ? (
        <CorrectOption disabled={true} movie={game.correct_option} />
      ) : game.status == GameStatus.lost ? (
        <div>
          {game.selected_option ? (
            <div class="flex flex-col space-y-2">
              <CorrectOption disabled={true} movie={game.correct_option} />
              <WrongOption disabled={true} movie={game.selected_option} />
            </div>
          ) : (
            <WrongOption disabled={true} movie={game.correct_option} />
          )}
        </div>
      ) : (
        <></>
      )}
      <div></div>
    </div>
  );
}
