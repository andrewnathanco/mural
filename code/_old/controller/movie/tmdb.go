package movie

import (
	"fmt"
	"mural/db"

	"github.com/ryanbradynd05/go-tmdb"
)

type TMDBController struct {
	TMDBApi *tmdb.TMDb
}

func NewTMDBAPI(
	tmdb_key string,
) *tmdb.TMDb {
	config := tmdb.Config{
		APIKey:   tmdb_key,
		Proxies:  nil,
		UseProxy: false,
	}

	return tmdb.Init(config)
}

func NewTMDBController(
	tmdb_key string,
) TMDBController {
	api := NewTMDBAPI(tmdb_key)
	return TMDBController{
		TMDBApi: api,
	}
}

func (mc TMDBController) GetMoviesByDecade(
	page_number int,
	decade string,
) ([]db.Movie, error) {
	lower_bound, upper_boud := getDecadeBounds(decade)

	parameters := map[string]string{
		"page":                     fmt.Sprintf("%d", page_number),
		"primary_release_date.lte": upper_boud,
		"primary_release_date.gte": lower_bound,
	}

	movie_results, err := mc.TMDBApi.DiscoverMovie(parameters)
	if err != nil {
		return nil, err
	}

	movies := []db.Movie{}
	for _, movie_short := range movie_results.Results {
		movie := db.ConvertShortToMovies(movie_short)
		movies = append(movies, movie)
	}

	return movies, nil
}
