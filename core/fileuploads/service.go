package fileuploads

import (
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tienloinguyen22/edwork-api-go/utils"
)

type FileUploadService struct {}

func NewFileUploadService() *FileUploadService {
	return &FileUploadService{}
}

func (s FileUploadService) UploadImage(ctx *gin.Context) (*UploadFileResult, error) {
	file, err := ctx.FormFile("image")
	if err != nil {
		return nil, utils.NewApiError(http.StatusBadRequest, "fileuploads.upload-image.cant-receive-file", err)
	}

	// Validate file ext
	ext := filepath.Ext(file.Filename)
	matched, err := regexp.MatchString(utils.REGEX_IMAGE_EXT, ext)
	if err != nil {
		return nil, utils.NewApiError(http.StatusBadRequest, "fileuploads.upload-image.cant-detect-file-ext", err)
	}
	if !matched {
		return nil, utils.NewApiError(http.StatusBadRequest, "fileuploads.upload-image.file-ext-not-allowed", fmt.Errorf("invalid image file"),)
	}

	// Validate file size
	if file.Size > int64(utils.ALLOW_FILE_SIZE_IMAGE) {
		return nil, utils.NewApiError(
			http.StatusBadRequest,
			"fileuploads.upload-image.file-too-large",
			fmt.Errorf("file too large (Max 5MB)"),
		)
	}

	// Save file to /temp folder
	err = utils.EnsureFolderExist("./temp")
	if err != nil {
		return nil, utils.NewApiError(
			http.StatusInternalServerError,
			"fileuploads.upload-image.cant-find-upload-folder",
			err,
		)
	}

	filename := fmt.Sprintf("%v%v", uuid.New().String(), ext)
	filepath := fmt.Sprintf("./temp/%v", filename)
	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		return nil, utils.NewApiError(
			http.StatusInternalServerError,
			"fileuploads.upload-image.cant-save-file",
			err,
		)
	}

	return &UploadFileResult{
		Filepath: filename,
	}, nil
}