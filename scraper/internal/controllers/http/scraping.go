package http

import (
	"TPBDM/scraper/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controllers) scrapingEndpoints(router *gin.RouterGroup) *gin.RouterGroup {

	type response struct {
		Error string `json:"error,omitempty"`
	}

	router.GET("", func(ctx *gin.Context) {

		var query entities.ScrapingQueryContract
		err := ctx.BindQuery(&query)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response{Error: err.Error()})
			return
		}

		err = c.service.BeginScraping(query)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response{Error: err.Error()})
			return
		}

		ctx.Status(http.StatusAccepted)
	})

	return router
}
