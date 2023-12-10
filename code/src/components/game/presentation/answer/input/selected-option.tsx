import { Movie } from "../../../../movie/model";

export default function SelectedOption(props: {
  disabled: boolean;
  movie: Movie;
}) {
  const disabled = props.disabled;
  const movie = props.movie;

  return (
    <div
      classList={{ "hover:cursor-pointer": !disabled }}
      class="text-lg w-full rounded-lg bg-desert-sand-300 text-river-bed-700 border-4 border-river-bed-800 py-2 px-4"
    >
      {movie.title} ({new Date(movie.release_date).getFullYear()})
    </div>
  );
}
