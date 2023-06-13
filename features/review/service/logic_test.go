package service

import (
	"errors"
	"testing"

	"github.com/GroupProject3-Kelompok2/BE/features/review"
	"github.com/GroupProject3-Kelompok2/BE/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestAddReview(t *testing.T) {
	data := mocks.NewReviewData(t)
	validate := validator.New()
	UserID := "nanoid"
	request := review.ReviewCore{
		Review:     "admin@gmail.com",
		Rating:     5,
		HomestayID: "nanoid",
	}
	service := New(data, validate)

	t.Run("success add review", func(t *testing.T) {
		data.On("AddReview", UserID, request).Return(nil).Once()
		err := service.AddReview(UserID, request)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("homestay_id cannot be empty", func(t *testing.T) {
		invalidRequest := review.ReviewCore{
			Review:     "admin@gmail.com",
			Rating:     5,
			HomestayID: "",
		}
		err := service.AddReview(UserID, invalidRequest)
		expectedErr := errors.New("homestay_id cannot be empty")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("review cannot be empty", func(t *testing.T) {
		invalidRequest := review.ReviewCore{
			Review:     "",
			Rating:     5,
			HomestayID: "nanoid",
		}
		err := service.AddReview(UserID, invalidRequest)
		expectedErr := errors.New("review cannot be empty")
		assert.NotNil(t, err)
		assert.EqualError(t, err, expectedErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("AddReview", UserID, request).Return(errors.New("internal server error")).Once()
		err := service.AddReview(UserID, request)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "internal server error")
		data.AssertExpectations(t)
	})
}

func TestEditReview(t *testing.T) {
	data := mocks.NewReviewData(t)
	validate := validator.New()
	UserID := "nanoid"
	ReviewID := "nanoid"
	request := review.ReviewCore{
		Review:     "admin@gmail.com",
		Rating:     5,
		HomestayID: "nanoid",
	}
	service := New(data, validate)

	t.Run("success edit review", func(t *testing.T) {
		data.On("EditReview", UserID, ReviewID, request).Return(nil).Once()
		err := service.EditReview(UserID, ReviewID, request)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("review record not found", func(t *testing.T) {
		notFoundErr := errors.New("review record not found")
		data.On("EditReview", UserID, ReviewID, request).Return(notFoundErr).Once()
		err := service.EditReview(UserID, ReviewID, request)
		assert.NotNil(t, err)
		assert.EqualError(t, err, notFoundErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		internalErr := errors.New("internal server error")
		data.On("EditReview", UserID, ReviewID, request).Return(internalErr).Once()
		err := service.EditReview(UserID, ReviewID, request)
		assert.NotNil(t, err)
		assert.EqualError(t, err, internalErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})
}
