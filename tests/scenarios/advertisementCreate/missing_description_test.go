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

// TestCreateMissingDescription проверяет отклонение объявления без поля description.
func TestCreateMissingDescription(t *testing.T) {
	f := fixtures.New(t)
	request := fixtures.Advertisement(t)
	delete(request.Fields, "description")
	base.Step("Попытаться создать объявление без поля description")
	actual := advertisements.CreateAdvertisementError(t, f.Owner.Token, request, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, testdata.DescriptionRequired)
}
