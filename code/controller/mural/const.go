package mural

import (
	"mural/model"
	"time"

	"github.com/google/uuid"
)

var (
	bbReleaseDate, _ = time.Parse("2006-01-02", "2023-08-16")
	MovBlueBeetle = model.Movie{
		Name: "Blue Beetle",
		ReleaseDate: bbReleaseDate,
		ID: uuid.New(),
		Poster: "/H6j5smdpRqP9a8UnhWp6zfl0SC.jpg",
		Description: "Recent college grad Jaime Reyes returns home full of aspirations for his future, only to find that home is not quite as he left it. As he searches to find his purpose in the world, fate intervenes when Jaime unexpectedly finds himself in possession of an ancient relic of alien biotechnology: the Scarab.",
	}

	ssReleaseDate, _ = time.Parse("2006-01-02", "2019-01-11")
	MovSurgStrike = model.Movie{
		Name: "Uri: The Surgical Strike",
		ReleaseDate: ssReleaseDate,
		ID: uuid.New(),
		Poster: "/yNySAgpAnWmPpYinim9E0tUzJWG.jpg",
		Description: "Following the roguish terrorist attacks at Uri Army Base camp in Kashmir, India takes the fight to the enemy, in its most successful covert operation till date with one and only one objective of avenging their fallen heroes.",
	}


	snipeReleaseDate, _ = time.Parse("2006-01-02", "2023-09-26")
	MovSniper = model.Movie{
		Name: "Sniper: G.R.I.T. - Global Response & Intelligence Team",
		ReleaseDate: snipeReleaseDate,
		ID: uuid.New(),
		Poster: "/a9bt9byTQ1MIfRWYQX240HiYPrl.jpg",
		Description: "When an international terrorist cult threatens global political stability and kidnaps a fellow agent, Ace Sniper Brandon Beckett and the newly-formed Global Response & Intelligence Team - or G.R.I.T. - led by Colonel Stone must travel across the world to Malta, infiltrate the cult, and take out its leader to free Lady Death and stop the global threat.",
	}

	megReleaseDate, _ = time.Parse("2006-01-02", "2023-08-02")
	MovMeg = model.Movie{
		Name: "Meg 2: The Trench",
		ReleaseDate: megReleaseDate,
		ID: uuid.New(),
		Poster: "/4m1Au3YkjqsxF8iwQy0fPYSxE0h.jpg",
		Description: "An exploratory dive into the deepest depths of the ocean of a daring research team spirals into chaos when a malevolent mining operation threatens their mission and forces them into a high-stakes battle for survival.",
	}
)