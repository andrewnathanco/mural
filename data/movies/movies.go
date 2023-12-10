package movies

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/ryanbradynd05/go-tmdb"
)

func must(err error) {
	if err != nil {
		slog.Error(err.Error())
		panic(1)
	}
}

type Genre struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type GenresList struct {
	Genres []Genre `json:"genres"`
}

const (
	Theme2020   = "2020"
	Theme2010   = "2010"
	Theme2000   = "2000"
	Theme1990   = "1990"
	Theme1980   = "1980"
	Theme1970   = "1970"
	ThemeRandom = "Random"
)

var (
	DecadeOptions = []string{
		Theme2020,
		Theme2010,
		Theme2000,
		Theme1990,
		Theme1980,
		Theme1970,
	}
)

type Movie struct {
	Genres      []Genre `json:"genres"`
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	Overview    string  `json:"overview"`
	VoteCount   uint32  `json:"vote_count"`
	PosterPath  string  `json:"poster_path"`
}

func getDecadeBounds(decade string) (string, string) {
	decadeStart, _ := strconv.Atoi(decade)
	decadeEnd := decadeStart + 9

	return strconv.Itoa(decadeStart - 1), strconv.Itoa(decadeEnd + 1)
}

func ConvertShortToMovies(tmdbShort tmdb.MovieShort) Movie {
	var all_genres GenresList
	file_data, err := os.ReadFile("./genres.json")
	must(err)
	must(json.Unmarshal(file_data, &all_genres))

	this_genres := []Genre{}
	for _, genre := range all_genres.Genres {
		for _, genre_int := range tmdbShort.GenreIDs {
			if genre.ID == genre_int {
				this_genres = append(this_genres, genre)
			}
		}
	}

	return Movie{
		ID:          tmdbShort.ID,
		Title:       tmdbShort.Title,
		ReleaseDate: tmdbShort.ReleaseDate,
		Overview:    tmdbShort.Overview,
		VoteCount:   tmdbShort.VoteCount,
		PosterPath:  tmdbShort.PosterPath,
		Genres:      this_genres,
	}
}

func GetMoviesByDecade(
	page_number int,
	decade string,
	tmdb_api *tmdb.TMDb,
) ([]Movie, error) {
	lower_bound, upper_boud := getDecadeBounds(decade)

	parameters := map[string]string{
		"page":                     fmt.Sprintf("%d", page_number),
		"primary_release_date.lte": upper_boud,
		"primary_release_date.gte": lower_bound,
		"sort_by":                  "vote_count.desc",
		"include_adult":            "false",
		"with_original_language":   "en",
	}

	movie_results, err := tmdb_api.DiscoverMovie(parameters)
	if err != nil {
		return nil, err
	}

	movies := []Movie{}
	for _, movie_short := range movie_results.Results {
		movie := ConvertShortToMovies(movie_short)
		movies = append(movies, movie)
	}

	return movies, nil
}
