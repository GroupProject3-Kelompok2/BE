package helper

import "github.com/aidarkhanov/nanoid"

func GenerateId() string {
	return nanoid.New()
}
