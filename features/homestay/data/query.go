package data

import (
	"errors"

	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
	"github.com/GroupProject3-Kelompok2/BE/utils/helper"
	"gorm.io/gorm"
)

type homestayQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) homestay.HomestayDataInterface {
	return &homestayQuery{
		db: db,
	}
}

func (repo *homestayQuery) Insert(input homestay.HomestayCore) error {
	homestayInputGorm := HomestayModel(input)
	homestayInputGorm.HomestayID = helper.GenerateId()

	tx := repo.db.Create(&homestayInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}
