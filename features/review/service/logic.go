package service

import (
	"errors"
	"strings"

	"github.com/GroupProject3-Kelompok2/BE/features/review"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"github.com/go-playground/validator/v10"
)

var log = middlewares.Log()

type reviewService struct {
	query    review.ReviewData
	validate *validator.Validate
}

func New(rd review.ReviewData, v *validator.Validate) review.ReviewService {
	return &reviewService{
		query:    rd,
		validate: v,
	}
}

// AddReview implements review.ReviewService
func (rs *reviewService) AddReview(userId string, request review.ReviewCore) error {
	errVal := rs.validate.Struct(request)
	if errVal != nil {
		switch {
		case strings.Contains(errVal.Error(), "HomestayID"):
			log.Warn("homestay_id cannot be empty")
			return errors.New("homestay_id cannot be empty")
		case strings.Contains(errVal.Error(), "Review"):
			log.Warn("review cannot be empty")
			return errors.New("review cannot be empty")
		}
	}

	err := rs.query.AddReview(userId, request)
	if err != nil {
		log.Error("internal server error")
		return errors.New("internal server error")
	}

	return nil
}

// EditReview implements review.ReviewService
func (rs *reviewService) EditReview(userId string, reviewId string, request review.ReviewCore) error {
	err := rs.query.EditReview(userId, reviewId, request)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "not found"):
			log.Error("review record not found")
			return errors.New("review record not found")
		default:
			log.Error("internal server error")
			return errors.New("internal server error")
		}
	}

	return nil
}

// DeleteReview implements review.ReviewService
func (rs *reviewService) DeleteReview(userId string, reviewId string) error {
	err := rs.query.DeleteReview(userId, reviewId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			log.Error("review record not found")
			return errors.New("review record not found")
		} else {
			log.Error("internal server error")
			return errors.New("internal server error")
		}
	}

	return nil
}
