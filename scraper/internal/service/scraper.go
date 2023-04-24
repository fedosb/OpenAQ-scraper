package service

import (
	"TPBDM/scraper/internal/entities"
	"TPBDM/scraper/internal/repositories/openaq"
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/semaphore"
	"strconv"
	"sync"
	"time"
)

func (s *service) BeginScraping(contract entities.ScrapingQueryContract) error {
	locked := !s.scrapingMutex.TryLock()
	if locked {
		return fmt.Errorf("WAIT WHILE PROCESS IS RUNNING")
	}

	go s.scrape(contract)

	return nil
}

func (s *service) scrape(query entities.ScrapingQueryContract) {
	defer s.scrapingMutex.Unlock()

	var totalCount int
	var err error

	if query.Count == 0 {
		totalCount, err = s.r.OpenAQ.GetMeasurementsCount(openaq.QueryContract{
			Country:   query.Country,
			Parameter: query.Parameter,
			City:      query.City,
		})
		if err != nil {
			log.Error().Msg(err.Error())
			totalCount = query.Count
		}
	} else {
		totalCount = query.Count
	}

	log.Info().Msg("BEGIN SCRAPING")
	log.Info().Msg("TOTAL COUNT: " + strconv.Itoa(totalCount))

	var (
		sem = semaphore.NewWeighted(1)
		wg  = sync.WaitGroup{}
		ctx = context.TODO()
	)

	for page := 1; page <= totalCount/query.Limit+1; page++ {
		_ = sem.Acquire(ctx, 1)
		wg.Add(1)

		go func(page int) {
			defer sem.Release(1)
			defer wg.Done()

			res, err := s.r.OpenAQ.GetMeasurements(openaq.QueryContract{
				Page:      page,
				Limit:     query.Limit,
				Country:   query.Country,
				Parameter: query.Parameter,
				City:      query.City,
			})
			if err != nil {
				log.Error().Msg(err.Error())
				return
			}

			var cnt int
			for _, measurement := range res {
				err = s.r.DB.Measurements.CreateMeasurement(measurement)
				if err != nil {
					log.Error().Msg(err.Error())
				} else {
					cnt++
				}
			}
			log.Info().Msg(fmt.Sprintf("PAGE %d (%d / %d): WROTE %d MEASUREMENTS OF %d",
				page,
				totalCount,
				query.Limit,
				cnt,
				len(res),
			))
		}(page)

		time.Sleep(time.Second * 10)
	}

	wg.Wait()
	log.Info().Msg("SCRAPING DONE")
}
