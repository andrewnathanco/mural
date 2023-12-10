import { createEffect, createSignal } from "solid-js";
import { BoardState, GameStatus, GameTile } from "../../../model/game";
import FlipButton from "../../../../buttons/flip-button";
import BoardStateSelector from "../state-selector/board-state-selector";
import HintButton from "../../../../buttons/hints-button";
import { useGame } from "../../../context/game";

export function getTiles(flipped: number[]) {
  const size = 10;
  const all_tiles: GameTile[][] = [];

  // build board
  for (let row = 0; row < size; row++) {
    const row_tiles = [];
    for (let col = 0; col < 10; col++) {
      const ring = Math.min(row, col, size - row - 1, size - col - 1);

      const key = parseInt(`${row}${col}`);

      row_tiles.push({
        flipped: flipped.includes(key),
        penalty: ring + 1,
        row,
        col,
        key,
      });
    }

    all_tiles.push(row_tiles);
  }

  return all_tiles;
}

export default function GameBoard() {
  const [game, set_game] = useGame();
  const [all_tiles, set_all_tiles] = createSignal(getTiles([]));

  createEffect(() => {
    set_all_tiles(getTiles(game.flipped));
  });

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
                if (game?.board_state == BoardState.flipped) {
                  return <FlippedTile />;
                }

                if (game?.selected_tile == tile.key) {
                  return <SelectedTile tile={tile} />;
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
      {game.status == GameStatus.started ? (
        <div class="w-full flex flex-col items-center space-y-2">
          <FlipButton />
          <HintButton />
        </div>
      ) : game.status != GameStatus.init ? (
        <BoardStateSelector />
      ) : (
        <></>
      )}
    </div>
  );
}

export function FlippedTile() {
  return (
    <div class="flex flex-col justify-center items-center md:w-16 md:h-16 w-8 h-8 border-2 border-river-bed-800"></div>
  );
}

export function SelectedTile(props: { tile: GameTile }) {
  let tile = props.tile;

  return (
    <div
      classList={{
        "bg-contessa-500":
          (tile.col % 2 == 0 && tile.row % 2 == 1) ||
          (tile.col % 2 == 1 && tile.row % 2 == 0),
        "bg-contessa-400":
          (tile.col % 2 == 1 && tile.row % 2 == 1) ||
          (tile.col % 2 == 0 && tile.row % 2 == 0),
      }}
      class="flex flex-col justify-center items-center md:w-16 md:h-16 w-8 h-8 md:text-xl text-xs grow-lg rounded-lg md:border-8 border-4 border-river-bed-600 "
    >
      -{tile.penalty}
    </div>
  );
}

export function DefaultTile(props: { tile: GameTile; disabled: boolean }) {
  const [game, set_game] = useGame();
  let tile = props.tile;
  let disabled = props.disabled;

  return (
    <button
      onclick={() => {
        set_game("selected_tile", tile.key);
        set_game("status", GameStatus.started);
      }}
      classList={{
        "bg-contessa-500":
          (tile.col % 2 == 0 && tile.row % 2 == 1) ||
          (tile.col % 2 == 1 && tile.row % 2 == 0),
        "bg-contessa-400":
          (tile.col % 2 == 1 && tile.row % 2 == 1) ||
          (tile.col % 2 == 0 && tile.row % 2 == 0),
        "cursor-pointer": !disabled,
      }}
      class="flex flex-col justify-center items-center md:w-16 md:h-16 w-8 h-8 border-2 border-river-bed-800"
    ></button>
  );
}
