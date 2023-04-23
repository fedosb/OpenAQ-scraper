package openaq

import (
	"TPBDM/scraper/config"
	"TPBDM/scraper/internal/entities"
	"encoding/json"
	"fmt"
	"net/http"
)

type openAQ struct {
	url string
}

// New ...
func New(config config.OpenAQConfig) (Repository, error) {
	return &openAQ{url: config.URL}, nil
}

func (o *openAQ) GetMeasurements(query QueryContract) ([]entities.Measurement, error) {
	var result []entities.Measurement

	reqUrl := o.url + fmt.Sprintf("/measurements?page=%d&limit=%d&country=%s", query.Page, query.Limit, query.Country)
	resp, err := http.Get(reqUrl)
	if err != nil {
		return result, err
	}

	if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
		return result, fmt.Errorf("ERROR: %s: GOT %s", reqUrl, resp.Status)
	}

	var contract ResponseContract
	err = json.NewDecoder(resp.Body).Decode(&contract)
	if err != nil {
		return result, err
	}

	result = buildMeasurementEntities(contract.Results)

	return result, nil
}

func (o *openAQ) GetMeasurementsCount(query QueryContract) (int, error) {
	var result int

	reqUrl := o.url + fmt.Sprintf("/measurements?page=%d&limit=1&country=%s", query.Page, query.Country)

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
