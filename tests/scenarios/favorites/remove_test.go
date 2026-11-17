package favorites

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

// TestRemoveFavorite проверяет удаление записи с сохранением самого объявления.
func TestRemoveFavorite(t *testing.T) {
	f := fixtures.New(t)
	buyer := f.Buyer()
	ad := f.Create(fixtures.Advertisement(t))
	favorites.AddFavorite(t, buyer.Token, ad.ID)
	assertions.FavoriteMatchesAdvertisement(t, favorites.GetFavorites(t, buyer.Token, ad.Title), ad.ID)
	base.Step("Удалить объявление из избранного")
	favorites.DeleteFavorite(t, buyer.Token, ad.ID)
	list := favorites.GetFavorites(t, buyer.Token, ad.Title)
	require.Zero(t, list.Total, "Удалённое избранное учитывается в total")
	require.Empty(t, list.Items, "Удалённое избранное осталось в списке")
	require.Equal(t, ad.ID, advertisements.GetAdvertisement(t, ad.ID).ID, "Само объявление должно быть доступно")
	base.Check("Запись исчезла из избранного, объявление доступно")
}
