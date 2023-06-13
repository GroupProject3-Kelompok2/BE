package handler

import "github.com/GroupProject3-Kelompok2/BE/features/homestay"

type HomestayRequest struct {
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Address     string  `json:"address" form:"address"`
	Price       float64 `json:"price" form:"price"`
	Status      bool    `json:"status" form:"status"`
}

func HomestayRequestCore(homestayRequest HomestayRequest) homestay.HomestayCore {
	return homestay.HomestayCore{
		Name:        homestayRequest.Name,
		Description: homestayRequest.Description,
		Address:     homestayRequest.Address,
		Price:       homestayRequest.Price,
		Status:      homestayRequest.Status,
	}
}
