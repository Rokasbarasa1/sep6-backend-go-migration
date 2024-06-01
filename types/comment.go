package types

import "time"

type Comment struct {
	MovieId        uint32     `json:"movieId" binding:"required,gte=1,lte=9999999"`
	UserId         uint32     `json:"userId" binding:"required,gte=1"`
	Nickname       *string    `json:"nickname,omitempty"`
	PhotoURL       *string    `json:"photoURL,omitempty"`
	CommentText    string     `json:"text" binding:"required"`
	HasReplies     *bool      `json:"hasReplies,omitempty"`
	ReplyCommentId *uint32    `json:"replyCommentId,omitempty" binding:"omitempty,gte=1,lte=9999999"`
	CommentTime    *time.Time `json:"commentTime,omitempty"`
}
