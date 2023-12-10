import { Movie } from "../../../../movie/model";

export default function WrongOption(props: {
  disabled: boolean;
  movie: Movie;
}) {
  const disabled = props.disabled;
  const movie = props.movie;

  return (
    <div
      classList={{ "hover:cursor-pointer": !disabled }}
      class="text-lg w-full rounded-lg bg-el-salva-300 text-el-salva-800 border-4 border-el-salva-800 py-2 px-4"
    >
      {movie.title} ({new Date(movie.release_date).getFullYear()})
    </div>
  );
}
