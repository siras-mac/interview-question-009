package memory_test

import (
	"fmt"
	"sync"
	"testing"

	"example.com/interview-question-009/api/internal/domain"
	"example.com/interview-question-009/api/internal/repository/memory"
)

func TestCommentRepoCreateAssignsIncrementingIDs(t *testing.T) {
	repo := memory.NewCommentRepo()

	first := &domain.Comment{PostID: 1, AuthorName: "Blend 285", Message: "first"}
	second := &domain.Comment{PostID: 1, AuthorName: "Blend 285", Message: "second"}
	if err := repo.Create(first); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := repo.Create(second); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if first.ID != 1 || second.ID != 2 {
		t.Errorf("ids = %d, %d, want 1, 2", first.ID, second.ID)
	}
}

func TestCommentRepoFindByPostIDFiltersAndKeepsOrder(t *testing.T) {
	repo := memory.NewCommentRepo(
		domain.Comment{PostID: 1, AuthorName: "Blend 285", Message: "first"},
		domain.Comment{PostID: 2, AuthorName: "Blend 285", Message: "other post"},
		domain.Comment{PostID: 1, AuthorName: "Blend 285", Message: "second"},
	)

	comments, err := repo.FindByPostID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(comments) != 2 {
		t.Fatalf("len = %d, want 2", len(comments))
	}
	if comments[0].Message != "first" || comments[1].Message != "second" {
		t.Errorf("wrong order: %q, %q", comments[0].Message, comments[1].Message)
	}
}

func TestCommentRepoFindByPostIDEmpty(t *testing.T) {
	repo := memory.NewCommentRepo()

	comments, err := repo.FindByPostID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if comments == nil {
		t.Error("expected empty slice, got nil")
	}
	if len(comments) != 0 {
		t.Errorf("len = %d, want 0", len(comments))
	}
}

func TestCommentRepoConcurrentCreate(t *testing.T) {
	repo := memory.NewCommentRepo()
	const n = 100

	var wg sync.WaitGroup
	for i := range n {
		wg.Go(func() {
			_ = repo.Create(&domain.Comment{PostID: 1, AuthorName: "Blend 285", Message: fmt.Sprintf("msg %d", i)})
		})
	}
	wg.Wait()

	comments, _ := repo.FindByPostID(1)
	if len(comments) != n {
		t.Fatalf("len = %d, want %d", len(comments), n)
	}

	seen := make(map[int]bool, n)
	for _, c := range comments {
		if seen[c.ID] {
			t.Fatalf("duplicate id %d", c.ID)
		}
		seen[c.ID] = true
	}
}
