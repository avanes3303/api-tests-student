package advertisementCreate

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/testdata"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
)

// TestCreateTooManyPhotos проверяет отклонение объявления с четырьмя фотографиями.
func TestCreateTooManyPhotos(t *testing.T) {
	f := fixtures.New(t)
	request := fixtures.Advertisement(t)
	request.Photos = fixtures.Photos(t, testdata.PhotoLimit+1)
	base.Step("Попытаться создать объявление с четырьмя фотографиями")
	actual := advertisements.CreateAdvertisementError(t, f.Owner.Token, request, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, "maximum 3 photos allowed")
}
