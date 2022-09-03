package routes

import (
	"os"
	"strings"

	"detfes/pkg"
	"detfes/vars"

	"github.com/gofiber/fiber/v2"
)

func Identify(c *fiber.Ctx) error {
	file, err := c.FormFile(vars.Config.Detect.FormName)
	if err != nil {
		return Return(c, fiber.StatusInternalServerError, err.Error(), 0, true)
	}

	if file.Header["Content-Type"][0] != "image/jpeg" {
		return Return(c, fiber.StatusBadRequest, "Only accept jpeg format", 0, true)
	}

	saveAs := strings.Join([]string{vars.Config.Path.Temp, file.Filename}, "/")

	if err := c.SaveFile(file, saveAs); err != nil {
		return Return(c, fiber.StatusInternalServerError, err.Error(), 0, true)
	}

	faces := pkg.FaceDetection(saveAs)
	_ = os.Remove(saveAs)

	if faces == 0 {
		return Return(c, fiber.StatusOK, "No face detected", 0, true)
	}

	giveError := false
	if !vars.Config.Detect.Multiple && faces > 1 {
		giveError = true
	}

	return Return(c, fiber.StatusOK, "Face detected", faces, giveError)
}
