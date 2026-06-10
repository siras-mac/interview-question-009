package domain

import "time"

// Post represents a feed post that users can comment on.
type Post struct {
	ID         int       `json:"id"`
	AuthorName string    `json:"authorName"`
	ImageURL   string    `json:"imageUrl"`
	PostedAt   time.Time `json:"postedAt"`
}

// PostRepository defines the persistence contract for posts.
type PostRepository interface {
	FindByID(id int) (*Post, error)
}
