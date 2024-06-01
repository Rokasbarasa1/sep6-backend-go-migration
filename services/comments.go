package services

import (
	"express-to-gin/models"
	"express-to-gin/types"
)

func GetCommentsFirstOrder(movieId int, number int, offset int) ([]types.Comment, error) {
	return models.GetFirstOrderCommentsForMovie(movieId, number, offset)
}

func GetCommentsSecondOrder(movieId int, commentId int, number int, offset int) ([]types.Comment, error) {
	return models.GetSecondOrderCommentsForMovie(movieId, commentId, number, offset)
}

func PostComment(commentBody types.Comment) error {
	return models.PostComment(commentBody)
}
