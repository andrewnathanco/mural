import { useGame } from "../game/context";
import { GameStatus } from "../game/model";
import { get_penalty_from_key } from "../game/service";

export default function FlipButton() {
  const [game, set_game] = useGame();

  return (
    <div
      onclick={() => {
        if (game.selected_tile != undefined) {
          set_game("flipped", (flipped) => {
            return [...flipped, game.selected_tile as number];
          });

          if (typeof game.score == "number") {
            set_game(
              "score",
              (game.score as number) - get_penalty_from_key(game.selected_tile)
            );
          }

          set_game("status", GameStatus.started);
          set_game("selected_tile", undefined);
        }
      }}
      classList={{
        "bg-river-bed-400": game.selected_tile == undefined,
        "border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700":
          game.selected_tile != undefined,
      }}
      class=" w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center "
    >
      Flip
    </div>
  );
}
