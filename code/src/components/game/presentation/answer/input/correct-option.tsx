import { Movie } from "../../../../movie/model";

export default function CorrectOption(props: {
  disabled: boolean;
  movie: Movie;
}) {
  const disabled = props.disabled;
  const movie = props.movie;

  return (
    <div
      classList={{ "hover:cursor-pointer": !disabled }}
      class="text-lg w-full rounded-lg bg-dingley-300 text-dingley-800 border-4 border-dingley-800 py-2 px-4"
    >
      {movie.title} ({new Date(movie.release_date).getFullYear()})
    </div>
  );
}
