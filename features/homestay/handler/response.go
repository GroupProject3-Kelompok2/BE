package handler

import "github.com/GroupProject3-Kelompok2/BE/features/homestay"

type HomestayResponse struct {
	HomestayID  string  `json:"homestay_id,omitempty" form:"homestay_id"`
	Name        string  `json:"name,omitempty" form:"name"`
	Description string  `json:"description,omitempty" form:"description"`
	Address     string  `json:"address,omitempty" form:"address"`
	Price       float64 `json:"price,omitempty" form:"price"`
	Status      bool    `json:"status,omitempty" form:"status"`
}

func HomestayCoreResponse(homestay homestay.HomestayCore) HomestayResponse {
	return HomestayResponse{
		HomestayID:  homestay.HomestayID,
		Name:        homestay.Name,
		Description: homestay.Description,
		Address:     homestay.Address,
		Price:       homestay.Price,
	}
}
