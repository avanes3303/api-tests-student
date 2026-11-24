package advertisementCreate

import (
	"net/http"
	"path/filepath"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/testdata"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
)

// TestCreateUnsupportedPhotoFormat проверяет отклонение текстового файла в поле фотографии.
func TestCreateUnsupportedPhotoFormat(t *testing.T) {
	f := fixtures.New(t)
	request := fixtures.Advertisement(t)
	request.Photos = []string{filepath.Join("testdata", "unsupported_photo.txt")}
	base.Step("Попытаться создать объявление с текстовым файлом вместо фотографии")
	actual := advertisements.CreateAdvertisementError(t, f.Owner.Token, request, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, `photo 1: file type not allowed: extension ".txt"`)
}
