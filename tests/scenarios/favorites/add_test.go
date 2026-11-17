package favorites

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

// TestAddFavorite проверяет появление чужого объявления в избранном покупателя.
func TestAddFavorite(t *testing.T) {
	f := fixtures.New(t)
	buyer := f.Buyer()
	ad := f.Create(fixtures.Advertisement(t))
	base.Step("Добавить чужое объявление в избранное")
	added := favorites.AddFavorite(t, buyer.Token, ad.ID)
	require.Equal(t, ad.ID, added.Advertisement.ID, "Добавлено другое объявление")
	base.Step("Получить избранное по уникальному названию")
	list := favorites.GetFavorites(t, buyer.Token, ad.Title)
	assertions.FavoriteMatchesAdvertisement(t, list, ad.ID)
	base.Check("Добавленное объявление присутствует в избранном")
}
