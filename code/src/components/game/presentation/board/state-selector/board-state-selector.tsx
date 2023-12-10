import { useGame } from "../../../context/game";
import { BoardState, GameStatus } from "../../../model/game";

export default function BoardStateSelector() {
  const [game, set_game] = useGame();

  return (
    <div class="flex space-x-2 items-center">
      <div
        class="inline-flex rounded-full text-sm text-river-bed-700 font-bold"
        role="group"
      >
        <button
          onclick={() => {
            set_game("board_state", BoardState.flipped);
          }}
          type="button"
          classList={{
            "bg-desert-sand-300 border-river-bed-700":
              game.board_state == BoardState.flipped,
            "border-river-bed-700": game.board_state != BoardState.flipped,
          }}
          class="px-4 py-1 rounded-s-full border-r border-y-2 border-l-2"
        >
          Flipped
        </button>
        <button
          type="button"
          onclick={() => {
            set_game("board_state", BoardState.game);
          }}
          classList={{
            "bg-desert-sand-300 border-river-bed-700":
              game.board_state == BoardState.game,
            "border-river-bed-700": game.board_state != BoardState.game,
          }}
          class=" px-4 py-1 rounded-e-full border-l border-y-2 border-r-2 border-river-bed-700 "
        >
          Game
        </button>
      </div>
    </div>
  );
}
