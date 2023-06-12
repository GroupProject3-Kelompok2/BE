package data

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ReviewID   string         `gorm:"primaryKey;type:varchar(21)"`
	UserID     string         `gorm:"type:varchar(21)"`
	HomestayID string         `gorm:"type:varchar(21)"`
	Review     string         `gorm:"type:text"`
	Rating     uint8          `gorm:"not null"`
	CreatedAt  time.Time      `gorm:"type:datetime"`
	UpdatedAt  time.Time      `gorm:"type:datetime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	User       User           `gorm:"foreignKey:UserID"`
	Homestay   Homestay       `gorm:"foreignKey:HomestayID"`
}

type User struct {
	UserID          string         `gorm:"primaryKey;type:varchar(21)"`
	Fullname        string         `gorm:"type:varchar(100);not null"`
	Email           string         `gorm:"type:varchar(100);not null;unique"`
	Password        string         `gorm:"type:text"`
	ProfilePricture string         `gorm:"type:text"`
	Role            string         `gorm:"type:enum('user', 'hoster'); default:'user'"`
	CreatedAt       time.Time      `gorm:"type:datetime"`
	UpdatedAt       time.Time      `gorm:"type:datetime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Homestays       []Homestay     `gorm:"foreignKey:UserID"`
	Reviews         []Review       `gorm:"foreignKey:UserID"`
}

type Homestay struct {
	HomestayID  string         `gorm:"primaryKey;type:varchar(21)"`
	UserID      string         `gorm:"type:varchar(21)"`
	Name        string         `gorm:"type:text"`
	Description string         `gorm:"type:text"`
	Address     string         `gorm:"type:text"`
	Status      bool           `gorm:"type:boolean"`
	CreatedAt   time.Time      `gorm:"type:datetime"`
	UpdatedAt   time.Time      `gorm:"type:datetime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	User        User           `gorm:"foreignKey:UserID"`
	Reviews     []Review       `gorm:"foreignKey:HomestayID"`
}
