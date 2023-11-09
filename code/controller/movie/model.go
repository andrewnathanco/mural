package movie

type MovieShort struct {
	Adult         bool    `json:"adult"`
	BackdropPath  string  `json:"backdrop_path"`
	ID            int     `json:"id"`
	OriginalTitle string  `json:"original_title"`
	GenreIDs      []int32 `json:"genre_ids"`
	Popularity    float32 `json:"popularity"`
	PosterPath    string  `json:"poster_path"`
	ReleaseDate   string  `json:"release_date"`
	Title         string  `json:"title"`
	Overview      string  `json:"overview"`
	Video         bool    `json:"video"`
	VoteAverage   float32 `json:"vote_average"`
	VoteCount     uint32  `json:"vote_count"`
}

// MoviePagedResults struct
type MoviePagedResults struct {
	ID                int
	Page              int
	Results           []MovieShort
	TotalPages        int                     `json:"total_pages"`
	TotalResults      int                     `json:"total_results"`
}