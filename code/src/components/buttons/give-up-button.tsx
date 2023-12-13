import { useGame } from "../game/context/game";
import { BoardState, GameStatus } from "../game/model/game";
import { update_number_of_games_played } from "../game/service";

export default function GiveUpButton() {
  const [game, set_game] = useGame();
  return (
    <button
      onclick={() => {
        set_game("selected_option", undefined);
        set_game("selected_tile", undefined);
        set_game("score", "❎");
        set_game("status", GameStatus.lost);
        set_game("board_state", BoardState.flipped);

        update_number_of_games_played();
      }}
      class="w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center hover:cursor-pointer bg-contessa-600"
    >
      Give Up
    </button>
  );
}
