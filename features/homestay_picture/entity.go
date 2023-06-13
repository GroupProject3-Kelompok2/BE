package homestay_picture

import (
	"time"
)

type HomestayPictureCore struct {
	PictureID  string
	HomestayID string
	URL        string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
