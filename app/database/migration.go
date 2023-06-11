package database

import (
	homestay "github.com/GroupProject3-Kelompok2/BE/features/homestay/data"
	payment "github.com/GroupProject3-Kelompok2/BE/features/payment/data"
	reservation "github.com/GroupProject3-Kelompok2/BE/features/reservation/data"
	review "github.com/GroupProject3-Kelompok2/BE/features/review/data"
	user "github.com/GroupProject3-Kelompok2/BE/features/user/data"
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) error {
	err := db.SetupJoinTable(&payment.Payment{}, "Item", &reservation.Reservation{})
	if err != nil {
		log.Sugar().Warnf("setup join table error", err)
		return err
	}

	err = db.AutoMigrate(
		&user.User{},
		&homestay.Homestay{},
		&homestay.HomestayPicture{},
		&payment.Payment{},
		&review.Review{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	return err
}
