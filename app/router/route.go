package router

import (
	_homestayData "github.com/GroupProject3-Kelompok2/BE/features/homestay/data"
	_homestayHandler "github.com/GroupProject3-Kelompok2/BE/features/homestay/handler"
	_homestayService "github.com/GroupProject3-Kelompok2/BE/features/homestay/service"
	_paymentdata "github.com/GroupProject3-Kelompok2/BE/features/payment/data"
	_paymenthandler "github.com/GroupProject3-Kelompok2/BE/features/payment/handler"
	_paymentservice "github.com/GroupProject3-Kelompok2/BE/features/payment/service"
	_reservationData "github.com/GroupProject3-Kelompok2/BE/features/reservation/data"
	_reservationHandler "github.com/GroupProject3-Kelompok2/BE/features/reservation/handler"
	_reservationService "github.com/GroupProject3-Kelompok2/BE/features/reservation/service"
	_reviewdata "github.com/GroupProject3-Kelompok2/BE/features/review/data"
	_reviewhandler "github.com/GroupProject3-Kelompok2/BE/features/review/handler"
	_reviewservice "github.com/GroupProject3-Kelompok2/BE/features/review/service"
	_userdata "github.com/GroupProject3-Kelompok2/BE/features/user/data"
	_userhandler "github.com/GroupProject3-Kelompok2/BE/features/user/handler"
	_userservice "github.com/GroupProject3-Kelompok2/BE/features/user/service"
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
	initPaymentRouter(db, e)
}

func initUserRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userdata.New(db)
	validate := validator.New()
	userService := _userservice.New(userData, validate)
	userHandler := _userhandler.New(userService)

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
	e.GET("/users/homestays", homestayHandler.GetAllHomestayByUserId(), middlewares.JWTMiddleware())
}

func initReservationRouter(db *gorm.DB, e *echo.Echo) {
	reservationData := _reservationData.New(db)
	reservationService := _reservationService.New(reservationData)
	reservationHandler := _reservationHandler.New(reservationService)

	e.POST("/reservations", reservationHandler.CreateReservation(), middlewares.JWTMiddleware())
	e.POST("/reservations/availability", reservationHandler.CheckAvailability(), middlewares.JWTMiddleware())
	e.GET("/reservations/:reservation_id", reservationHandler.GetReservationById(), middlewares.JWTMiddleware())
	e.GET("/users/reservations", reservationHandler.GetAllReservationsByUserId(), middlewares.JWTMiddleware())
}

func initReviewRouter(db *gorm.DB, e *echo.Echo) {
	reviewData := _reviewdata.New(db)
	validate := validator.New()
	reviewService := _reviewservice.New(reviewData, validate)
	reviewHandler := _reviewhandler.New(reviewService)

	e.POST("/reviews", reviewHandler.AddReview(), middlewares.JWTMiddleware())
	e.PUT("/reviews/:id", reviewHandler.EditReview(), middlewares.JWTMiddleware())
	e.DELETE("/reviews/:id", reviewHandler.DeleteReview(), middlewares.JWTMiddleware())
}

func initPaymentRouter(db *gorm.DB, e *echo.Echo) {
	paymentData := _paymentdata.New(db)
	validate := validator.New()
	paymentService := _paymentservice.New(paymentData, validate)
	paymentHandler := _paymenthandler.New(paymentService)

	e.POST("/payments", paymentHandler.Payment(), middlewares.JWTMiddleware())
	e.POST("/payments/callback", paymentHandler.Notification())
}
