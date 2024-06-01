package routes

import (
	"express-to-gin/middleware"
	"express-to-gin/services"
	"express-to-gin/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitComments(server *gin.Engine) {
	/**
	 * Get first order comments that are in movie description
	 * @param movieId - integer, id for which to get comments for.
	 * @param number - int, how many comments to return
	 * @param offset - int, how many comments to skip by
	 *
	 * @example - GET {BaseURL}/comments/getFirstOrderComments/123456/1/0
	 */
	server.GET("/comments/getFirstOrderComments/:movieId/:number/:offset",
		middleware.ValidateInt("movieId", "required,gte=1,lte=9999999"),
		middleware.ValidateInt("number", "required,gte=1,lte=1000"),
		middleware.ValidateInt("offset", "gte=0,lte=9999999"), // Not required as 0 clashes with 'required'
		func(ctx *gin.Context) {
			// Retrieve the value set by the middleware
			movieId := ctx.GetInt("movieId") // Retrieve the value set by the middleware
			number := ctx.GetInt("number")
			offset := ctx.GetInt("offset")

			comments, error := services.GetCommentsFirstOrder(movieId, number, offset)

			if error != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Internal server error: %s", error.Error())})
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"comments": comments,
			})
		})

	/**
	 * Get comments that are replying to a comment that are in movie description
	 * @param movieId - integer, id for which to get comments for.
	 * @param number - int, how many comments to return
	 * @param offset - int, how many comments to skip by
	 *
	 * @example - GET {BaseURL}/comments/getSecondOrderComments/123456/123456/1/0
	 */
	server.GET("/comments/getSecondOrderComments/:movieId/:commentId/:number/:offset",
		middleware.ValidateInt("movieId", "required,gte=1,lte=9999999"),
		middleware.ValidateInt("commentId", "required,gte=1,lte=9999999"),
		middleware.ValidateInt("number", "required,gte=1,lte=1000"),
		middleware.ValidateInt("offset", "gte=0,lte=9999999"), // Not required as 0 clashes with 'required'
		func(ctx *gin.Context) {
			// Retrieve the value set by the middleware
			movieId := ctx.GetInt("movieId")
			commentId := ctx.GetInt("commentId")
			number := ctx.GetInt("number")
			offset := ctx.GetInt("offset")

			comments, error := services.GetCommentsSecondOrder(movieId, commentId, number, offset)

			if error != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Internal server error: %s", error.Error())})
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"comments": comments,
			})
		})

	/**
	 * Post comment for user
	 * @param userId - integer, id for which to post comments for.
	 *
	 * @example - POST {BaseURL}/comments
	 * @body -
	 * {
	 *     "movieId": 123,
	 *     "userId": 123,
	 *     "text": "TEXT FOR COMMENT HERE",
	 *     "replyCommentId": null,
	 *     "commentTime": null
	 * }
	 */
	server.POST("/comments",
		func(ctx *gin.Context) {
			var comment types.Comment
			if err := ctx.BindJSON(&comment); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return // Do not proceed if there's an error
			}

			error := services.PostComment(comment)

			if error != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Internal server error: %s", error.Error())})
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": "comment processed",
			})
		})
}
