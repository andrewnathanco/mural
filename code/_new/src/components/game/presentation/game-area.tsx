import FlipButton from "../../buttons/flip-button";
import { useGame } from "../context";
import AnswerArea from "./answer-area";
import DifficultySelector from "./board/difficulty-selector";
import GameBoard from "./board/game-board";

export default function GameArea() {
  const [game, _] = useGame();

  return (
    <main class="text-river-bed-600 font-extrabold flex flex-col items-center justify-center space-y-4">
      <div class="flex flex-col items-center space-y-2">
        <div class="flex space-x-2 text-4xl">
          <div>Score:</div>
          <div>{game.score}</div>
        </div>
        <div>
          <DifficultySelector />
        </div>
      </div>

      <GameBoard />
      <AnswerArea />
    </main>
  );
}
