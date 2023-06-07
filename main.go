package main

import (
	"os"

	"auto_block_snapshot/cmd"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	if err := cmd.Execute(); err != nil {
		log.Error().Err(err).Msg("An error occurred while executing the command")
		os.Exit(1)
	}
}
