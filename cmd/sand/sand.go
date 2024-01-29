package main

import (
	"os"

	pb "github.com/the-capt/SaND/pkg/proto/sand"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("This is an Info message.")

	rr := pb.Range{Beginning: 0, Ending: 1}
	log.Info.Str("range", rr).Msg("")
}
