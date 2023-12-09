import SubmitButton from "../../buttons/submit-button";
import GiveUpButton from "../../buttons/give-up-button";
import { useGame } from "../context";
import { GameStatus } from "../model";
import EasyModeInputArea from "./answer/easy-mode-area";
import HardModeInput from "./answer/hard-mode-input";
import { GameDifficulty } from "./board/difficulty-selector";
import ShareButton from "../../buttons/share-button";

export default function AnswerArea() {
  const [game, set_game] = useGame();

  return (
    <main class="w-full text-river-bed-600 font-extrabold flex flex-col space-y-4">
      {game.flipped.length == 0 ? (
        <></>
      ) : (
        <>
          <div class="h-0.5 w-full rounded-full bg-river-bed-600"></div>
          {game.status == GameStatus.won || game.status == GameStatus.lost ? (
            <>
              <ShareButton />
            </>
          ) : (
            <></>
          )}
          {game.difficulty == GameDifficulty.easy ? (
            <EasyModeInputArea />
          ) : (
            <HardModeInput />
          )}

          {game.status == GameStatus.won || game.status == GameStatus.lost ? (
            <> </>
          ) : (
            <div class="flex flex-col space-y-2">
              <SubmitButton />
              {game.difficulty == GameDifficulty.hard ? (
                <GiveUpButton />
              ) : (
                <></>
              )}
            </div>
          )}
        </>
      )}
    </main>
  );
}
