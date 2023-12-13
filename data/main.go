package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"mural-data/movies"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/ryanbradynd05/go-tmdb"
	"github.com/samber/lo"
)

func must(err error) {
	if err != nil {
		slog.Error(err.Error())
		panic(1)
	}
}

func exportToJSON(data interface{}, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		must(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	// Encode data and write to file
	err = encoder.Encode(data)
	if err != nil {
		must(err)
	}
}

func getDecadeFromReleaseDate(dateStr string) string {
	if dateStr == "" {
		return ""
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return ""
	}

	year := date.Year()
	switch {
	case year >= 2020:
		return movies.Theme2020
	case year >= 2010:
		return movies.Theme2010
	case year >= 2000:
		return movies.Theme2000
	case year >= 1990:
		return movies.Theme1990
	case year >= 1980:
		return movies.Theme1980
	case year >= 1970:
		return movies.Theme1970
	default:
		return movies.ThemeRandom
	}
}

func main() {
	must(godotenv.Load())
	tmdb_config := tmdb.Config{
		APIKey:   os.Getenv("TMDB_KEY"),
		Proxies:  nil,
		UseProxy: false,
	}

	tmdb_api := tmdb.Init(tmdb_config)

	// get all movies in JSON form
	movie_map := map[string][]movies.Movie{}
	all_movies := map[int]movies.Movie{}

	for i := 1; i < 10; i++ {
		fmt.Printf(".")
		for _, decade := range movies.DecadeOptions {
			movs, err := movies.GetMoviesByDecade(i, decade, tmdb_api)
			if err != nil {
				slog.Error(fmt.Errorf("could not get answers: %w", err).Error())
				return
			}

			for _, mov := range movs {
				all_movies[mov.ID] = mov
			}
		}
	}

	for _, mov := range all_movies {
		decade := getDecadeFromReleaseDate(mov.ReleaseDate)
		movs := lo.ValueOr[string, []movies.Movie](movie_map, decade, []movies.Movie{})
		movie_map[decade] = append(movs, mov)
	}

	exportToJSON(movie_map, "movies.json")
	fmt.Println()
	for decade, movies := range movie_map {
		fmt.Printf("%s: %d\n", decade, len(movies))
	}
}
