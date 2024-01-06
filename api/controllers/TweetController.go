package controllers

import (
	"crud/api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	if len(t.tweets) < 1 {
		ctx.JSON(http.StatusNoContent, t.tweets)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  t.tweets,
		"error": false,
	})
}
