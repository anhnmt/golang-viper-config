package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	NewConfig()

	log.Info().
		Str("name", viper.GetString("app.name")).
		Msg("Hello, world!")
}
