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

// TestAddWithoutID проверяет корректный отказ для отсутствующего или некорректного id.
func TestAddWithoutID(t *testing.T) {
	f := fixtures.New(t)
	base.Step("Добавить в избранное без advertisement_id")
	actual := favorites.AddFavoriteError(t, f.Owner.Token, "", http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, testdata.FavoriteIDRequired)
}
