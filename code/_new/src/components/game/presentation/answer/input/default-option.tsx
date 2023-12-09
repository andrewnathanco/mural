import { Movie } from "../../../../movie/model";
import { useGame } from "../../../context";

export default function DefaultOption(props: {
  disabled: boolean;
  movie: Movie;
}) {
  const [game, set_game] = useGame();
  const disabled = props.disabled;
  const movie = props.movie;

  return (
    <button
      onclick={() => {
        if (!disabled) {
          set_game("selected_option", { ...movie });
        }
      }}
      classList={{ "hover:cursor-pointer": !disabled }}
      class="text-left text-lg w-full rounded-lg py-2 px-4 bg-transparent text-river-bed-700 border-2 border-river-bed-700 "
    >
      {movie.title} ({new Date(movie.release_date).getFullYear()})
    </button>
  );
}
