package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/the-capt/SaND/pkg/proto/sand"
	"google.golang.org/protobuf/proto"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	update := sand.SandUpdate{
		ProviderName: "/Users/mikeb/some_file.bin",
		Selections: []*sand.MultiSelection{
			{AddressSpaceName: "RAM",
				AddressSpaceExtent: &sand.AddressRange{Beginning: 0x0100_1000, Ending: 0x0100_1fff},
				IsSynthetic:        false,
				SelectionRanges: []*sand.AddressRange{
					{Beginning: 0x0100_0010, Ending: 0x0100_0020},
					{Beginning: 0x0100_0040, Ending: 0x0100_01C0},
				},
			}},
		CurrentLocation: &sand.Location{
			AddressSpaceName: "RAM", Address: 0x0100_0000,
		},
	}

	log.Info().Interface("update", &update).Send()

	encoded, err := proto.Marshal(&update)
	if err != nil {
		log.Error().Err(err)
	}

	log.Info().Hex("encoded", encoded).Send()
}
