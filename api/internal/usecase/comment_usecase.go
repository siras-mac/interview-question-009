package usecase

import (
	"strings"
	"time"

	"example.com/interview-question-009/api/internal/domain"
)

// CommentUsecase contains business logic for comments.
type CommentUsecase struct {
	posts    domain.PostRepository
	comments domain.CommentRepository
}

func NewCommentUsecase(posts domain.PostRepository, comments domain.CommentRepository) *CommentUsecase {
	return &CommentUsecase{posts: posts, comments: comments}
}

// ListComments returns all comments of a post, oldest first.
func (u *CommentUsecase) ListComments(postID int) ([]domain.Comment, error) {
	if _, err := u.posts.FindByID(postID); err != nil {
		return nil, err
	}
	return u.comments.FindByPostID(postID)
}

// AddComment validates and stores a new comment under a post.
func (u *CommentUsecase) AddComment(postID int, authorName, message string) (*domain.Comment, error) {
	authorName = strings.TrimSpace(authorName)
	message = strings.TrimSpace(message)
	if authorName == "" {
		return nil, domain.ErrEmptyAuthor
	}
	if message == "" {
		return nil, domain.ErrEmptyMessage
	}
	if _, err := u.posts.FindByID(postID); err != nil {
		return nil, err
	}

	comment := &domain.Comment{
		PostID:     postID,
		AuthorName: authorName,
		Message:    message,
		CreatedAt:  time.Now(),
	}
	if err := u.comments.Create(comment); err != nil {
		return nil, err
	}
	return comment, nil
}
