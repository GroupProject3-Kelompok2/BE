package main

import (
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"github.com/GroupProject3-Kelompok2/BE/app/config"
	"github.com/GroupProject3-Kelompok2/BE/app/database"
	"github.com/GroupProject3-Kelompok2/BE/app/router"
	"github.com/robfig/cron"

	"github.com/GroupProject3-Kelompok2/BE/utils/email"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := database.InitDatabase(*cfg)
	database.InitMigration(db)
	router.InitRouter(db, e)

	log := middlewares.Log()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	cron := cron.New()
	cron.AddFunc("1 * * * *", func() {
		email.SendEmail(db)
	})
	cron.Start()

	go func() {
		<-c
		log.Info("Program closed ...")
		pprof.StopCPUProfile()
		os.Exit(0)
	}()
	e.Logger.Fatal(e.Start(":8080"))
}
