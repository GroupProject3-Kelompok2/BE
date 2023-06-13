package data

import (
	"errors"

	"github.com/GroupProject3-Kelompok2/BE/features/review"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

var log = middlewares.Log()

type reviewQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) review.ReviewData {
	return &reviewQuery{
		db: db,
	}
}

// AddReview implements review.ReviewData
func (rq *reviewQuery) AddReview(userId string, request review.ReviewCore) error {
	reviewId, err := gonanoid.New()
	if err != nil {
		log.Warn("error while create nano_id for user_id")
		return nil
	}

	request.ReviewID = reviewId
	request.UserID = userId
	req := reviewEntities(request)
	query := rq.db.Table("reviews").Create(&req)
	if query.Error != nil {
		log.Error("error inserting data")
		return query.Error
	}

	if query.RowsAffected == 0 {
		log.Warn("no feedback has been registered")
		return errors.New("no feedback has been registered")
	}

	return nil
}

// EditReview implements review.ReviewData
func (rq *reviewQuery) EditReview(userId string, reviewId string, request review.ReviewCore) error {
	req := reviewEntities(request)
	query := rq.db.Table("reviews").Where("review_id = ? AND user_id = ?", reviewId, userId).Updates(&req)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("review record not found")
		return errors.New("feedback record not found")
	}

	if query.RowsAffected == 0 {
		log.Warn("no review has been created")
		return errors.New("row affected : 0")
	}

	if query.Error != nil {
		log.Error("error while updating review")
		return errors.New("duplicate data entry")
	}

	return nil
}
