package usecase

import "example.com/interview-question-09/api/internal/domain"

// PostUsecase contains business logic for posts.
type PostUsecase struct {
	posts domain.PostRepository
}

func NewPostUsecase(posts domain.PostRepository) *PostUsecase {
	return &PostUsecase{posts: posts}
}

// GetPost returns a single post by its ID.
func (u *PostUsecase) GetPost(id int) (*domain.Post, error) {
	return u.posts.FindByID(id)
}
