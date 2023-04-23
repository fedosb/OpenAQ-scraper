package openaq

import (
	"TPBDM/scraper/config"
	"TPBDM/scraper/internal/entities"
	"encoding/json"
	"net/http"
)

type openAQ struct {
	url string
}

// New ...
func New(config config.OpenAQConfig) (Repository, error) {
	return &openAQ{url: config.URL}, nil
}

func (o *openAQ) GetMeasurements() ([]entities.Measurement, error) {
	var result []entities.Measurement

	reqUrl := o.url + "/measurements"

	resp, err := http.Get(reqUrl)
	if err != nil {
		return result, err
	}

	var contract ResponseContract
	err = json.NewDecoder(resp.Body).Decode(&contract)
	if err != nil {
		return result, err
	}

	result = buildMeasurementEntities(contract.Results)

	return result, nil
}

func (o *openAQ) GetMeasurementsCount() (int, error) {
	var result int

	reqUrl := o.url + "/measurements?limit=1"

	resp, err := http.Get(reqUrl)
	if err != nil {
		return result, err
	}

	var contract ResponseContract
	err = json.NewDecoder(resp.Body).Decode(&contract)
	if err != nil {
		return result, err
	}

	result = contract.Meta.Found

	return result, nil
}
