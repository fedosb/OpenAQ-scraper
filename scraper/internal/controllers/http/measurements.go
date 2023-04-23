package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controllers) measurementsEndpoints(router *gin.RouterGroup) *gin.RouterGroup {

	type response struct {
		Result any   `json:"result,omitempty"`
		Error  error `json:"error,omitempty"`
	}

	router.GET("", func(ctx *gin.Context) {

		res, err := c.service.GetMeasurementsList()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response{Error: err})
			return
		}

		ctx.JSON(http.StatusOK, response{Result: res})
	})

	return router
}
