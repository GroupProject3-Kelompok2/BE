package service

import (
	"errors"
	"strings"

	"github.com/GroupProject3-Kelompok2/BE/features/user"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"github.com/GroupProject3-Kelompok2/BE/utils/validation"
)

var log = middlewares.Log()

type userService struct {
	query user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userService{
		query: ud,
	}
}

// Register implements user.UserService
func (us *userService) Register(request user.UserCore) (user.UserCore, error) {
	if request.Fullname == "" || request.Email == "" || request.Password == "" {
		log.Error("request cannot be empty")
		return user.UserCore{}, errors.New("request cannot be empty")
	}

	_, err := validation.UserValidate("register", request)
	if err != nil {
		if strings.Contains(err.Error(), "email") {
			return user.UserCore{}, errors.New("invalid email format")
		}
		return user.UserCore{}, errors.New("check password strength, low password")
	}

	result, err := us.query.Register(request)
	if err != nil {
		message := ""
		if strings.Contains(err.Error(), "error while hashing password") {
			log.Error("error while hashing password")
			message = "error while hashing password"
		} else if strings.Contains(err.Error(), "error insert data, duplicated") {
			log.Error("error insert data, duplicated")
			message = "error insert data, duplicated"
		} else {
			log.Error("internal server error")
			message = "internal server error"
		}
		return user.UserCore{}, errors.New(message)
	}

	return result, nil
}
