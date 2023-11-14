package test

import "mural/db"

var (
	MovMissionImpossible = db.Movie{
		Adult:         false,
		BackdropPath:  "/h56edmERPTkey89SqyKu4hINVNy.jpg",
		ID:            575264,
		OriginalTitle: "Mission: Impossible - Dead Reckoning Part One",
		Popularity:    2645.664,
		PosterPath:    "/NNxYkU70HPurnNCSiCjYAmacwm.jpg",
		ReleaseDate:   "2023-07-08",
		Title:         "Mission: Impossible - Dead Reckoning Part One",
		Overview:      "Ethan Hunt and his IMF team embark on their most dangerous mission yet: To track down a terrifying new weapon that threatens all of humanity before it falls into the wrong hands. With control of the future and the world's fate at stake and dark forces from Ethan's past closing in, a deadly race around the globe begins. Confronted by a mysterious, all-powerful enemy, Ethan must consider that nothing can matter more than his mission—not even the lives of those he cares about most.",
		Video:         false,
		VoteAverage:   7.7,
		VoteCount:     2198,
	}

	MovAriel = db.Movie{
		ID:            2,
		Title:         "Ariel",
		OriginalTitle: "Ariel",
		ReleaseDate:   "1988-10-21",
		Overview:      "After the coal mine he works at closes and his father commits suicide, a Finnish man leaves for the city to make a living but there, he is framed and imprisoned for various crimes.",
		VoteAverage:   7.0,
		VoteCount:     270,
		Popularity:    15.079999923706055,
		Adult:         false,
		Video:         false,
		PosterPath:    "/dQL2wJZo05GDd21VgOacMeCuyZy.jpg",
		BackdropPath:  "/ojDg0PGvs6R9xYFodRct2kdI6wC.jpg",
	}

	MovShadows = db.Movie{
		ID:            3,
		Title:         "Shadows in Paradise",
		OriginalTitle: "Varjoja paratiisissa",
		ReleaseDate:   "1986-10-17",
		Overview:      "Nikander, a rubbish collector and would-be entrepreneur finds his plans for success dashed when his business associate dies. One evening, he meets Ilona, a down-on-her luck cashier in a local supermarket—and, falteringly, a bond begins to develop between them.",
		VoteAverage:   7.2,
		VoteCount:     288,
		Popularity:    7.959,
		Adult:         false,
		Video:         false,
		PosterPath:    "/l94l89eMmFKh7na2a1u5q67VgNx.jpg",
		BackdropPath:  "/nj01hspawPof0mJmlgfjuLyJuRN.jpg",
	}

	MovFourRooms = db.Movie{
		ID:            5,
		Title:         "Four Rooms",
		OriginalTitle: "Four Rooms",
		ReleaseDate:   "1995-12-09",
		Overview:      "It's Ted the Bellhop's first night on the job...and the hotel's very unusual guests are about to place him in some outrageous predicaments. It seems that this evening's room service is serving up one unbelievable happening after another.",
		VoteAverage:   5.8,
		VoteCount:     2453,
		Popularity:    20.748,
		Adult:         false,
		Video:         false,
		PosterPath:    "/f2t4JbUvQIjUF5FstG1zZFAp02N.jpg",
		BackdropPath:  "/75aHn1NOYXh4M7L5shoeQ6NGykP.jpg",
	}

	MovJudgement = db.Movie{
		ID:            6,
		Title:         "Judgment Night",
		OriginalTitle: "Judgment Night",
		ReleaseDate:   "1993-10-15",
		Overview:      "While racing to a boxing match, Frank, Mike, John and Rey get more than they bargained for. A wrong turn lands them directly in the path of Fallon, a vicious, wise-cracking drug lord. After accidentally witnessing Fallon murder a disloyal henchman, the four become his unwilling prey in a savage game of cat and mouse as they are mercilessly stalked through the urban jungle in this taut suspense drama.",
		VoteAverage:   6.5,
		VoteCount:     303,
		Popularity:    12.086,
		Adult:         false,
		Video:         false,
		PosterPath:    "/bGMqHn0H2UY6SPZ5Vch4YJM2jDO.jpg",
		BackdropPath:  "/3rvvpS9YPM5HB2f4HYiNiJVtdam.jpg",
	}

	MovFindingNemo = db.Movie{
		ID:            12,
		Title:         "Finding Nemo",
		OriginalTitle: "Finding Nemo",
		ReleaseDate:   "2003-05-30",
		Overview:      "Nemo, an adventurous young clownfish, is unexpectedly taken from his Great Barrier Reef home to a dentist's office aquarium. It's up to his worrisome father Marlin and a friendly but forgetful fish Dory to bring Nemo home -- meeting vegetarian sharks, surfer dude turtles, hypnotic jellyfish, hungry seagulls, and more along the way.",
		VoteAverage:   7.8,
		VoteCount:     18182,
		Popularity:    94.569,
		Adult:         false,
		Video:         false,
		PosterPath:    "/h3uqFk7sZRJvLZDdLiFB9qwbL07.jpg",
		BackdropPath:  "/ggQ6o8X5984OCh3kZi2UIJQJY5y.jpg",
	}

	MovBlueBeetle = db.Movie{
		ID:            565770,
		Title:         "Blue Beetle",
		OriginalTitle: "Blue Beetle",
		ReleaseDate:   "2023-08-16",
		Overview:      "Recent college grad Jaime Reyes returns home full of aspirations for his future, only to find that home is not quite as he left it. As he searches to find his purpose in the world, fate intervenes when Jaime unexpectedly finds himself in possession of an ancient relic of alien biotechnology: the Scarab.",
		VoteAverage:   7.0,
		VoteCount:     1528,
		Popularity:    583.432,
		Adult:         false,
		Video:         false,
		PosterPath:    "/3H9NA1KWEQN0ItL3Wl3SFZYP6yV.jpg",
		BackdropPath:  "/mXLOHHc1Zeuwsl4xYKjKh2280oL.jpg",
	}
)

var AllMovies = []db.Movie{
	MovAriel,
	MovFindingNemo,
	MovFourRooms,
	MovJudgement,
	MovMissionImpossible,
	MovShadows,
}
