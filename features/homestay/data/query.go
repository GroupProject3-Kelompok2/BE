package data

import (
	"errors"

	"github.com/GroupProject3-Kelompok2/BE/features/homestay"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	gonanoid "github.com/matoous/go-nanoid/v2"

	"gorm.io/gorm"
)

var log = middlewares.Log()

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

func (repo *homestayQuery) SelectAll() ([]homestay.HomestayCore, error) {
	var homestaysData []Homestay
	tx := repo.db.Find(&homestaysData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("error homestays not found")
	}

	var homestaysCoreAll []homestay.HomestayCore
	for _, value := range homestaysData {
		homestayCore := HomestayCore(value)
		homestaysCoreAll = append(homestaysCoreAll, homestayCore)
	}
	return homestaysCoreAll, nil
}

func (repo *homestayQuery) SelectById(homestayId string) (homestay.HomestayCore, error) {
	var homestayGorm Homestay
	tx := repo.db.Where("homestay_id = ?", homestayId).First(&homestayGorm)
	if tx.Error != nil {
		return homestay.HomestayCore{}, errors.New("error homestay not found")
	}

	homestayCore := HomestayCore(homestayGorm)
	return homestayCore, nil
}

// UpdateHomestayPictures implements homestay.HomestayDataInterface
func (hq *homestayQuery) HomestayPictures(homestayId string, req homestay.HomestayPictureCore) error {
	pictureId, err := gonanoid.New()
	if err != nil {
		log.Warn("error while creating nano_id for user_id")
		return nil
	}

	req.PictureID = pictureId
	req.HomestayID = homestayId
	pic := homestayPictureEntities(req)
	query := hq.db.Table("homestay_pictures").Create(&pic)
	if query.Error != nil {
		log.Error("error inserting data, duplicated")
		return errors.New("error inserting data, duplicated")
	}

	rowAffected := query.RowsAffected
	if rowAffected == 0 {
		log.Warn("no picture has been created")
		return errors.New("row affected: 0")
	}

	log.Sugar().Infof("new homestay_picture has been created: %s", req.PictureID)
	return nil
}
