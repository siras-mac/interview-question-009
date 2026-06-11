package usecase_test

import (
	"errors"
	"testing"
	"time"

	"example.com/interview-question-09/api/internal/domain"
	"example.com/interview-question-09/api/internal/repository/memory"
	"example.com/interview-question-09/api/internal/usecase"
)

func newFixture() (*usecase.PostUsecase, *usecase.CommentUsecase) {
	postRepo := memory.NewPostRepo(domain.Post{
		ID:         1,
		AuthorName: "Change can",
		ImageURL:   "assets/post.png",
		PostedAt:   time.Date(2021, time.October, 16, 16, 0, 0, 0, time.UTC),
	})
	commentRepo := memory.NewCommentRepo(domain.Comment{
		PostID:     1,
		AuthorName: "Blend 285",
		Message:    "have a good day",
		CreatedAt:  time.Date(2021, time.October, 16, 16, 5, 0, 0, time.UTC),
	})
	return usecase.NewPostUsecase(postRepo), usecase.NewCommentUsecase(postRepo, commentRepo)
}

func TestGetPost(t *testing.T) {
	postUC, _ := newFixture()

	post, err := postUC.GetPost(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if post.AuthorName != "Change can" {
		t.Errorf("author = %q, want %q", post.AuthorName, "Change can")
	}
}

func TestGetPostNotFound(t *testing.T) {
	postUC, _ := newFixture()

	if _, err := postUC.GetPost(99); !errors.Is(err, domain.ErrNotFound) {
		t.Errorf("error = %v, want ErrNotFound", err)
	}
}

func TestListComments(t *testing.T) {
	_, commentUC := newFixture()

	comments, err := commentUC.ListComments(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(comments) != 1 {
		t.Fatalf("len = %d, want 1", len(comments))
	}
	if comments[0].Message != "have a good day" {
		t.Errorf("message = %q, want %q", comments[0].Message, "have a good day")
	}
}

func TestListCommentsPostNotFound(t *testing.T) {
	_, commentUC := newFixture()

	if _, err := commentUC.ListComments(99); !errors.Is(err, domain.ErrNotFound) {
		t.Errorf("error = %v, want ErrNotFound", err)
	}
}

func TestAddComment(t *testing.T) {
	tests := []struct {
		name       string
		postID     int
		authorName string
		message    string
		wantErr    error
	}{
		{name: "valid comment", postID: 1, authorName: "Blend 285", message: "nice photo"},
		{name: "trims whitespace", postID: 1, authorName: "  Blend 285  ", message: "  hello  "},
		{name: "empty message", postID: 1, authorName: "Blend 285", message: "   ", wantErr: domain.ErrEmptyMessage},
		{name: "empty author", postID: 1, authorName: "", message: "hi", wantErr: domain.ErrEmptyAuthor},
		{name: "post not found", postID: 99, authorName: "Blend 285", message: "hi", wantErr: domain.ErrNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, commentUC := newFixture()

			comment, err := commentUC.AddComment(tt.postID, tt.authorName, tt.message)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("error = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if comment.ID == 0 {
				t.Error("comment ID was not assigned")
			}
			if comment.AuthorName != "Blend 285" {
				t.Errorf("author = %q, want trimmed %q", comment.AuthorName, "Blend 285")
			}

			comments, _ := commentUC.ListComments(tt.postID)
			if len(comments) != 2 {
				t.Errorf("comment count = %d, want 2", len(comments))
			}
		})
	}
}
