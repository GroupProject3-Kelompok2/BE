package review

import (
	"time"

	"gorm.io/gorm"
)

type ReviewCore struct {
	ReviewID   uint
	Review     string
	Rating     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	Reviewer   string
	HomestayID string
	User       UserCore
	Homestay   HomestayCore
}

type UserCore struct {
	UserID      string
	Username    string
	Email       string
	Password    string
	UserPicture string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Homestays   []HomestayCore
	Reviews     []ReviewCore
}

type HomestayCore struct {
	HomestayID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Hoster     string
}
