import SubmitButton from "../../buttons/submit-button";
import GiveUpButton from "../../buttons/give-up-button";
import { GameStatus } from "../model/game";
import EasyModeInputArea from "./answer/easy-mode-area";
import HardModeInput from "./answer/hard-mode-input";
import { GameDifficulty } from "./board/difficulty-selector";
import { useGame } from "../context/game";
import { useUser } from "../context/game-difficulty";
import { ShareButton } from "../../buttons/share-button";

export default function AnswerArea() {
  const [game, set_game] = useGame();
  const [user, __] = useUser();

  return (
    <main class="w-full text-river-bed-600 font-extrabold flex flex-col space-y-4">
      {game.flipped.length == 0 ? (
        <></>
      ) : (
        <>
          <div class="h-0.5 w-full rounded-full bg-river-bed-600"></div>
          {game.status == GameStatus.won || game.status == GameStatus.lost ? (
            <>
              <ShareButton
                onclick={() => {
                  window.scrollTo(0, 0);
                  open();
                }}
              />
            </>
          ) : (
            <></>
          )}
          {user.difficulty == GameDifficulty.easy ? (
            <EasyModeInputArea />
          ) : (
            <HardModeInput />
          )}

          {game.status == GameStatus.won || game.status == GameStatus.lost ? (
            <> </>
          ) : (
            <div class="flex flex-col space-y-2">
              <SubmitButton />
              {user.difficulty == GameDifficulty.hard ? (
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
