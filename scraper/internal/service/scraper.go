package service

import (
	"TPBDM/scraper/internal/repositories/openaq"
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/semaphore"
	"strconv"
	"sync"
	"time"
)

func (s *service) BeginScraping() error {
	locked := !s.scrapingMutex.TryLock()
	if locked {
		return fmt.Errorf("WAIT WHILE PROCESS IS RUNNING")
	}

	go s.scrape()

	return nil
}

func (s *service) scrape() {
	defer s.scrapingMutex.Unlock()

	var limit = 100

	log.Info().Msg("BEGIN SCRAPING")

	totalCount, err := s.r.OpenAQ.GetMeasurementsCount(openaq.QueryContract{
		Page:    1,
		Country: "US",
	})
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	log.Info().Msg("TOTAL COUNT: " + strconv.Itoa(totalCount))

	var (
		sem = semaphore.NewWeighted(1)
		wg  = sync.WaitGroup{}
		ctx = context.TODO()
	)

	for page := 1; page <= totalCount/limit+1; page++ {
		_ = sem.Acquire(ctx, 1)
		wg.Add(1)

		go func(page int) {
			defer sem.Release(1)
			defer wg.Done()

			res, err := s.r.OpenAQ.GetMeasurements(openaq.QueryContract{
				Page:    page,
				Limit:   limit,
				Country: "US",
			})
			if err != nil {
				log.Error().Msg(err.Error())
				return
			}

			for _, measurement := range res {
				err = s.r.DB.Measurements.CreateMeasurement(measurement)
				if err != nil {
					log.Error().Msg(err.Error())
				}
			}
			log.Info().Msg("WROTE " + strconv.Itoa(len(res)) + " MEASUREMENTS")
		}(page)

		time.Sleep(time.Second + time.Millisecond*100)
	}

	wg.Wait()
	log.Info().Msg("SCRAPING DONE")
}
