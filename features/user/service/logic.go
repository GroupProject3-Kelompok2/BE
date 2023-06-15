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

// ProfileUser implements user.UserService
func (us *userService) ProfileUser(userId string) (user.UserCore, error) {
	result, err := us.query.ProfileUser(userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("not found, error while retrieving user profile")
			return user.UserCore{}, errors.New("not found, error while retrieving user profile")
		} else {
			log.Error("internal server error")
			return user.UserCore{}, errors.New("internal server error")
		}
	}
	return result, nil
}

// UpdateProfile implements user.UserService
func (us *userService) UpdateProfile(userId string, request user.UserCore) error {
	err := us.query.UpdateProfile(userId, request)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("user profile record not found")
			return errors.New("user profile record not found")
		} else if strings.Contains(err.Error(), "duplicate data entry") {
			log.Error("failed to update user, duplicate data entry")
			return errors.New("failed to update user, duplicate data entry")
		} else {
			log.Error("internal server error")
			return errors.New("internal server error")
		}
	}

	return nil
}

// DeactiveUser implements user.UserService
func (us *userService) DeactiveUser(userId string) error {
	err := us.query.DeactiveUser(userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("user profile record not found")
			return errors.New("user profile record not found")
		} else {
			log.Error("internal server error")
			return errors.New("internal server error")
		}
	}

	return nil
}

// UpgradeProfile implements user.UserService
func (us *userService) UpgradeProfile(userId string, request user.UserCore) error {
	err := us.query.UpgradeProfile(userId, request)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("user profile record not found")
			return errors.New("user profile record not found")
		} else {
			log.Error("internal server error")
			return errors.New("internal server error")
		}
	}

	return nil
}

// MyHomestays implements user.UserService
func (*userService) MyHomestays(userId string) ([]user.UserCore, error) {
	panic("unimplemented")
}
