package main

import (
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/tienloinguyen22/edwork-api-go/adapters"
	"github.com/tienloinguyen22/edwork-api-go/configs"
	"github.com/tienloinguyen22/edwork-api-go/core/auth"
	"github.com/tienloinguyen22/edwork-api-go/core/consumers"
	"github.com/tienloinguyen22/edwork-api-go/core/fileuploads"
	"github.com/tienloinguyen22/edwork-api-go/core/healthcheck"
	"github.com/tienloinguyen22/edwork-api-go/core/profiles"
	"github.com/tienloinguyen22/edwork-api-go/core/resetpassword"
	"github.com/tienloinguyen22/edwork-api-go/core/users"
)

func main() {
	// Prerequisites
	cfg := configs.InitializeConfigs()
	firebaseAdmin := adapters.InitializeFirebaseAdmin(cfg.FIREBASE_CREDENTIALS_FILE)
	db := adapters.InitializePostgresql(cfg.DB_URI)
	mq := adapters.InitializeMessageQueue(cfg.REDIS_URI)
	emailClient := adapters.InitializeSmtpClient(cfg.EMAIL_HOST, cfg.EMAIL_PORT, cfg.EMAIL_USERNAME, cfg.EMAIL_PASSWORD)

	// Repositories
	userRepo := users.NewUserRepository(db)
	resetPasswordTokenRepo := resetpassword.NewResetPasswordTokenRepository(db)

	// Service
	authService := auth.NewAuthService(firebaseAdmin, userRepo)
	profileService := profiles.NewProfileService(mq, userRepo)
	fileUploadService := fileuploads.NewFileUploadService()
	consumerService := consumers.NewConsumerService()
	resetPasswordService := resetpassword.NewResetPasswordService(firebaseAdmin, emailClient, resetPasswordTokenRepo, userRepo)

	// Message queue
	mq.Consume(adapters.ConsumerConfig{
		PrefetchCount: 10,
		PollInterval: time.Second,
		QueueName: "RESIZE_IMAGE",
		Callback: consumerService.ResizeImage,
	})

	// Controller
	r := gin.Default()
	r.Use(static.Serve("/images", static.LocalFile("./uploads", false)))
	r.Use(static.Serve("/images", static.LocalFile("./temp", false)))
	healthcheck.NewHealthcheckController(r)
	auth.NewAuthController(r, firebaseAdmin, userRepo, authService)
	profiles.NewProfileController(r, firebaseAdmin, userRepo, profileService)
	fileuploads.NewFileUploadController(r, firebaseAdmin, userRepo, fileUploadService)
	resetpassword.NewResetPasswordController(r, resetPasswordService)

	// Start app
	r.Run(cfg.ADDRESS)
}