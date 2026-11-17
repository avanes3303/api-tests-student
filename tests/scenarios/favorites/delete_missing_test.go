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

// TestDeleteMissingAdvertisement проверяет корректный отказ для отсутствующего или некорректного id.
func TestDeleteMissingAdvertisement(t *testing.T) {
	f := fixtures.New(t)
	base.Step("Удалить несуществующее объявление из избранного")
	actual := favorites.DeleteFavoriteError(t, f.Owner.Token, testdata.MissingAdvertisementID, http.StatusNotFound)
	assertions.ErrorResponse(t, actual, testdata.NotFound, testdata.FavoriteNotFound)
}
