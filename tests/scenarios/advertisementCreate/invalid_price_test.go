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

// TestCreateInvalidPrice проверяет отклонение объявления с нечисловой ценой.
func TestCreateInvalidPrice(t *testing.T) {
	f := fixtures.New(t)
	request := fixtures.Advertisement(t)
	request.Fields["price"] = "abc"
	base.Step("Попытаться создать объявление с нечисловой ценой")
	actual := advertisements.CreateAdvertisementError(t, f.Owner.Token, request, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, "cannot convert string value abc to float")
}
