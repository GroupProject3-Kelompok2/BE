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

type UpdateProfileRequest struct {
	Fullname       *string `json:"fullname" form:"fullname"`
	Email          *string `json:"email" form:"email"`
	Phone          *string `json:"phone" form:"phone"`
	Password       *string `json:"password" form:"password"`
	ProfilePicture *string `json:"profile_picture" form:"profile_picture"`
}

type UpgradeProfileRequest struct {
	Role *string `json:"role" form:"role"`
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
	case *UpdateProfileRequest:
		if v.Fullname != nil {
			res.Fullname = *v.Fullname
		}
		if v.Email != nil {
			res.Email = *v.Email
		}
		if v.Phone != nil {
			res.Password = *v.Phone
		}
		if v.Password != nil {
			res.Password = *v.Password
		}
		if v.ProfilePicture != nil {
			res.ProfilePicture = *v.ProfilePicture
		}
	case *UpgradeProfileRequest:
		if v.Role != nil {
			res.Role = *v.Role
		}
	default:
		return user.UserCore{}
	}
	return res
}
