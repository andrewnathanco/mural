import { useGame } from "../../context";
import { GameStatus } from "../../model";
import CorrectOption from "./input/correct-option";
import DefaultOption from "./input/default-option";
import SelectedOption from "./input/selected-option";
import WrongOption from "./input/wrong-option";

export default function EasyModeInputArea() {
  const [game, set_game] = useGame();

  return (
    <ul class="w-full flex flex-col space-y-1">
      {game.easy_mode_options.map((option) => {
        if (game.selected_option?.id == option.id) {
          if (
            option.id == game.correct_option?.id &&
            game.status != GameStatus.init &&
            game.status != GameStatus.started
          ) {
            return <CorrectOption disabled={true} movie={option} />;
          }

          if (
            option.id != game.correct_option?.id &&
            game.status != GameStatus.init &&
            game.status != GameStatus.started
          )
            return <WrongOption disabled={true} movie={option} />;

          return (
            <SelectedOption
              disabled={
                !(
                  game.status != GameStatus.init &&
                  game.status != GameStatus.started
                )
              }
              movie={option}
            />
          );
        }

        if (
          option.id == game.correct_option?.id &&
          game.status != GameStatus.init &&
          game.status != GameStatus.started
        ) {
          return <CorrectOption disabled={true} movie={option} />;
        }

        return (
          <DefaultOption
            disabled={
              game.status != GameStatus.init &&
              game.status != GameStatus.started
            }
            movie={option}
          />
        );
      })}
    </ul>
  );
}
