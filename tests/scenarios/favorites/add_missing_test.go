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

// TestAddMissingAdvertisement проверяет корректный отказ для отсутствующего или некорректного id.
func TestAddMissingAdvertisement(t *testing.T) {
	f := fixtures.New(t)
	base.Step("Добавить несуществующее объявление")
	actual := favorites.AddFavoriteError(t, f.Owner.Token, testdata.MissingAdvertisementID, http.StatusNotFound)
	assertions.ErrorResponse(t, actual, testdata.NotFound, testdata.AdvertisementNotFound)
}
