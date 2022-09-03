package pkg

import (
	"detfes/vars"

	"github.com/Kagami/go-face"
	"github.com/rs/zerolog/log"
)

func FaceDetection(image string) int {
	rec, err := face.NewRecognizer(vars.Config.Path.Models)
	if err != nil {
		if vars.Config.Verbose {
			log.Error().Msg(err.Error())
		}
		return 0
	}
	defer rec.Close()

	faces, err := rec.RecognizeFile(image)
	if err != nil {
		if vars.Config.Verbose {
			log.Error().Msg(err.Error())
		}
		return 0
	}

	return len(faces)
}
