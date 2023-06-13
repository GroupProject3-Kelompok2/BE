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
