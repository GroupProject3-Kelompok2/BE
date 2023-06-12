package data

import (
	"time"

	homestay "github.com/GroupProject3-Kelompok2/BE/features/homestay/data"
	reservation "github.com/GroupProject3-Kelompok2/BE/features/reservation/data"
	review "github.com/GroupProject3-Kelompok2/BE/features/review/data"
	"github.com/GroupProject3-Kelompok2/BE/features/user"
	"gorm.io/gorm"
)

type User struct {
	UserID         string                    `gorm:"primaryKey;type:varchar(21)"`
	Fullname       string                    `gorm:"type:varchar(100);not null"`
	Email          string                    `gorm:"type:varchar(100);not null;unique"`
	Phone          string                    `gorm:"type:varchar(15);not null"`
	Password       string                    `gorm:"type:text;not null"`
	ProfilePicture string                    `gorm:"type:text"`
	Role           string                    `gorm:"type:enum('user', 'hoster'); default:'user'"`
	CreatedAt      time.Time                 `gorm:"type:datetime"`
	UpdatedAt      time.Time                 `gorm:"type:datetime"`
	DeletedAt      gorm.DeletedAt            `gorm:"index"`
	Homestays      []homestay.Homestay       `gorm:"foreignKey:UserID"`
	Reservations   []reservation.Reservation `gorm:"foreignKey:UserID"`
	Reviews        []review.Review           `gorm:"foreignKey:UserID"`
}

// User-model to user-core
func userModels(u User) user.UserCore {
	return user.UserCore{
		UserID:         u.UserID,
		Fullname:       u.Fullname,
		Email:          u.Email,
		Phone:          u.Phone,
		Password:       u.Password,
		ProfilePicture: u.ProfilePicture,
		Role:           u.Role,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}

// User-core to user-model
func userEntities(u user.UserCore) User {
	return User{
		UserID:         u.UserID,
		Fullname:       u.Fullname,
		Email:          u.Email,
		Phone:          u.Phone,
		Password:       u.Password,
		ProfilePicture: u.ProfilePicture,
		Role:           u.Role,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}
