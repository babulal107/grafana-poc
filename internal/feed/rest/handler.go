package rest

import (
	"github.com/gin-gonic/gin"
	feedModel "github.com/grafana-poc/internal/feed/model"
	"github.com/grafana-poc/internal/model"
	"net/http"
)

type handler struct {
}

// NewHandler represents implementation of handler to be separated with Gin
func NewHandler() handler {
	return handler{}
}

func (h handler) HandleHTTP(c *gin.Context) {
	var feeds []feedModel.FeedCard
	feed := feedModel.FeedCard{
		CardHash: "342SDFASF324324234",
		Title:    "Test static feed title",
		Body:     "this is test static feed body",
	}
	feeds = append(feeds, feed)
	c.JSON(http.StatusOK, model.SuccessResponse(feeds))
}
