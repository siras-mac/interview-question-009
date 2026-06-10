package memory

import (
	"sync"

	"example.com/interview-question-009/api/internal/domain"
)

// CommentRepo is an in-memory implementation of domain.CommentRepository.
type CommentRepo struct {
	mu       sync.RWMutex
	comments []domain.Comment
	nextID   int
}

func NewCommentRepo(seed ...domain.Comment) *CommentRepo {
	r := &CommentRepo{nextID: 1}
	for _, c := range seed {
		c.ID = r.nextID
		r.nextID++
		r.comments = append(r.comments, c)
	}
	return r
}

func (r *CommentRepo) FindByPostID(postID int) ([]domain.Comment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]domain.Comment, 0)
	for _, c := range r.comments {
		if c.PostID == postID {
			result = append(result, c)
		}
	}
	return result, nil
}

func (r *CommentRepo) Create(comment *domain.Comment) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	comment.ID = r.nextID
	r.nextID++
	r.comments = append(r.comments, *comment)
	return nil
}
