package app

import (
	"os"
	"os/signal"
	"syscall"

	"detfes/vars"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jenggo/fiberlog"
	"github.com/rs/zerolog/log"
)

func RunServer() error {
	config := fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		BodyLimit:             1024 * 1024 * 1024,
		DisableStartupMessage: true,
	}

	app := fiber.New(config)
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(fiberlog.New(fiberlog.Config{
		Next: func(ctx *fiber.Ctx) bool {
			return ctx.Path() == "/ping"
		},
	}))

	Router(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Info().Msg(vars.AppName + " v" + vars.Version + " - listen " + vars.Config.Listen)

	go func() {
		if err := app.Listen(vars.Config.Listen); err != nil {
			log.Panic().Err(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	return app.Shutdown()
}
