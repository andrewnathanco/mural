import { useGame } from "../game/context/game";
import { useUser } from "../game/context/game-difficulty";
import { BoardState, Game, GameStatus } from "../game/model/game";
import { User } from "../game/model/user";
import { GameDifficulty } from "../game/presentation/board/difficulty-selector";

function create_share_url(game: Game) {
  let share_url = `${import.meta.env.VITE_SHARE_URL}/share`;

  let flipped = game.flipped.join(",");

  share_url += `?flipped=${flipped}`;
  share_url += `&name=${game.user_name}`;
  if (game.selected_option) {
    share_url += `&answer_id=${game.selected_option?.id}`;
  }
  share_url += `&correct_id=${game.correct_option.id}`;

  return share_url;
}

function generate_share_data(game: Game, user: User): ShareData {
  return {
    url: create_share_url(game),
    title: "Mural Share",
    text: `${game.user_name ? `${game.user_name}'s ` : ""}Mural${
      user.difficulty == GameDifficulty.hard ? "*" : ""
    } #${game.game_key}\nScore: ${game.score}`,
  };
}

export function ShareButton(props: { onclick: Function }) {
  const [game, set_game] = useGame();
  const [user, _] = useUser();


  return (
    <div
      onclick={() => {
        const data = generate_share_data(game, user)
        try {
          navigator?.share(data);
        } catch(e) {
        }
      }}
      class="w-full p-1 text-base text-desert-sand-200 rounded-md flex justify-center items-center border-dingley-700 border-2 hover:cursor-pointer bg-dingley-700 hover:bg-dingley-800"
    >
      Share
    </div>
  );
}
