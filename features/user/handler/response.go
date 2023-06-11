package handler

import "github.com/GroupProject3-Kelompok2/BE/features/user"

type registerResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func register(u user.UserCore) registerResponse {
	return registerResponse{
		UserID:   u.UserID,
		Username: u.Username,
		Email:    u.Email,
	}
}
