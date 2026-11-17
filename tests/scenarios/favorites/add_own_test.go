package favorites

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/testdata"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

// TestAddOwnAdvertisement проверяет запрет добавления собственного объявления.
func TestAddOwnAdvertisement(t *testing.T) {
	f := fixtures.New(t)
	ad := f.Create(fixtures.Advertisement(t))
	base.Step("Добавить собственное объявление в избранное")
	actual := favorites.AddFavoriteError(t, f.Owner.Token, ad.ID, http.StatusForbidden)
	assertions.ErrorResponse(t, actual, testdata.Forbidden, testdata.OwnFavorite)
	list := favorites.GetFavorites(t, f.Owner.Token, ad.Title)
	require.Empty(t, list.Items, "Своё объявление появилось в избранном после отказа")
	require.Zero(t, list.Total, "Своё объявление учитывается в total избранного")
	base.Check("Собственное объявление не добавлено")
}
