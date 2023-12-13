import { useSearchParams } from "@solidjs/router";
import { GameProvider, useGame } from "../components/game/context/game";
import CorrectOption from "../components/game/presentation/answer/input/correct-option";
import WrongOption from "../components/game/presentation/answer/input/wrong-option";
import ShareBoard from "../components/game/presentation/board/game-board/share-board";
import { Movie } from "../components/movie/model";
import { get_movie_from_id } from "../components/movie/service";

function ShareBody(props: {
  name: string;
  flipped: string;
  answer: string;
  correct: string;
}) {
  const [game, set_game] = useGame();
  const answer = get_movie_from_id(parseInt(props.answer));
  const correct = get_movie_from_id(parseInt(props.correct));

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
          ) : (
            <div class="flex flex-col space-y-2">
              <CorrectOption disabled={true} movie={correct} />
              <WrongOption disabled={true} movie={answer} />
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default function Share() {
  const [params, _] = useSearchParams();
  const name = params["name"];
  const flipped = params["flipped"];
  const correct = params["correct_id"];
  const answer = params["answer_id"];

  return (
    <GameProvider>
      {name && flipped && correct ? (
        <ShareBody
          name={name}
          flipped={flipped}
          answer={answer}
          correct={correct}
        />
      ) : (
        <div class="text-4xl">
          Sharing not working for that user. This is probably Andrew's fault.
        </div>
      )}
    </GameProvider>
  );
}
