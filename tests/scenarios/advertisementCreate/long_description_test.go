package advertisementCreate

import (
	"net/http"
	"strings"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/testdata"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
)

// TestCreateLongDescription проверяет отклонение объявления с описанием длиннее 500 символов.
func TestCreateLongDescription(t *testing.T) {
	f := fixtures.New(t)
	request := fixtures.Advertisement(t)
	request.Fields["description"] = strings.Repeat("Я", testdata.DescriptionLimit+1)
	base.Step("Попытаться создать объявление с описанием длиннее 500 символов")
	actual := advertisements.CreateAdvertisementError(t, f.Owner.Token, request, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, "description must be between 1 and 500 characters")
}
