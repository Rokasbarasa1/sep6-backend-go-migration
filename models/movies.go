package models

import (
	"express-to-gin/connections"
	"express-to-gin/types"
	"strconv"
)

// const escapeSansQuotes(connection, val){
// return connection.escape(val).match(/^'(\w+)'$/)[1];
// }

func GetAllMoviesWithSorting(sorting string, number int, offset int, category string, descending int, search string) ([]types.Movie, error) {
	var order string
	if descending == 1 {
		order = "DESC"
	} else {
		order = "ASC"
	}

	var parameters []string

	categorySQL := ""
	if category != "any" {
		categorySQL = "INNER JOIN movieToGenre " +
			"ON movies.id = movieToGenre.movieId " +
			"AND movieToGenre.genreId in ( SELECT genre.genreId FROM genre WHERE genre.genreName = ? ) "
		parameters = append(parameters, category)
	}

	searchSQL := ""
	if search != "" {
		searchSQL = "WHERE title like ? "
		parameters = append(parameters, "%"+search+"%")
	}

	sorting =

	parameters = append(parameters, strconv.Itoa(offset))
	parameters = append(parameters, strconv.Itoa(number))

	movies := []types.Movie{}
	interfaces := make([]interface{}, len(parameters))
	for i, v := range parameters {
		interfaces[i] = v
	}

	error := connections.PostgreSQLQuery(
		&movies,
		"SELECT movies.id, movies.title, movies.posterURL, substring(description,1,100) as description "+
			"FROM movies "+
			categorySQL+
			searchSQL+
			"ORDER BY "+sorting+" "+order+" "+

			"LIMIT ?,? ",
		interfaces,
	)

	return movies, error
}

// getMovieByIDThirdParty = async (id)
