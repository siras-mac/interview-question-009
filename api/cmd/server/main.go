package main

import (
	"log"
	"os"
	"time"

	"example.com/interview-question-009/api/internal/domain"
	"example.com/interview-question-009/api/internal/handler"
	"example.com/interview-question-009/api/internal/repository/memory"
	"example.com/interview-question-009/api/internal/usecase"
)

func main() {
	bangkok := time.FixedZone("ICT", 7*60*60)

	postRepo := memory.NewPostRepo(domain.Post{
		ID:         1,
		AuthorName: "Change can",
		ImageURL:   "assets/post.png",
		PostedAt:   time.Date(2021, time.October, 16, 16, 0, 0, 0, bangkok),
	})
	commentRepo := memory.NewCommentRepo(domain.Comment{
		PostID:     1,
		AuthorName: "Blend 285",
		Message:    "have a good day",
		CreatedAt:  time.Date(2021, time.October, 16, 16, 5, 0, 0, bangkok),
	})

	postUC := usecase.NewPostUsecase(postRepo)
	commentUC := usecase.NewCommentUsecase(postRepo, commentRepo)
	router := handler.NewRouter(handler.NewPostHandler(postUC, commentUC))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("API server listening on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
