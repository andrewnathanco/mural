import { useGame } from "../game/context";
import { get_penalty_from_key } from "../game/service";

export default function FlipButton() {
  const [game, set_game] = useGame();

  return (
    <div
      onclick={() => {
        if (game.selected != undefined) {
          set_game("flipped", (flipped) => {
            return [...flipped, game.selected as number];
          });

          set_game("score", game.score - get_penalty_from_key(game.selected));
          set_game("selected", undefined);
        }
      }}
      classList={{
        "bg-river-bed-400": game.selected == undefined,
        "border-river-bed-600 border-2 hover:cursor-pointer bg-river-bed-600 hover:bg-river-bed-700":
          game.selected != undefined,
      }}
      class=" w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center "
    >
      Flip
    </div>
  );
}
