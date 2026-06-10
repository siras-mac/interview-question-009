package domain

import "time"

// Comment represents a user comment under a post.
type Comment struct {
	ID         int       `json:"id"`
	PostID     int       `json:"postId"`
	AuthorName string    `json:"authorName"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"createdAt"`
}

// CommentRepository defines the persistence contract for comments.
type CommentRepository interface {
	FindByPostID(postID int) ([]Comment, error)
	Create(comment *Comment) error
}
