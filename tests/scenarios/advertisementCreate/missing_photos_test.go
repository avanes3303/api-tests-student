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

// TestCreateMissingPhotos проверяет отклонение объявления без фотографий.
func TestCreateMissingPhotos(t *testing.T) {
	f := fixtures.New(t)
	request := fixtures.Advertisement(t)
	request.Photos = nil
	base.Step("Попытаться создать объявление без фотографий")
	actual := advertisements.CreateAdvertisementError(t, f.Owner.Token, request, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, testdata.PhotoRequired)
}
