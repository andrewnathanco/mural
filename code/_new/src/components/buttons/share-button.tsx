import { useGame } from "../game/context";
import { BoardState, GameStatus } from "../game/model";

export default function ShareButton() {
  const [game, set_game] = useGame();

  return (
    <div
      onclick={() => {}}
      class="w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center border-dingley-700 border-2 hover:cursor-pointer bg-dingley-700 hover:bg-dingley-800"
    >
      Share
    </div>
  );
}
