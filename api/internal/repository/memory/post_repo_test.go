package memory_test

import (
	"errors"
	"testing"
	"time"

	"example.com/interview-question-09/api/internal/domain"
	"example.com/interview-question-09/api/internal/repository/memory"
)

func TestPostRepoFindByID(t *testing.T) {
	repo := memory.NewPostRepo(domain.Post{
		ID:         1,
		AuthorName: "Change can",
		ImageURL:   "assets/post.png",
		PostedAt:   time.Date(2021, time.October, 16, 16, 0, 0, 0, time.UTC),
	})

	post, err := repo.FindByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if post.AuthorName != "Change can" {
		t.Errorf("author = %q, want %q", post.AuthorName, "Change can")
	}
}

func TestPostRepoFindByIDNotFound(t *testing.T) {
	repo := memory.NewPostRepo()

	if _, err := repo.FindByID(1); !errors.Is(err, domain.ErrNotFound) {
		t.Errorf("error = %v, want ErrNotFound", err)
	}
}

func TestPostRepoReturnsCopy(t *testing.T) {
	repo := memory.NewPostRepo(domain.Post{ID: 1, AuthorName: "Change can"})

	first, _ := repo.FindByID(1)
	first.AuthorName = "mutated"

	second, _ := repo.FindByID(1)
	if second.AuthorName != "Change can" {
		t.Errorf("stored post was mutated through returned pointer")
	}
}
