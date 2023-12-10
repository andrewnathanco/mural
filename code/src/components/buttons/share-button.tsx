import { useShareDialog } from "../dialog/share/context";
import { generate_share_data } from "../dialog/share/share-dialog";
import { useGame } from "../game/context/game";
import { BoardState, GameStatus } from "../game/model/game";

export default function ShareButton(props: { onclick: Function }) {
  const [game, set_game] = useGame();

  return (
    <div
      onclick={() => {
        props.onclick();
      }}
      class="w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center border-dingley-700 border-2 hover:cursor-pointer bg-dingley-700 hover:bg-dingley-800"
    >
      Share
    </div>
  );
}
