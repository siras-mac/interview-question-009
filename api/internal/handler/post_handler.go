package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"example.com/interview-question-009/api/internal/domain"
	"example.com/interview-question-009/api/internal/usecase"
)

// PostHandler exposes post and comment endpoints over HTTP.
type PostHandler struct {
	postUC    *usecase.PostUsecase
	commentUC *usecase.CommentUsecase
}

func NewPostHandler(postUC *usecase.PostUsecase, commentUC *usecase.CommentUsecase) *PostHandler {
	return &PostHandler{postUC: postUC, commentUC: commentUC}
}

type addCommentRequest struct {
	AuthorName string `json:"authorName"`
	Message    string `json:"message"`
}

// GetPost handles GET /api/posts/:id
func (h *PostHandler) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	post, err := h.postUC.GetPost(id)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, post)
}

// ListComments handles GET /api/posts/:id/comments
func (h *PostHandler) ListComments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	comments, err := h.commentUC.ListComments(id)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, comments)
}

// AddComment handles POST /api/posts/:id/comments
func (h *PostHandler) AddComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	var req addCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	comment, err := h.commentUC.AddComment(id, req.AuthorName, req.Message)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusCreated, comment)
}

func respondError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
	case errors.Is(err, domain.ErrEmptyMessage), errors.Is(err, domain.ErrEmptyAuthor):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
