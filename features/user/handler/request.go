package handler

import "github.com/GroupProject3-Kelompok2/BE/features/user"

type RegisterRequest struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(data interface{}) user.UserCore {
	res := user.UserCore{}
	switch v := data.(type) {
	case RegisterRequest:
		res.Username = v.Username
		res.Email = v.Email
		res.Password = v.Password
	default:
		return user.UserCore{}
	}
	return res
}
