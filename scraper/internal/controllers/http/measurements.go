package http

import (
	"TPBDM/scraper/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controllers) measurementsEndpoints(router *gin.RouterGroup) *gin.RouterGroup {

	type response struct {
		Result any    `json:"result,omitempty"`
		Error  string `json:"error,omitempty"`
	}

	router.GET("", func(ctx *gin.Context) {

		var query entities.MeasurementsQueryContract
		err := ctx.BindQuery(&query)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response{Error: err.Error()})
			return
		}

		res, err := c.service.GetMeasurementsList(query)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response{Result: res})
	})

	router.GET("/cities", func(ctx *gin.Context) {

		var query entities.MeasurementsQueryContract
		_ = ctx.Bind(&query)

		res, err := c.service.GetMeasurementCitiesList(query)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response{Result: res})
	})

	router.GET("/locations", func(ctx *gin.Context) {

		var query entities.MeasurementsQueryContract
		_ = ctx.BindQuery(&query)

		res, err := c.service.GetMeasurementLocationsList(query)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response{Result: res})
	})

	router.GET("/parameters", func(ctx *gin.Context) {

		var query entities.MeasurementsQueryContract
		_ = ctx.BindQuery(&query)

		res, err := c.service.GetMeasurementParameterList(query)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response{Error: err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response{Result: res})
	})

	return router
}
