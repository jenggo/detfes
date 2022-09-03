package main

import (
	"os"

	"detfes/app"
	"detfes/vars"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}).With().Timestamp().Logger()

	if err := cleanenv.ReadConfig(vars.FileConfig, &vars.Config); err != nil {
		log.Fatal().Msg(err.Error())
	}

	if _, err := os.ReadDir(vars.Config.Path.Temp); err != nil {
		if err := os.Mkdir(vars.Config.Path.Temp, 0o744); err != nil {
			log.Fatal().Msg(err.Error())
		}
	}

	if err := app.RunServer(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
