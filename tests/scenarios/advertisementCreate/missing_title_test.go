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

// TestCreateMissingTitle проверяет отклонение объявления без поля title.
func TestCreateMissingTitle(t *testing.T) {
	f := fixtures.New(t)
	request := fixtures.Advertisement(t)
	delete(request.Fields, "title")
	base.Step("Попытаться создать объявление без поля title")
	actual := advertisements.CreateAdvertisementError(t, f.Owner.Token, request, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, testdata.TitleRequired)
}
