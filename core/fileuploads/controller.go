package fileuploads

import (
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/tienloinguyen22/edwork-api-go/core/users"
	"github.com/tienloinguyen22/edwork-api-go/middlewares"
	"github.com/tienloinguyen22/edwork-api-go/utils"
)

type FileUploadController struct {
	FirebaseAdmin *firebase.App
	UserRepo *users.UserRepository
	FileUploadService *FileUploadService
}

func NewFileUploadController(router *gin.Engine, firebaseAdmin *firebase.App, userRepo *users.UserRepository, fileUploadService *FileUploadService) {
	fileUploadController := &FileUploadController{
		FirebaseAdmin: firebaseAdmin,
		UserRepo: userRepo,
		FileUploadService: fileUploadService,
	}
	fileUploadController.SetupRouter(router)
}

func (c FileUploadController) SetupRouter(router *gin.Engine) {
	controller := router.Group("/api/file-uploads")

	controller.POST(
		"/image",
		middlewares.VerifyToken(c.FirebaseAdmin, c.UserRepo),
		middlewares.Authenticated(),
		func (ctx *gin.Context) {
			uploadedFile, err := c.FileUploadService.UploadImage(ctx)
			if err != nil {
				utils.HandleError(ctx, err)
				return
			}

			ctx.JSON(200, uploadedFile)
		},
	)
}