package main

import "fmt"

type Movie struct {
	Title  string
	Year   int
	Rating float64
	Genres []string
}

func TopRatedMovie(movies []Movie) Movie {
	if len(movies) == 0 {
		return Movie{}
	}

	topMovie := movies[0]
	for _, movie := range movies {
		if movie.Rating > topMovie.Rating {
			topMovie = movie
		}
	}

	return topMovie
}


func addGenres(movie *Movie, newGenres ...string) {
	movie.Genres = append(movie.Genres, newGenres...)
}


func searchByGenre(movies []Movie, genre string) []Movie {
	var result []Movie

	for _, movie := range movies {
		for _, g := range movie.Genres {
			if g == genre {
				result = append(result, movie)
				break
			}
		}
	}

	return result
}

func main() {
	movies := []Movie{
		{Title: "Начало", Year: 2010, Rating: 8.8, Genres: []string{"фантастика", "боевик"}},
		{Title: "Интерстеллар", Year: 2014, Rating: 8.6, Genres: []string{"фантастика", "драма"}},
		{Title: "Побег из Шоушенка", Year: 1994, Rating: 9.3, Genres: []string{"драма"}},
		{Title: "Тёмный рыцарь", Year: 2008, Rating: 9.0, Genres: []string{"боевик", "драма"}},
		{Title: "Криминальное чтиво", Year: 1994, Rating: 8.9, Genres: []string{"криминал", "драма"}},
	}

	topMovie := TopRatedMovie(movies)
	fmt.Printf("Фильм с самым высоким рейтингом: %s (Рейтинг: %.1f)\n\n", topMovie.Title, topMovie.Rating)

	addGenres(&movies[0], "триллер")
	fmt.Printf("Обновленные жанры фильма '%s': %v\n\n", movies[0].Title, movies[0].Genres)

	foundMovies := searchByGenre(movies, "боевик")
	fmt.Println("Фильмы в жанре 'боевик':")
	for _, m := range foundMovies {
		fmt.Printf("- %s (%d г., Рейтинг: %.1f)\n", m.Title, m.Year, m.Rating)
	}
}