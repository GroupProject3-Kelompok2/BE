package data

import (
	"errors"

	"github.com/GroupProject3-Kelompok2/BE/features/user"
	"github.com/GroupProject3-Kelompok2/BE/utils/helper"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

var log = middlewares.Log()

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

// Register implements user.UserData
func (uq *userQuery) Register(request user.UserCore) (user.UserCore, error) {
	userID, err := gonanoid.New()
	if err != nil {
		log.Warn("error while create nano_id for user_id")
		return user.UserCore{}, nil
	}

	hashed, err := helper.HashPassword(request.Password)
	if err != nil {
		log.Error("error while hashing password")
		return user.UserCore{}, errors.New("error while hashing password")
	}

	request.UserID = userID
	request.Password = hashed
	request.ProfilePicture = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png"
	req := userEntities(request)
	query := uq.db.Table("users").Create(&req)
	if query.Error != nil {
		log.Error("error insert data, duplicated")
		return user.UserCore{}, errors.New("error insert data, duplicated")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		log.Warn("no user has been created")
		return user.UserCore{}, errors.New("row affected : 0")
	}

	log.Sugar().Infof("new user has been created: %s", req.UserID)
	return userModels(req), nil
}

// Login implements user.UserData
func (uq *userQuery) Login(request user.UserCore) (user.UserCore, string, error) {
	result := User{}
	query := uq.db.Table("users").Where("email = ?", request.Email).First(&result)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("user record not found, invalid email and password")
		return user.UserCore{}, "", errors.New("invalid email and password")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		log.Warn("no user has been created")
		return user.UserCore{}, "", errors.New("row affected : 0")
	}

	match1 := helper.MatchPassword(request.Password, result.Password)
	log.Sugar().Warnf("match password: %v", match1)
	if !match1 {
		return user.UserCore{}, "", errors.New("password does not match")
	}

	token, err := middlewares.CreateToken(result.UserID, result.Role)
	if err != nil {
		log.Error("error while creating jwt token")
		return user.UserCore{}, "", errors.New("error while creating jwt token")
	}

	log.Sugar().Infof("user has been logged in: %s", result.UserID)
	return userModels(result), token, nil
}

// ProfileUser implements user.UserData
func (uq *userQuery) ProfileUser(userId string) (user.UserCore, error) {
	users := User{}
	query := uq.db.First(&users, "user_id = ?", userId)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("user profile record not found")
		return user.UserCore{}, errors.New("user profile record not found")
	}

	return userModels(users), nil
}

// UpdateProfile implements user.UserData
func (uq *userQuery) UpdateProfile(userId string, request user.UserCore) error {
	req := userEntities(request)
	query := uq.db.Table("users").Where("user_id = ?", userId).Updates(&req)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("user profile record not found")
		return errors.New("user profile record not found")
	}

	if query.RowsAffected == 0 {
		log.Warn("no user has been created")
		return errors.New("row affected : 0")
	}

	if query.Error != nil {
		log.Error("error while updating user")
		return errors.New("duplicate data entry")
	}

	return nil
}

// DeactiveUser implements user.UserData
func (uq *userQuery) DeactiveUser(userId string) error {
	req := User{}
	query := uq.db.Table("users").Where("user_id = ?", userId).Delete(&req)

	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("user profile record not found")
		return errors.New("user profile record not found")
	}

	if query.RowsAffected == 0 {
		log.Warn("no user has been created")
		return errors.New("row affected : 0")
	}

	if query.Error != nil {
		log.Error("error while deactivate user")
		return errors.New("duplicate data entry")
	}

	return nil
}

// UpgradeProfile implements user.UserData
func (uq *userQuery) UpgradeProfile(userId string, request user.UserCore) error {
	query := uq.db.Table("users").Where("user_id = ?", userId).Updates(map[string]interface{}{
		"role": "hoster",
	})

	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("user profile record not found")
		return errors.New("user profile record not found")
	}

	if query.RowsAffected == 0 {
		log.Warn("no user has been created")
		return errors.New("row affected : 0")
	}

	if query.Error != nil {
		log.Error("error while deactivate user")
		return errors.New("duplicate data entry")
	}

	return nil
}
