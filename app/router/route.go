package router

import (
	_homestayData "github.com/GroupProject3-Kelompok2/BE/features/homestay/data"
	_homestayHandler "github.com/GroupProject3-Kelompok2/BE/features/homestay/handler"
	_homestayService "github.com/GroupProject3-Kelompok2/BE/features/homestay/service"
	_reservationData "github.com/GroupProject3-Kelompok2/BE/features/reservation/data"
	_reservationHandler "github.com/GroupProject3-Kelompok2/BE/features/reservation/handler"
	_reservationService "github.com/GroupProject3-Kelompok2/BE/features/reservation/service"
	rd "github.com/GroupProject3-Kelompok2/BE/features/review/data"
	rh "github.com/GroupProject3-Kelompok2/BE/features/review/handler"
	rs "github.com/GroupProject3-Kelompok2/BE/features/review/service"
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
	initReservationRouter(db, e)
	initReviewRouter(db, e)
}

func initUserRouter(db *gorm.DB, e *echo.Echo) {
	userData := ud.New(db)
	validate := validator.New()
	userService := us.New(userData, validate)
	userHandler := uh.New(userService)

	e.POST("/register", userHandler.Register())
	e.POST("/login", userHandler.Login())
	e.GET("/users", userHandler.ProfileUser(), middlewares.JWTMiddleware())
	e.PUT("/users", userHandler.UpdateUser(), middlewares.JWTMiddleware())
	e.PUT("/users/role", userHandler.UpgradeUser(), middlewares.JWTMiddleware())
	e.DELETE("/users", userHandler.DeactiveUser(), middlewares.JWTMiddleware())
}

func initHomestayRouter(db *gorm.DB, e *echo.Echo) {
	homestayData := _homestayData.New(db)
	homestayService := _homestayService.New(homestayData)
	homestayHandler := _homestayHandler.New(homestayService)

	e.POST("/homestays", homestayHandler.CreateHomestay(), middlewares.JWTMiddleware())
	e.GET("/homestays", homestayHandler.GetAllHomestay(), middlewares.JWTMiddleware())
	e.GET("/homestays/:homestay_id", homestayHandler.GetHomestayById(), middlewares.JWTMiddleware())
	e.PUT("/homestays/:homestay_id", homestayHandler.UpdateHomestayById(), middlewares.JWTMiddleware())
	e.DELETE("/homestays/:homestay_id", homestayHandler.DeleteHomestayById(), middlewares.JWTMiddleware())
	e.POST("/homestays/:id/pictures", homestayHandler.HomestayPictures(), middlewares.JWTMiddleware())
}

func initReservationRouter(db *gorm.DB, e *echo.Echo) {
	reservationData := _reservationData.New(db)
	reservationService := _reservationService.New(reservationData)
	reservationHandler := _reservationHandler.New(reservationService)

	e.POST("/reservations", reservationHandler.CreateReservation(), middlewares.JWTMiddleware())
}

func initReviewRouter(db *gorm.DB, e *echo.Echo) {
	reviewData := rd.New(db)
	validate := validator.New()
	reviewService := rs.New(reviewData, validate)
	reviewHandler := rh.New(reviewService)

	e.POST("/reviews", reviewHandler.AddReview(), middlewares.JWTMiddleware())
	e.PUT("/reviews/:id", reviewHandler.EditReview(), middlewares.JWTMiddleware())
	e.DELETE("/reviews/:id", reviewHandler.DeleteReview(), middlewares.JWTMiddleware())
}
