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

// TestCreateInvalidQuantity проверяет отклонение объявления с нечисловым количеством.
func TestCreateInvalidQuantity(t *testing.T) {
	f := fixtures.New(t)
	request := fixtures.Advertisement(t)
	request.Fields["quantity"] = "abc"
	base.Step("Попытаться создать объявление с нечисловым количеством")
	actual := advertisements.CreateAdvertisementError(t, f.Owner.Token, request, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, "cannot convert string value to int")
}
