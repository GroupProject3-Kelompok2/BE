package handler

import "github.com/GroupProject3-Kelompok2/BE/features/review"

type AddReviewRequest struct {
	Review     string `json:"review" form:"review"`
	Rating     uint8  `json:"rating" form:"rating"`
	HomestayID string `json:"homestay_id" form:"homestay_id"`
}

func RequestToCore(data interface{}) review.ReviewCore {
	res := review.ReviewCore{}
	switch v := data.(type) {
	case AddReviewRequest:
		res.Review = v.Review
		res.Rating = v.Rating
		res.HomestayID = v.HomestayID
	default:
		return res
	}
	return res
}
