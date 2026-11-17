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

// TestCreateNegativePrice проверяет отклонение объявления с отрицательной ценой.
func TestCreateNegativePrice(t *testing.T) {
	f := fixtures.New(t)
	request := fixtures.Advertisement(t)
	request.Fields["price"] = "-1"
	base.Step("Попытаться создать объявление с отрицательной ценой")
	actual := advertisements.CreateAdvertisementError(t, f.Owner.Token, request, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, testdata.PriceRange)
}
