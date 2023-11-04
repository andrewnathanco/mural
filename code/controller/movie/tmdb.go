package movie

import (
	"fmt"
	"mural/api"
	"mural/model"
	"os"

	"github.com/ryanbradynd05/go-tmdb"
)


type TMDBController struct {
	TMDBApi *tmdb.TMDb
}

func NewTMDBAPI() *tmdb.TMDb {
	config := tmdb.Config{
		APIKey:   os.Getenv("TMDB_KEY"),
		Proxies:  nil,
		UseProxy: false,
	}

	return tmdb.Init(config)
}

func NewTMDBController() *TMDBController {
	api := NewTMDBAPI()
	return &TMDBController{
		TMDBApi: api,
	}
}

func (mc TMDBController) GetAnswers(
) (*model.Movie, []model.Answer, error) {
	parameters := map[string]string{
		"sort_by": "popularity.desc",
		"page": fmt.Sprintf("%d", api.RandomAnswerKey),
	}

	movie_results, err := mc.TMDBApi.DiscoverMovie(parameters)
	if err != nil {
		return nil,nil, err
	}


	var answers []model.Answer
	var correct_movie model.Movie
	for i, mov := range movie_results.Results {
		movie := model.Movie{
			Poster: mov.PosterPath,
			ID: mov.ID,
			Name: mov.Title,
			Description: mov.Overview,
			ReleaseDate: mov.ReleaseDate,
		}

		answer := model.Answer{
			Movie: movie,
			Selected: false,
			IsCorrect: i == api.RandomAnswerKey ,
		}

		if i == int(api.RandomAnswerKey)  {
			correct_movie = movie
		}

		if i >= 4 {
			break
		}

		answers = append(answers, answer)
	}


	return &correct_movie, answers, nil
}