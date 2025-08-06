package handler

import (
	"errors"
	"fmt"
	"github.com/Kosench/go-shortener/internal/shortener/model"
	"github.com/Kosench/go-shortener/internal/shortener/repository"
	"github.com/Kosench/go-shortener/internal/shortener/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShortenerHandler struct {
	service *service.ShortenerService
}

func NewShortenerHandler(s *service.ShortenerService) *ShortenerHandler {
	return &ShortenerHandler{service: s}
}

func (h *ShortenerHandler) ShortenURL(c *gin.Context) {
	var req model.ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request: %v", err)})
		return
	}

	code, err := h.service.Shorten(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to shorten URL: %v", err)})
		return
	}

	resp := model.ShortenResponse{ShortURL: fmt.Sprintf("http://%s/%s", c.Request.Host, code)}
	c.JSON(http.StatusOK, resp)
}

func (h *ShortenerHandler) RedirectURL(c *gin.Context) {
	code := c.Param("code")
	urlModel, err := h.service.GetOriginalURL(code)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.Redirect(http.StatusMovedPermanently, urlModel.Original)
}
