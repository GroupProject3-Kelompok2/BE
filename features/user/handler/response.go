package handler

import "github.com/GroupProject3-Kelompok2/BE/features/user"

type registerResponse struct {
	UserID   string `json:"user_id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

func register(u user.UserCore) registerResponse {
	return registerResponse{
		UserID:   u.UserID,
		Fullname: u.Fullname,
		Email:    u.Email,
	}
}
