package service

import (
	"errors"
	"strings"

	"github.com/GroupProject3-Kelompok2/BE/features/user"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"github.com/go-playground/validator/v10"
)

var log = middlewares.Log()

type userService struct {
	query    user.UserData
	validate *validator.Validate
}

func New(ud user.UserData, v *validator.Validate) user.UserService {
	return &userService{
		query:    ud,
		validate: v,
	}
}

// Register implements user.UserService
func (us *userService) Register(request user.UserCore) (user.UserCore, error) {
	err := us.validate.Struct(request)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "Fullname"):
			log.Warn("fullname cannot be empty")
			return user.UserCore{}, errors.New("fullname cannot be empty")
		case strings.Contains(err.Error(), "Email"):
			log.Warn("invalid email format")
			return user.UserCore{}, errors.New("invalid email format")
		case strings.Contains(err.Error(), "Phone"):
			log.Warn("phone cannot be empty")
			return user.UserCore{}, errors.New("phone cannot be empty")
		case strings.Contains(err.Error(), "Password"):
			log.Warn("password cannot be empty")
			return user.UserCore{}, errors.New("password cannot be empty")
		}
	}

	result, err := us.query.Register(request)
	if err != nil {
		message := ""
		switch {
		case strings.Contains(err.Error(), "error while hashing password"):
			log.Error("error while hashing password")
			message = "error while hashing password"
		case strings.Contains(err.Error(), "error insert data, duplicated"):
			log.Error("error insert data, duplicated")
			message = "error insert data, duplicated"
		default:
			log.Error("internal server error")
			message = "internal server error"
		}
		log.Error("request cannot be empty")
		return user.UserCore{}, errors.New(message)
	}

	log.Sugar().Infof("new user has been created: %s", result.UserID)
	return result, nil
}

// Login implements user.UserService
func (us *userService) Login(request user.UserCore) (user.UserCore, string, error) {
	err := us.validate.Struct(request)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "Email"):
			log.Warn("invalid email format")
			return user.UserCore{}, "", errors.New("invalid email format")
		case strings.Contains(err.Error(), "Password"):
			log.Warn("password cannot be empty")
			return user.UserCore{}, "", errors.New("password cannot be empty")
		}
	}

	result, token, err := us.query.Login(request)
	if err != nil {
		message := ""
		switch {
		case strings.Contains(err.Error(), "invalid email and password"):
			log.Error("invalid email and password")
			message = "invalid email and password"
		case strings.Contains(err.Error(), "password does not match"):
			log.Error("password does not match")
			message = "password does not match"
		case strings.Contains(err.Error(), "error while creating jwt token"):
			log.Error("error while creating jwt token")
			message = "error while creating jwt token"
		default:
			log.Error("internal server error")
			message = "internal server error"
		}
		return user.UserCore{}, "", errors.New(message)
	}

	log.Sugar().Infof("user has been logged in: %s", result.UserID)
	return result, token, nil
}
