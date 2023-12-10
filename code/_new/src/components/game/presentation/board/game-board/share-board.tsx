import { createEffect, createSignal } from "solid-js";
import { useGame } from "../../../context/game";
import { DefaultTile, FlippedTile, SelectedTile, getTiles } from "./game-board";
import { BoardState, GameStatus } from "../../../model/game";
import ShareBoardStateSelector from "../state-selector/share-state-selector";
import CorrectOption from "../../answer/input/correct-option";
import WrongOption from "../../answer/input/wrong-option";

export default function ShareBoard(props: { flipped: number[] }) {
  const [game, set_game] = useGame();

  const [board_state, set_board_state] = createSignal(BoardState.game);
  const [all_tiles, set_all_tiles] = createSignal(getTiles(props.flipped));

  return (
    <div class="flex flex-col space-y-4 items-center">
      <div
        class="flex flex-col border-2 border-river-bed-800 bg-cover"
        style={`background-image: url('https://image.tmdb.org/t/p/w1280/${game.correct_option.poster_path}');`}
      >
        {all_tiles().map((tile_row) => {
          return (
            <div class="flex text-sm">
              {tile_row.map((tile) => {
                if (board_state() == BoardState.flipped) {
                  return <FlippedTile />;
                }

                if (tile.flipped) {
                  return <FlippedTile />;
                } else {
                  return <DefaultTile tile={tile} disabled={false} />;
                }
              })}
            </div>
          );
        })}
      </div>
      <ShareBoardStateSelector
        board_state_signal={[board_state, set_board_state]}
      />
    </div>
  );
}
