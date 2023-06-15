package handler

import "github.com/GroupProject3-Kelompok2/BE/features/homestay"

type HomestayResponse struct {
	HomestayID      string            `json:"homestay_id,omitempty"`
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
	Address         string            `json:"address,omitempty"`
	Price           float64           `json:"price,omitempty"`
	TotalReviews    uint              `json:"total_reviews,omitempty"`
	AverageRating   float32           `json:"average_rating,omitempty"`
	HomestayPicture []HomestayPicture `json:"homestay_pictures,omitempty"`
}

type HomestayPicture struct {
	HomestayPictureURL string `json:"homestay_picture,omitempty"`
}

func searchHomestay(h homestay.HomestayCore) HomestayResponse {
	pictures := make([]HomestayPicture, len(h.Pictures))
	for i, p := range h.Pictures {
		pictures[i] = HomestayPicture{
			HomestayPictureURL: p.URL,
		}
	}

	response := HomestayResponse{
		HomestayID:      h.HomestayID,
		Name:            h.Name,
		Description:     h.Description,
		Address:         h.Address,
		Price:           h.Price,
		HomestayPicture: pictures,
		TotalReviews:    h.TotalReviews,
		AverageRating:   h.AverageRating,
	}

	return response
}
