package entities

type ScrapingQueryContract struct {
	Parameter string `form:"parameter"`
	Count     int    `form:"count"`
	City      string `form:"city"`
	Location  string `form:"location"`
	Country   string `form:"country,default=US"`
	Limit     int    `form:"limit,default=1000"`
}

type MeasurementsQueryContract struct {
	Parameter string `form:"parameter"`
	City      string `form:"city"`
	Country   string `form:"country,default=US"`
	Location  string `form:"location"`
}

type List[T Measurement | string] struct {
	Data []T `json:"data"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}
