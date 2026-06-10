package memory

import (
	"sync"

	"example.com/interview-question-009/api/internal/domain"
)

// PostRepo is an in-memory implementation of domain.PostRepository.
type PostRepo struct {
	mu    sync.RWMutex
	posts map[int]domain.Post
}

func NewPostRepo(seed ...domain.Post) *PostRepo {
	r := &PostRepo{posts: make(map[int]domain.Post)}
	for _, p := range seed {
		r.posts[p.ID] = p
	}
	return r
}

func (r *PostRepo) FindByID(id int) (*domain.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	post, ok := r.posts[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return &post, nil
}
