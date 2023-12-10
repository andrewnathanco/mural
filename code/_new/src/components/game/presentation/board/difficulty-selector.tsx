import { createEffect } from "solid-js";
import { useGame } from "../../context/game";
import { GameStatus } from "../../model/game";
import { useUser } from "../../context/game-difficulty";

export enum GameDifficulty {
  easy,
  hard,
}

export default function DifficultySelector() {
  const [game, set_game] = useGame();
  const [user, set_user] = useUser();

  return (
    <div class="flex space-x-2 items-center">
      <div
        class="inline-flex rounded-full text-sm text-river-bed-700 font-bold"
        role="group"
      >
        <button
          onclick={() => {
            if (game.flipped.length == 0) {
              set_user("difficulty", GameDifficulty.easy);
            }
          }}
          type="button"
          classList={{
            "bg-desert-sand-300 border-river-bed-700":
              user.difficulty == GameDifficulty.easy,
            "border-river-bed-700":
              user.difficulty != GameDifficulty.easy &&
              game.flipped.length == 0,
            "border-desert-sand-300 text-desert-sand-300":
              user.difficulty != GameDifficulty.easy &&
              game.flipped.length != 0,
          }}
          class="px-4 py-1 rounded-s-full border-r border-y-2 border-l-2"
        >
          Easy Mode
        </button>
        <button
          type="button"
          onclick={() => {
            if (game.flipped.length == 0) {
              set_user("difficulty", GameDifficulty.hard);
            }
          }}
          classList={{
            "bg-desert-sand-300 border-river-bed-700":
              user.difficulty == GameDifficulty.hard,
            "border-river-bed-700":
              user.difficulty != GameDifficulty.hard &&
              game.flipped.length == 0,
            "border-desert-sand-300 text-desert-sand-300":
              user.difficulty != GameDifficulty.hard &&
              game.flipped.length != 0,
          }}
          class=" px-4 py-1 rounded-e-full border-l border-y-2 border-r-2 border-river-bed-700 "
        >
          Hard Mode
        </button>
      </div>
    </div>
  );
}
