package favorites

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/testdata"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

// TestDeleteAdvertisementRemovesFavorite проверяет каскадное удаление избранного.
func TestDeleteAdvertisementRemovesFavorite(t *testing.T) {
	f := fixtures.New(t)
	buyer := f.Buyer()
	ad := f.Create(fixtures.Advertisement(t))
	favorites.AddFavorite(t, buyer.Token, ad.ID)
	assertions.FavoriteMatchesAdvertisement(t, favorites.GetFavorites(t, buyer.Token, ad.Title), ad.ID)
	base.Step("Владелец удаляет объявление, добавленное в избранное покупателя")
	advertisements.DeleteAdvertisement(t, f.Owner.Token, ad.ID)
	err := advertisements.GetAdvertisementError(t, ad.ID, http.StatusNotFound)
	assertions.ErrorResponse(t, err, testdata.NotFound, testdata.AdvertisementNotFound)
	list := favorites.GetFavorites(t, buyer.Token, ad.Title)
	require.Zero(t, list.Total, "Удалённое объявление учитывается в total избранного")
	require.Empty(t, list.Items, "Удалённое объявление осталось в избранном")
	base.Check("Удаление объявления убрало запись из избранного")
}
