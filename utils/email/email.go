package email

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"

	"github.com/GroupProject3-Kelompok2/BE/app/config"
	"github.com/GroupProject3-Kelompok2/BE/features/user/data"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

var log = middlewares.Log()

func SendEmailVerCode(user data.User) error {
	var basePath string
	wd, _ := os.Getwd()
	if string(wd[len(wd)-13]) == "u" {
		basePath = filepath.Join(wd, "../../../", "./utils/email/info.html")
	} else {
		basePath = filepath.Join(wd, "../", "./utils/email/info.html")
	}
	var body bytes.Buffer
	t, err := template.ParseFiles(basePath)
	if err != nil {
		return err
	}
	t.Execute(&body, user)
	m := gomail.NewMessage()
	m.SetHeader("From", config.GOMAIL_EMAIL)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Verification Code")
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer(config.GOMAIL_HOST, config.GOMAIL_PORT, config.GOMAIL_EMAIL, config.GOMAIL_PASSWORD)
	err = d.DialAndSend(m)
	if err != nil {
		log.Sugar().Warn("Error sending email:" + err.Error())
		return err
	}
	log.Info("Email sent.")
	return nil
}

func SendEmail(db *gorm.DB) {
	var user []data.User
	result := db.Find(&user)
	if result.Error != nil {
		log.Error("can't get user data")
	}

	var body bytes.Buffer
	t, err := template.ParseFiles("./utils/email/info.html")
	if err != nil {
		log.Error(err.Error())
	}
	m := gomail.NewMessage()
	m.SetHeader("From", config.GOMAIL_EMAIL)
	m.SetHeader("Subject", "Information")
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer(config.GOMAIL_HOST, config.GOMAIL_PORT, config.GOMAIL_EMAIL, config.GOMAIL_PASSWORD)
	for _, v := range user {
		m.SetHeader("To", v.Email)
		t.Execute(&body, v)
		err := d.DialAndSend(m)
		if err != nil {
			log.Sugar().Error("Error sending email:" + err.Error())

		}
	}
	log.Info("Email sent.")
}
