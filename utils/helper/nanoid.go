package helper

import gonanoid "github.com/matoous/go-nanoid/v2"

func GenerateId() (string, error) {
	return gonanoid.New()
}
