package handler

import (
	"github.com/GroupProject3-Kelompok2/BE/features/user"
	"github.com/GroupProject3-Kelompok2/BE/utils/helper"
)

type loginResponse struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

type profileResponse struct {
	UserID         string           `json:"user_id"`
	Fullname       string           `json:"fullname"`
	Email          string           `json:"email"`
	Phone          string           `json:"phone"`
	ProfilePicture string           `json:"profile_picture"`
	Role           string           `json:"role"`
	CreatedAt      helper.LocalTime `json:"created_at"`
	UpdatedAt      helper.LocalTime `json:"updated_at"`
}

func profileUser(u user.UserCore) profileResponse {
	return profileResponse{
		UserID:         u.UserID,
		Fullname:       u.Fullname,
		Email:          u.Email,
		Phone:          u.Phone,
		ProfilePicture: u.ProfilePicture,
		Role:           u.Role,
		CreatedAt:      helper.LocalTime(u.CreatedAt),
		UpdatedAt:      helper.LocalTime(u.UpdatedAt),
	}
}
