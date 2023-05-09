package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// NewConfig initializes the config
func NewConfig() {
	viper.AutomaticEnv()

	// Replace env key
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("Error get current directory")
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath(fmt.Sprintf("%s/config", pwd))

	env := getDefaultEnv()
	viper.SetConfigFile(fmt.Sprintf("%s/config/%s.yaml", pwd, env))
	viper.SetConfigType("yaml")
	viper.SetConfigName(env)

	if err = viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Error reading config file")
	}

	log.Info().
		Str("env", env).
		Str("goarch", runtime.GOARCH).
		Str("goos", runtime.GOOS).
		Str("version", runtime.Version()).
		Msg("Runtime information")
}

// getDefaultEnv
func getDefaultEnv() string {
	env := strings.ToLower(os.Getenv("env"))
	switch env {
	case "prod":
		return "prod"
	case "dev":
		return "dev"
	}
	return "local"
}
