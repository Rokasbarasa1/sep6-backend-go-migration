package routes

import (
	"express-to-gin/middleware"
	"express-to-gin/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitMovies(server *gin.Engine) {

	/**
	 * Get list of movies
	 * @param sorting - string, what parameter in movie object to sort by
	 * @param number - int, how many movies to return
	 * @param offset - int, how many movies to skip by
	 * @param category - string, how many movies to return
	 * @param descending - 1 or 0, 1 - sort and show descending, 0 - sort and show ascending
	 *
	 * @example - GET {BaseURL}/movies/list/title/10/0/Drama/1
	 */
	server.GET("/movies/list/:sorting/:number/:offset/:category/:descending",
		middleware.ValidateString("sorting", "required,len=3,len=20"),
		middleware.ValidateInt("number", "required,gte=1,lte=1000"),
		middleware.ValidateInt("offset", "gte=0,lte=9999999"),
		middleware.ValidateString("category", "required,len=3,len=20"),
		middleware.ValidateInt("descending", "gte=0,lte=1"),
		func(ctx *gin.Context) {
			// Retrieve the value set by the middleware
			sorting := ctx.GetString("sorting")
			number := ctx.GetInt("number") // Retrieve the value set by the middleware
			offset := ctx.GetInt("offset")
			category := ctx.GetString("category")
			descending := ctx.GetInt("descending")

			comments, error := services.GetCommentsFirstOrder(movieId, number, offset)

			if error != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Internal server error: %s", error.Error())})
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"movies": movies,
			})
		})
}
