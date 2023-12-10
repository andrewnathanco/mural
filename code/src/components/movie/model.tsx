export interface Genre {
  id: number;
  name: string;
}

export interface Movie {
  genres: Genre[];
  id: number;
  title: string;
  release_date: string;
  overview: string;
  vote_count: number;
  poster_path: string;
}
