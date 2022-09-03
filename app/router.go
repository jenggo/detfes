package app

import (
	"detfes/app/routes"
	"detfes/pkg"
	"detfes/vars"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/keyauth/v2"
	"github.com/rs/zerolog/log"
)

func Router(app *fiber.App) {
	app.Use(keyauth.New(keyauth.Config{
		Validator: func(c *fiber.Ctx, apikey string) (bool, error) {
			// Bypass checking apikey (bcrypt hashed)
			if vars.Config.Key.API == apikey {
				return true, nil
			}

			// Bypassed
			if pkg.ValidAPIKey(apikey, vars.Config.Key.API) {
				return true, nil
			}

			if vars.Config.Verbose {
				log.Error().Msg(keyauth.ErrMissingOrMalformedAPIKey.Error() + ": " + apikey)
			}

			return false, keyauth.ErrMissingOrMalformedAPIKey
		},
	}))

	app.Post("/identify", routes.Identify)
}
