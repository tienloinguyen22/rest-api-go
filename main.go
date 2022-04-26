package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/tienloinguyen22/edwork-api-go/adapters"
	"github.com/tienloinguyen22/edwork-api-go/configs"
	"github.com/tienloinguyen22/edwork-api-go/core/auth"
	"github.com/tienloinguyen22/edwork-api-go/core/healthcheck"
	"github.com/tienloinguyen22/edwork-api-go/core/profiles"
	"github.com/tienloinguyen22/edwork-api-go/core/users"
)

func main() {
	// Prerequisites
	cfg := configs.InitializeConfigs()
	firebaseAdmin := adapters.InitializeFirebaseAdmin(cfg.FIREBASE_CREDENTIALS_FILE)
	db := adapters.InitializePostgresql(cfg.DB_URI)

	// Repositories
	userRepo := users.NewUserRepository(db)

	// Service
	authService := auth.NewAuthService(firebaseAdmin, userRepo)
	profileService := profiles.NewProfileService(userRepo)

	// Controller
	r := gin.Default()
	healthcheck.NewHealthcheckController(r)
	auth.NewAuthController(r, firebaseAdmin, userRepo, authService)
	profiles.NewProfileController(r, firebaseAdmin, userRepo, profileService)

	// Start app
	r.Run(cfg.ADDRESS)
}