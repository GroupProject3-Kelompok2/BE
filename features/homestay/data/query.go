package data

import (
	"errors"

	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
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
	tx := repo.db.Create(&homestayInputGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}

func (repo *homestayQuery) UpdateById(userId string, homestayId string, input homestay.HomestayCore) error {
	var homestayGorm Homestay
	tx := repo.db.Where("homestay_id = ?", homestayId).First(&homestayGorm)
	if tx.Error != nil {
		return errors.New("error homestay not found")
	}

	if homestayGorm.UserID != userId {
		return errors.New("unauthorize")
	}

	homestayInputGorm := HomestayModel(input)
	tx = repo.db.Model(&homestayGorm).Updates(homestayInputGorm)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + "failed to update homestay")
	}

	return nil
}

func (repo *homestayQuery) DeleteById(userId string, homestayId string) error {
	var homestayGorm Homestay
	tx := repo.db.Where("homestay_id = ?", homestayId).First(&homestayGorm)
	if tx.Error != nil {
		return errors.New("error homestay not found")
	}

	if homestayGorm.UserID != userId {
		return errors.New("unauthorize")
	}

	tx = repo.db.Delete(&homestayGorm)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + "failed to delete homestay")
	}

	return nil
}
