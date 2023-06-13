package service

import (
	"errors"
	"testing"

	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
	"github.com/GroupProject3-Kelompok2/BE/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHomestayPictures(t *testing.T) {
	data := mocks.NewHomestayDataInterface(t)
	HomestayID := "nanoid"
	request := homestay.HomestayPictureCore{
		URL: "https://example.com/image.jpg",
	}

	service := New(data)

	t.Run("success", func(t *testing.T) {
		data.On("HomestayPictures", HomestayID, request).Return(nil).Once()

		err := service.HomestayPictures(HomestayID, request)

		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("homestay profile not found", func(t *testing.T) {
		notFoundErr := errors.New("homestay profile record not found")
		data.On("HomestayPictures", HomestayID, request).Return(notFoundErr).Once()

		err := service.HomestayPictures(HomestayID, request)

		assert.NotNil(t, err)
		assert.EqualError(t, err, notFoundErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		internalErr := errors.New("internal server error")
		data.On("HomestayPictures", HomestayID, request).Return(internalErr).Once()

		err := service.HomestayPictures(HomestayID, request)

		assert.NotNil(t, err)
		assert.EqualError(t, err, internalErr.Error(), "Expected error message does not match")
		data.AssertExpectations(t)
	})
}
