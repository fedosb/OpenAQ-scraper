package main

import (
	"TPBDM/scraper/config"
	"TPBDM/scraper/internal/controllers/http"
	"TPBDM/scraper/internal/repositories"
	"TPBDM/scraper/internal/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var c *config.Config

func init() {
	// log settings init
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// config init
	c = config.New()
}

func main() {

	r, err := repositories.New(*c)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	controllers, err := http.New(
		service.New(*r),
	)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	log.Info().Msg("Server is going to start")

	err = controllers.Router.Run()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
