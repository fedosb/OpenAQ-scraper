package http

import (
	"TPBDM/scraper/internal/service"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	Router  *gin.Engine
	service service.Service
}

func New(
	service service.Service,
) (*Controllers, error) {
	r := gin.Default()

	c := &Controllers{
		Router:  r,
		service: service,
	}

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			c.measurementsEndpoints(v1.Group("/measurements"))
			c.scrapingEndpoints(v1.Group("/scraping"))
		}
	}

	return c, nil
}
