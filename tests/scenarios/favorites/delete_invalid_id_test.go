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

// TestDeleteInvalidID проверяет корректный отказ для отсутствующего или некорректного id.
func TestDeleteInvalidID(t *testing.T) {
	f := fixtures.New(t)
	base.Step("Удалить из избранного по невалидному UUID")
	actual := favorites.DeleteFavoriteError(t, f.Owner.Token, testdata.InvalidID, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, "Invalid advertisement ID")
}
