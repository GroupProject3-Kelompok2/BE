package homestay_picture

import (
	"time"

	"github.com/GroupProject3-Kelompok2/BE/features/review"
	"gorm.io/gorm"
)

type HomestayPictureCore struct {
	HomestayID       string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	Hoster           string
	HomestayPictures []HomestayPictureCore
	Reviews          []review.ReviewCore
}
