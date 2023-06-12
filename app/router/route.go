package router

import (
	_homestayData "github.com/GroupProject3-Kelompok2/BE/features/homestay/data"
	_homestayHandler "github.com/GroupProject3-Kelompok2/BE/features/homestay/handler"
	_homestayService "github.com/GroupProject3-Kelompok2/BE/features/homestay/service"
	ud "github.com/GroupProject3-Kelompok2/BE/features/user/data"
	uh "github.com/GroupProject3-Kelompok2/BE/features/user/handler"
	us "github.com/GroupProject3-Kelompok2/BE/features/user/service"
	"github.com/GroupProject3-Kelompok2/BE/utils/middlewares"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	initUserRouter(db, e)
	initHomestayRouter(db, e)
}

func initUserRouter(db *gorm.DB, e *echo.Echo) {
	userData := ud.New(db)
	validate := validator.New()
	userService := us.New(userData, validate)
	userHandler := uh.New(userService)

	e.POST("/register", userHandler.Register())
	e.POST("/login", userHandler.Login())
	e.GET("/users/:id", userHandler.ProfileUser(), middlewares.JWTMiddleware())
	e.PUT("/users/:id", userHandler.UpdateUser(), middlewares.JWTMiddleware())
}

func initHomestayRouter(db *gorm.DB, e *echo.Echo) {
	homestayData := _homestayData.New(db)
	homestayService := _homestayService.New(homestayData)
	homestayHandler := _homestayHandler.New(homestayService)

	e.POST("/homestays", homestayHandler.CreateHomestay)
}
