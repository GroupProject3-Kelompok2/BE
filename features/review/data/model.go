package data

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ReviewID   uint           `gorm:"primaryKey;autoIncrement"`
	Review     string         `gorm:"type:text"`
	Rating     uint           `gorm:"not null"`
	CreatedAt  time.Time      `gorm:"type:datetime"`
	UpdatedAt  time.Time      `gorm:"type:datetime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Reviewer   string         `gorm:"type:varchar(50)"`
	HomestayID string         `gorm:"type:varchar(50)"`
	User       User           `gorm:"foreignKey:Reviewer"`
	Homestay   Homestay
}

type User struct {
	UserID      string         `gorm:"primaryKey;type:varchar(50)"`
	Username    string         `gorm:"type:varchar(100);not null"`
	Email       string         `gorm:"type:varchar(100);not null;unique"`
	Password    string         `gorm:"type:text"`
	UserPicture string         `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time      `gorm:"type:datetime"`
	UpdatedAt   time.Time      `gorm:"type:datetime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Homestays   []Homestay     `gorm:"foreignKey:Hoster"`
	Reviews     []Review       `gorm:"foreignKey:Reviewer"`
}

type Homestay struct {
	HomestayID string         `gorm:"primaryKey;type:varchar(50)"`
	CreatedAt  time.Time      `gorm:"type:datetime"`
	UpdatedAt  time.Time      `gorm:"type:datetime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Hoster     string         `gorm:"type:varchar(50)"`
	Reviews    []Review       `gorm:"foreignKey:HomestayID"`
}
