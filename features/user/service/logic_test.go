package service

import (
	"errors"
	"testing"

	"github.com/GroupProject3-Kelompok2/BE/features/user"
	"github.com/GroupProject3-Kelompok2/BE/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	data := mocks.NewUserData(t)
	validate := validator.New()
	arguments := user.UserCore{
		Fullname: "admin",
		Email:    "admin@gmail.com",
		Phone:    "081235288543",
		Password: "@S3#cr3tP4ss#word123",
		Role:     "user",
	}
	result := user.UserCore{
		UserID:   "550e8400-e29b-41d4-a716-446655440000",
		Fullname: "admin",
		Email:    "admin@gmail.com",
		Phone:    "081235288543",
		Password: "@S3#cr3tP4ss#word123",
		Role:     "user",
	}
	service := New(data, validate)

	t.Run("fullname cannot be empty", func(t *testing.T) {
		request := user.UserCore{
			Fullname: "",
			Email:    "admin@gmail.com",
			Phone:    "081235288543",
			Password: "@S3#cr3tP4ss#word123",
			Role:     "user",
		}
		_, err := service.Register(request)
		expectedErr := errors.New("fullname cannot be empty")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("invalid email format", func(t *testing.T) {
		request := user.UserCore{
			Fullname: "admin",
			Email:    "admin@.com",
			Phone:    "081235288543",
			Password: "@S3#cr3tP4ss#word123",
			Role:     "user",
		}
		_, err := service.Register(request)
		expectedErr := errors.New("invalid email format")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("phone cannot be empty", func(t *testing.T) {
		request := user.UserCore{
			Fullname: "admin",
			Email:    "admin@gmail.com",
			Phone:    "",
			Password: "@S3#cr3tP4ss#word123",
			Role:     "user",
		}
		_, err := service.Register(request)
		expectedErr := errors.New("phone cannot be empty")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("password cannot be empty", func(t *testing.T) {
		request := user.UserCore{
			Fullname: "admin",
			Email:    "admin@gmail.com",
			Phone:    "081235288543",
			Password: "",
			Role:     "user",
		}
		_, err := service.Register(request)
		expectedErr := errors.New("password cannot be empty")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("validation error", func(t *testing.T) {
		request := user.UserCore{
			Fullname: "",
			Email:    "",
			Phone:    "",
			Password: "",
			Role:     "",
		}
		_, err := service.Register(request)
		expectedErr := errors.New("validation error")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("success create account", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(result, nil).Once()
		res, err := service.Register(arguments)
		assert.Nil(t, err)
		assert.Equal(t, result.UserID, res.UserID)
		assert.NotEmpty(t, result.Fullname)
		assert.NotEmpty(t, result.Email)
		assert.NotEmpty(t, result.Password)
		data.AssertExpectations(t)
	})

	t.Run("error while hashing password", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.UserCore{}, errors.New("error while hashing password")).Once()
		res, err := service.Register(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.UserID)
		assert.ErrorContains(t, err, "password")
		data.AssertExpectations(t)
	})

	t.Run("error insert data, duplicated", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.UserCore{}, errors.New("error insert data, duplicated")).Once()
		res, err := service.Register(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.UserID)
		assert.ErrorContains(t, err, "duplicated")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.UserCore{}, errors.New("server error")).Once()
		res, err := service.Register(arguments)
		assert.NotNil(t, err)
		assert.Equal(t, "", res.UserID)
		assert.ErrorContains(t, err, "internal server error")
		data.AssertExpectations(t)
	})
}
