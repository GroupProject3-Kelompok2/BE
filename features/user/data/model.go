package data

import (
	"time"

	homestay "github.com/GroupProject3-Kelompok2/BE/features/homestay/data"
	payment "github.com/GroupProject3-Kelompok2/BE/features/payment/data"
	review "github.com/GroupProject3-Kelompok2/BE/features/review/data"
	"github.com/GroupProject3-Kelompok2/BE/features/user"
	"gorm.io/gorm"
)

type User struct {
	UserID      string              `gorm:"primaryKey;type:varchar(50)"`
	Username    string              `gorm:"type:varchar(100);not null"`
	Email       string              `gorm:"type:varchar(100);not null;unique"`
	Password    string              `gorm:"type:text"`
	UserPicture string              `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time           `gorm:"type:datetime"`
	UpdatedAt   time.Time           `gorm:"type:datetime"`
	DeletedAt   gorm.DeletedAt      `gorm:"index"`
	Homestays   []homestay.Homestay `gorm:"foreignKey:Hoster"`
	Reviews     []review.Review     `gorm:"foreignKey:Reviewer"`
	Payments    []payment.Payment   `gorm:"foreignKey:Buyer"`
}

// User-model to user-core
func userModels(u User) user.UserCore {
	return user.UserCore{
		UserID:      u.UserID,
		Username:    u.Username,
		Email:       u.Email,
		Password:    u.Password,
		UserPicture: u.UserPicture,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

// User-core to user-model
func userEntities(u user.UserCore) User {
	return User{
		UserID:      u.UserID,
		Username:    u.Username,
		Email:       u.Email,
		Password:    u.Password,
		UserPicture: u.UserPicture,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
