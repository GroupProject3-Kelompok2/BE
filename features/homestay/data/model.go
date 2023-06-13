package data

import (
	"time"

	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
	reservation "github.com/GroupProject3-Kelompok2/BE/features/reservation/data"
	review "github.com/GroupProject3-Kelompok2/BE/features/review/data"

	"gorm.io/gorm"
)

type Homestay struct {
	HomestayID   string                    `gorm:"primaryKey;type:varchar(21)"`
	UserID       string                    `gorm:"type:varchar(21)"`
	Name         string                    `gorm:"type:text"`
	Description  string                    `gorm:"type:text"`
	Address      string                    `gorm:"type:text"`
	Price        float64                   `gorm:"type:decimal(15,2)"`
	Status       bool                      `gorm:"type:boolean;default:true"`
	CreatedAt    time.Time                 `gorm:"type:datetime"`
	UpdatedAt    time.Time                 `gorm:"type:datetime"`
	DeletedAt    gorm.DeletedAt            `gorm:"index"`
	Pictures     []HomestayPicture         `gorm:"foreignKey:HomestayID"`
	Reservations []reservation.Reservation `gorm:"foreignKey:HomestayID"`
	Reviews      []review.Review           `gorm:"foreignKey:HomestayID"`
}

type HomestayPicture struct {
	PictureID  string `gorm:"primaryKey;type:varchar(21)"`
	HomestayID string `gorm:"type:varchar(21)"`
	URL        string `gorm:"type:text"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func HomestayCore(homestayData Homestay) homestay.HomestayCore {
	return homestay.HomestayCore{
		HomestayID:  homestayData.HomestayID,
		UserID:      homestayData.UserID,
		Name:        homestayData.Name,
		Description: homestayData.Description,
		Address:     homestayData.Address,
		Price:       homestayData.Price,
		Status:      homestayData.Status,
		CreatedAt:   homestayData.CreatedAt,
		UpdatedAt:   homestayData.UpdatedAt,
	}
}

func HomestayPictureCore(pictureData HomestayPicture) homestay.HomestayPictureCore {
	return homestay.HomestayPictureCore{
		PictureID:  pictureData.PictureID,
		HomestayID: pictureData.HomestayID,
		URL:        pictureData.URL,
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
		Status:      dataCore.Status,
		CreatedAt:   dataCore.CreatedAt,
		UpdatedAt:   dataCore.UpdatedAt,
	}
}
