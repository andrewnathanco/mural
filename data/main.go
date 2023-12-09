package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"mural-data/movies"
	"os"

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
	for i := 1; i < 50; i++ {
		fmt.Printf(".")
		for _, decade := range movies.DecadeOptions {
			movs, err := movies.GetMoviesByDecade(i, decade, tmdb_api)
			if err != nil {
				slog.Error(fmt.Errorf("could not get answers: %w", err).Error())
				return
			}

			existing_movies := lo.ValueOr(movie_map, decade, []movies.Movie{})
			movie_map[decade] = append(existing_movies, movs...)
		}
	}

	exportToJSON(movie_map, "movies.json")
	fmt.Println()
	for decade, movies := range movie_map {
		fmt.Printf("%s: %d\n", decade, len(movies))
	}
}
