package validation

import (
	"fmt"
	"log"

	"github.com/GroupProject3-Kelompok2/BE/features/user"
	"github.com/go-playground/validator/v10"
)

type Register struct {
	Fullname string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,strongPassword"`
}

type Login struct {
	Fullname string `validate:"required"`
	Password string `validate:"required"`
}

func UserValidate(option string, data interface{}) (interface{}, error) {
	switch option {
	case "register":
		res := Register{}
		if v, ok := data.(user.UserCore); ok {
			res.Fullname = v.Fullname
			res.Email = v.Email
			res.Password = v.Password
		}
		err := Authenticate(res)
		if err != nil {
			return nil, err
		}
		return res, nil
	case "login":
		res := Login{}
		if v, ok := data.(Login); ok {
			res.Fullname = v.Fullname
			res.Password = v.Password
		}
		err := Authenticate(res)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, fmt.Errorf("invalid option")
	}
}

func Authenticate(data interface{}) error {
	validate := validator.New()
	err := validate.RegisterValidation("strongPassword", StrongPassword)
	if err != nil {
		return err
	}

	err = validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePasswordValidator(password string) error {
	minChars := strongPassword(password)
	if minChars > 0 {
		return fmt.Errorf("password strength is low, minimum %d characters need to be added", minChars)
	}
	return nil
}

func StrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	minChars := strongPassword(password)
	if minChars > 0 {
		log.Printf("Password strength is low. Minimum %d characters need to be added.", minChars)
		return false
	}
	return true
}

func strongPassword(password string) uint32 {
	hasDigit := false
	hasLower := false
	hasUpper := false
	hasSpecial := false
	toAdd := 0

	for i := 0; i < len(password); i++ {
		val := uint32(password[i])
		if val >= 65 && val <= 90 {
			hasUpper = true
		}
		if val >= 97 && val <= 122 {
			hasLower = true
		}
		if val >= 48 && val <= 57 {
			hasDigit = true
		}
		if val >= 33 && val <= 45 {
			hasSpecial = true
		}
	}

	if !hasUpper {
		toAdd++
	}
	if !hasLower {
		toAdd++
	}
	if !hasDigit {
		toAdd++
	}
	if !hasSpecial {
		toAdd++
	}

	addedLen := len(password) + toAdd
	if addedLen > 6 {
		return uint32(toAdd)
	}

	return uint32(6 - addedLen + toAdd)
}
