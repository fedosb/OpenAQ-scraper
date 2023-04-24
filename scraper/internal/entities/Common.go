package entities

type ScrapingQueryContract struct {
	Parameter string `form:"parameter" binding:"required"`
	Count     int    `form:"count"`
	City      string `form:"city"`
	Country   string `form:"country,default=US"`
	Limit     int    `form:"limit,default=1000"`
}
