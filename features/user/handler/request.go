package handler

import "github.com/GroupProject3-Kelompok2/BE/features/user"

type RegisterRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(data interface{}) user.UserCore {
	res := user.UserCore{}
	switch v := data.(type) {
	case RegisterRequest:
		res.Fullname = v.Fullname
		res.Email = v.Email
		res.Phone = v.Phone
		res.Password = v.Password
	case LoginRequest:
		res.Email = v.Email
		res.Password = v.Password
	default:
		return user.UserCore{}
	}
	return res
}
