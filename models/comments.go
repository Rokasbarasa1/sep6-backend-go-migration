package models

import (
	"express-to-gin/connections"
	"express-to-gin/types"
)

func GetFirstOrderCommentsForMovie(movieId int, number int, offset int) ([]types.Comment, error) {
	comments := []types.Comment{}

	error := connections.PostgreSQLQuery(
		&comments,
		"SELECT *, appUser.nickname, appUser.photoURL, IF((SELECT COUNT(*) FROM movieComment as sc WHERE replyCommentId = fc.commentId), true, false) as hasReplies "+
			"FROM movieComment as fc "+
			"INNER JOIN appUser ON fc.userId = appUser.userId "+
			"WHERE fc.movieId = $1 "+
			"ORDER BY commentTime DESC LIMIT $2,$3",
		[]interface{}{movieId, number, offset},
	)

	return comments, error
}

func GetSecondOrderCommentsForMovie(movieId int, replyCommentId int, number int, offset int) ([]types.Comment, error) {
	comments := []types.Comment{}

	error := connections.PostgreSQLQuery(
		&comments,
		"SELECT *, appUser.nickname, appUser.photoURL "+
			"FROM movieComment "+
			"INNER JOIN appUser ON fc.userId = appUser.userId "+
			"WHERE movieId = $1 AND replyCommentId = $2 "+
			"LIMIT $3,$4 ",
		[]interface{}{movieId, replyCommentId, number, offset},
	)

	return comments, error
}

func PostComment(commentBody types.Comment) error {
	comments := []types.Comment{}

	error := connections.PostgreSQLQuery(
		&comments,
		"INSERT INTO movieComment (movieId, userId, commentText, replyCommentId, commentTime) VALUES ($1, $2, $3, $4, NOW())",
		[]interface{}{commentBody.MovieId, commentBody.UserId, commentBody.CommentText, commentBody.ReplyCommentId},
	)

	return error
}
