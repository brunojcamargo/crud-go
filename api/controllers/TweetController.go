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

func (t *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
		})
		return
	}

	t.tweets = append(t.tweets, *tweet)

	ctx.JSON(http.StatusCreated, gin.H{
		"data":  tweet,
		"error": false,
	})
}

func (t *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for idx, tweet := range t.tweets {
		if tweet.ID == id {
			t.tweets = append(t.tweets[0:idx], t.tweets[idx+1:]...)
			ctx.JSON(http.StatusNoContent, gin.H{
				"data":  tweet,
				"error": false,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": true,
	})
}
