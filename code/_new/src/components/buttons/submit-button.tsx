import { useGame } from "../game/context/game";
import { BoardState, GameStatus } from "../game/model";

export default function SubmitButton() {
  const [game, set_game] = useGame();

  return (
    <div
      onclick={() => {
        let game_status =
          game.selected_option?.id == game.correct_option.id
            ? GameStatus.won
            : GameStatus.lost;

        set_game("selected_tile", undefined);
        set_game("score", game_status == GameStatus.lost ? "❎" : game.score);
        set_game("status", game_status);
        set_game("board_state", BoardState.flipped);
      }}
      classList={{
        "bg-river-bed-400": !game.selected_option,
        "border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700":
          !!game.selected_option,
      }}
      class=" w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center "
    >
      Submit
    </div>
  );
}
