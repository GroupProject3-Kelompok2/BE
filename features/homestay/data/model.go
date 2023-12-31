package data

import (
	"time"

	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
	reservation "github.com/GroupProject3-Kelompok2/BE/features/reservation/data"
	review "github.com/GroupProject3-Kelompok2/BE/features/review/data"

	"gorm.io/gorm"
)

type Homestay struct {
	HomestayID       string                    `gorm:"primaryKey;type:varchar(21)"`
	UserID           string                    `gorm:"type:varchar(21)"`
	Name             string                    `gorm:"type:text"`
	Description      string                    `gorm:"type:text"`
	Address          string                    `gorm:"type:text"`
	Price            float64                   `gorm:"type:decimal(15,2)"`
	CreatedAt        time.Time                 `gorm:"type:datetime"`
	UpdatedAt        time.Time                 `gorm:"type:datetime"`
	DeletedAt        gorm.DeletedAt            `gorm:"index"`
	HomestayPictures []HomestayPicture         `gorm:"foreignKey:HomestayID"`
	Reservations     []reservation.Reservation `gorm:"foreignKey:HomestayID"`
	Reviews          []review.Review           `gorm:"foreignKey:HomestayID"`
}

type HomestayPicture struct {
	HomestayPictureID  string `gorm:"primaryKey;type:varchar(21)"`
	HomestayID         string `gorm:"type:varchar(21)"`
	HomestayPictureURL string `gorm:"type:text"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func HomestayCore(homestayData Homestay) homestay.HomestayCore {
	return homestay.HomestayCore{
		HomestayID:  homestayData.HomestayID,
		UserID:      homestayData.UserID,
		Name:        homestayData.Name,
		Description: homestayData.Description,
		Address:     homestayData.Address,
		Price:       homestayData.Price,
		CreatedAt:   homestayData.CreatedAt,
		UpdatedAt:   homestayData.UpdatedAt,
	}
}

func HomestayPictureCore(pictureData HomestayPicture) homestay.HomestayPictureCore {
	return homestay.HomestayPictureCore{
		PictureID:  pictureData.HomestayPictureID,
		HomestayID: pictureData.HomestayID,
		URL:        pictureData.HomestayPictureURL,
		CreatedAt:  pictureData.CreatedAt,
		UpdatedAt:  pictureData.UpdatedAt,
	}
}

func HomestayModel(dataCore homestay.HomestayCore) Homestay {
	return Homestay{
		HomestayID:  dataCore.HomestayID,
		UserID:      dataCore.UserID,
		Name:        dataCore.Name,
		Description: dataCore.Description,
		Address:     dataCore.Address,
		Price:       dataCore.Price,
		CreatedAt:   dataCore.CreatedAt,
		UpdatedAt:   dataCore.UpdatedAt,
	}
}

// homestayPicture-core to homestayPicture-model
func homestayPictureEntities(pictureData homestay.HomestayPictureCore) HomestayPicture {
	return HomestayPicture{
		HomestayPictureID:  pictureData.PictureID,
		HomestayID:         pictureData.HomestayID,
		HomestayPictureURL: pictureData.URL,
		CreatedAt:          pictureData.CreatedAt,
		UpdatedAt:          pictureData.UpdatedAt,
	}
}

// homestay-model to homestay-core
func modelToCore(h Homestay) homestay.HomestayCore {
	var pictureURL string
	if len(h.HomestayPictures) > 0 {
		pictureURL = h.HomestayPictures[0].HomestayPictureURL
	}

	var pictures []homestay.HomestayPictureCore
	if pictureURL != "" {
		pictures = []homestay.HomestayPictureCore{
			{
				URL: pictureURL,
			},
		}
	}

	var reviews []homestay.ReviewCore
	totalRating := 0
	for _, r := range h.Reviews {
		reviewCore := homestay.ReviewCore{
			Review: r.Review,
			Rating: uint8(r.Rating),
		}
		reviews = append(reviews, reviewCore)

		totalRating += int(r.Rating)
	}

	var averageRating float32
	if len(h.Reviews) > 0 {
		averageRating = float32(totalRating) / float32(len(h.Reviews))
	}

	response := homestay.HomestayCore{
		HomestayID:    h.HomestayID,
		Name:          h.Name,
		Description:   h.Description,
		Address:       h.Address,
		Price:         h.Price,
		Pictures:      pictures,
		Reviews:       reviews,
		TotalReviews:  uint(len(h.Reviews)),
		AverageRating: averageRating,
	}

	return response
}
