package openaq

import (
	"TPBDM/scraper/config"
	"TPBDM/scraper/internal/entities"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type openAQ struct {
	url string
}

// New ...
func New(config config.OpenAQConfig) (Repository, error) {
	return &openAQ{url: config.URL}, nil
}

func (o *openAQ) GetMeasurementsCount(query QueryContract) (int, error) {
	var result int
	var err error

	query.Limit = 1
	reqUrl := o.url + "/measurements" + o.processQueryParams(query)

	var resp *http.Response
	resp, err = http.Get(reqUrl)
	if err != nil {
		return result, err
	}

	if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
		return result, fmt.Errorf("ERROR: %s: GOT %s", reqUrl, resp.Status)
	} else {
		var contract ResponseContract
		err = json.NewDecoder(resp.Body).Decode(&contract)
		if err != nil {
			return result, err
		}

		result = contract.Meta.Found
	}

	return result, err
}

func (o *openAQ) GetMeasurements(query QueryContract) ([]entities.Measurement, error) {
	var result []entities.Measurement
	var err error

	for attempt := 0; attempt <= 100; attempt++ {
		reqUrl := o.url + "/measurements" + o.processQueryParams(query)

		var resp *http.Response
		resp, err = http.Get(reqUrl)
		if err != nil {
			return result, err
		}

		if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
			err = fmt.Errorf("ERROR: %s: GOT %s", reqUrl, resp.Status)
			if attempt < 100 {
				log.Error().Msg(err.Error() + "; RETRY IN 1m")
				time.Sleep(time.Minute)
			}
		} else {
			var contract ResponseContract
			err = json.NewDecoder(resp.Body).Decode(&contract)
			if err != nil {
				return result, err
			}

			result = buildMeasurementEntities(contract.Results)
			err = nil
			break
		}
	}

	return result, err
}

func (o *openAQ) processQueryParams(query QueryContract) (res string) {
	if query.Page != 0 {
		res += "&page=" + url.QueryEscape(strconv.Itoa(query.Page))
	}

	if query.Limit != 0 {
		res += "&limit=" + url.QueryEscape(strconv.Itoa(query.Limit))
	}

	if query.Country != "" {
		res += "&country=" + url.QueryEscape(query.Country)
	}

	if query.City != "" {
		res += "&city=" + url.QueryEscape(query.City)
	}

	if query.Parameter != "" {
		res += "&parameter=" + url.QueryEscape(query.Parameter)
	}

	if res[0] == '&' {
		res = res[1:]
	}

	if res != "" {
		res = "?" + res
	}

	return
}
