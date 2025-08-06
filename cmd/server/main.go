package main

import (
	"github.com/Kosench/go-shortener/internal/shortener/handler"
	"github.com/Kosench/go-shortener/internal/shortener/repository"
	"github.com/Kosench/go-shortener/internal/shortener/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	storage := repository.NewMemoryRepository()
	shortener := service.NewShortenerService(storage)
	httpHandler := handler.NewShortenerHandler(shortener)

	r.POST("/shorten", httpHandler.ShortenURL)
	r.GET("/:code", httpHandler.RedirectURL)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
