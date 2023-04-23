package openaq

import "time"

type QueryContract struct {
	Page    int
	Limit   int
	Country string
}

type ResponseContract struct {
	Meta    MeasurementsMetaContract `json:"meta"`
	Results []MeasurementContract    `json:"results"`
}

type MeasurementsMetaContract struct {
	Name    string `json:"name"`
	License string `json:"license"`
	Website string `json:"website"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	Found   int    `json:"found"`
}

type MeasurementContract struct {
	Location  string                  `json:"location"`
	Parameter string                  `json:"parameter"`
	Value     float32                 `json:"value"`
	Date      MeasurementDateContract `json:"date"`
	Unit      string                  `json:"unit"`
	Country   string                  `json:"country"`
	City      string                  `json:"city"`
}

type MeasurementDateContract struct {
	Utc   time.Time `json:"utc"`
	Local time.Time `json:"local"`
}
