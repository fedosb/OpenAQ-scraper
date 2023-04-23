package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controllers) scrapingEndpoints(router *gin.RouterGroup) *gin.RouterGroup {

	type response struct {
		Error string `json:"error,omitempty"`
	}

	router.GET("", func(ctx *gin.Context) {

		err := c.service.BeginScraping()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response{Error: err.Error()})
			return
		}

		ctx.Status(http.StatusAccepted)
	})

	return router
}
