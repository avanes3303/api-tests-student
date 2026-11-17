package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/testdata"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

// TestAddInvalidID проверяет корректный отказ для отсутствующего или некорректного id.
func TestAddInvalidID(t *testing.T) {
	f := fixtures.New(t)
	base.Step("Добавить в избранное по невалидному UUID")
	actual := favorites.AddFavoriteError(t, f.Owner.Token, testdata.InvalidID, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, testdata.InvalidUUID)
}
