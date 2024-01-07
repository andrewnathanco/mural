import { useSearchParams } from "@solidjs/router";
import { Title } from "@solidjs/meta";
import { GameProvider, useGame } from "../components/game/context/game";
import CorrectOption from "../components/game/presentation/answer/input/correct-option";
import WrongOption from "../components/game/presentation/answer/input/wrong-option";
import ShareBoard from "../components/game/presentation/board/game-board/share-board";
import { Movie } from "../components/movie/model";
import { get_movie_from_id } from "../components/movie/service";
import { ShareDialogProvider } from "../components/dialog/share/context";
import {
  ShareWarningDialogProvider,
  useShareWarningDialog,
} from "../components/dialog/share-warning/context";
import ShareWarningDialog from "../components/dialog/share-warning/share-warning-dialog";
import { createEffect } from "solid-js";
import { GameStatus } from "../components/game/model/game";
import { get_todays_game } from "../components/game/service";

function ShareBody(props: {
  name: string;
  flipped: string;
  answer: string;
  correct: string;
}) {
  const answer_id = props.answer;
  const [game, set_game] = useGame();
  const answer = get_movie_from_id(parseInt(answer_id));
  const correct = get_movie_from_id(parseInt(props.correct));
  const [_, { open }] = useShareWarningDialog();

  createEffect(() => {
    if (
      game.status == GameStatus.init ||
      game.status == GameStatus.started ||
      game.game_key != get_todays_game().game_key
    ) {
      open();
    }
  });

  return (
    <div class="flex flex-col items-center justify-center">
      <div class="flex flex-col items-center space-y-4">
        <div class="flex flex-col space-y-4 w-full">
          <div class="text-5xl flex space-x-2 items-center">
            <div>
              {props.name ? props.name + "'s " : ""}Mural #{game.game_key}
            </div>
            <div
              id="game-version"
              class="font-semibold w-min h-min text-gray-600 text-xs border-2 px-1 border-river-bed-700 rounded-lg"
            >
              {import.meta.env.VITE_VERSION ?? "v0.1.1"}
            </div>
          </div>
          <div class="flex justify-between">
            <div class="flex flex-col space-y-1 items-start">
              <div id="game-theme" class="text-contessa-500 text-4xl">
                {game.theme}
              </div>
              <div class="text-md">Today's Theme</div>
            </div>
          </div>
          <div class="h-0.5 w-full rounded-full bg-river-bed-600"></div>
          <ShareBoard
            flipped={props.flipped.split(",").map((val) => parseInt(val))}
            correct={correct}
          />
          <div class="text-3xl flex space-x-2 items-center flex-col"></div>
          {correct.id == answer.id ? (
            <CorrectOption disabled={true} movie={correct} />
          ) : answer_id != undefined ? (
            <div class="flex flex-col space-y-2">
              <CorrectOption disabled={true} movie={correct} />
              <WrongOption disabled={true} movie={answer} />
            </div>
          ) : (
            <WrongOption disabled={true} movie={correct} />
          )}
        </div>
      </div>
      <ShareWarningDialog />
    </div>
  );
}

export default function App() {
  const [params, _] = useSearchParams();
  const name = params["name"];
  const flipped = params["flipped"];
  const correct = params["correct_id"];
  const answer = params["answer_id"];

  return (
    <>
      <Title>{name ? `${name}'s Mural` : "Mural Share"}</Title>
      <ShareWarningDialogProvider>
        <GameProvider>
          {flipped && correct ? (
            <ShareBody
              name={name || ""}
              flipped={flipped}
              answer={answer || ""}
              correct={correct}
            />
          ) : (
            <div class="text-4xl">
              Sharing not working for that user. This is probably Andrew's
              fault.
            </div>
          )}
        </GameProvider>
      </ShareWarningDialogProvider>
    </>
  );
}
