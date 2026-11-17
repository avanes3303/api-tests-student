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

// TestCreateQuantityAboveLimit проверяет отклонение объявления с количеством выше максимального.
func TestCreateQuantityAboveLimit(t *testing.T) {
	f := fixtures.New(t)
	request := fixtures.Advertisement(t)
	request.Fields["quantity"] = "101"
	base.Step("Попытаться создать объявление с количеством выше максимального")
	actual := advertisements.CreateAdvertisementError(t, f.Owner.Token, request, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, testdata.QuantityRange)
}
