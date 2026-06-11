package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"example.com/interview-question-09/api/internal/domain"
	"example.com/interview-question-09/api/internal/handler"
	"example.com/interview-question-09/api/internal/repository/memory"
	"example.com/interview-question-09/api/internal/usecase"
)

func newTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	postRepo := memory.NewPostRepo(domain.Post{
		ID:         1,
		AuthorName: "Change can",
		ImageURL:   "assets/post.png",
		PostedAt:   time.Date(2021, time.October, 16, 16, 0, 0, 0, time.UTC),
	})
	commentRepo := memory.NewCommentRepo()

	postUC := usecase.NewPostUsecase(postRepo)
	commentUC := usecase.NewCommentUsecase(postRepo, commentRepo)
	return handler.NewRouter(handler.NewPostHandler(postUC, commentUC))
}

func TestGetPostEndpoint(t *testing.T) {
	router := newTestRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/posts/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	var post domain.Post
	if err := json.Unmarshal(w.Body.Bytes(), &post); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if post.AuthorName != "Change can" {
		t.Errorf("author = %q, want %q", post.AuthorName, "Change can")
	}
}

func TestGetPostEndpointErrors(t *testing.T) {
	tests := []struct {
		name string
		path string
		want int
	}{
		{name: "post not found", path: "/api/posts/99", want: http.StatusNotFound},
		{name: "invalid id", path: "/api/posts/abc", want: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := newTestRouter()

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			router.ServeHTTP(w, req)

			if w.Code != tt.want {
				t.Errorf("status = %d, want %d", w.Code, tt.want)
			}
		})
	}
}

func TestAddCommentEndpoint(t *testing.T) {
	tests := []struct {
		name string
		path string
		body string
		want int
	}{
		{name: "created", path: "/api/posts/1/comments", body: `{"authorName":"Blend 285","message":"have a good day"}`, want: http.StatusCreated},
		{name: "empty message", path: "/api/posts/1/comments", body: `{"authorName":"Blend 285","message":""}`, want: http.StatusBadRequest},
		{name: "invalid body", path: "/api/posts/1/comments", body: `not-json`, want: http.StatusBadRequest},
		{name: "post not found", path: "/api/posts/99/comments", body: `{"authorName":"Blend 285","message":"hi"}`, want: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := newTestRouter()

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, tt.path, strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			if w.Code != tt.want {
				t.Fatalf("status = %d, want %d", w.Code, tt.want)
			}
			if tt.want == http.StatusCreated {
				var comment domain.Comment
				if err := json.Unmarshal(w.Body.Bytes(), &comment); err != nil {
					t.Fatalf("invalid JSON response: %v", err)
				}
				if comment.Message != "have a good day" {
					t.Errorf("message = %q, want %q", comment.Message, "have a good day")
				}
			}
		})
	}
}

func TestListCommentsEndpoint(t *testing.T) {
	router := newTestRouter()

	// add one comment first
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/posts/1/comments", strings.NewReader(`{"authorName":"Blend 285","message":"hello"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/api/posts/1/comments", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
	var comments []domain.Comment
	if err := json.Unmarshal(w.Body.Bytes(), &comments); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if len(comments) != 1 {
		t.Errorf("len = %d, want 1", len(comments))
	}
}
