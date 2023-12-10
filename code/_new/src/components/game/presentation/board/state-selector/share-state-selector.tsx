import { Signal, createSignal } from "solid-js";
import { BoardState } from "../../../model";

export default function ShareBoardStateSelector(props: {
  board_state_signal: Signal<BoardState>;
}) {
  const [board_state, set_board_state] = props.board_state_signal;

  return (
    <div class="flex space-x-2 items-center">
      <div
        class="inline-flex rounded-full text-sm text-river-bed-700 font-bold"
        role="group"
      >
        <button
          onclick={() => {
            set_board_state(BoardState.flipped);
          }}
          type="button"
          classList={{
            "bg-desert-sand-300 border-river-bed-700":
              board_state() == BoardState.flipped,
            "border-river-bed-700": board_state() != BoardState.flipped,
          }}
          class="px-4 py-1 rounded-s-full border-r border-y-2 border-l-2"
        >
          Flipped
        </button>
        <button
          type="button"
          onclick={() => {
            set_board_state(BoardState.game);
          }}
          classList={{
            "bg-desert-sand-300 border-river-bed-700":
              board_state() == BoardState.game,
            "border-river-bed-700": board_state() != BoardState.game,
          }}
          class=" px-4 py-1 rounded-e-full border-l border-y-2 border-r-2 border-river-bed-700 "
        >
          Game
        </button>
      </div>
    </div>
  );
}
