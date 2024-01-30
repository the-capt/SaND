package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/the-capt/SaND/pkg/proto/sand"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("This is an Info message.")

	rr := &sand.AddressRange{Beginning: 0, Ending: 1}
	log.Info().Str("range", rr.String()).Msg("")
	if rr.Beginning == 0 {
		log.Info().Msg("beginning is 0")
	}
}
