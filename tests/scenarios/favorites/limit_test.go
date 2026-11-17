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

// TestFavoritesLimit проверяет допустимые 50 записей и отказ при добавлении 51-й.
func TestFavoritesLimit(t *testing.T) {
	f := fixtures.New(t)
	buyer := f.Buyer()
	initial := favorites.GetFavorites(t, buyer.Token, "")
	require.Zero(t, initial.Total, "Технический покупатель должен начинать сценарий с пустым избранным")
	request := fixtures.Advertisement(t)
	expectedIDs := make([]string, 0, testdata.FavoritesLimit)
	base.Step("Добавить 50 разных чужих объявлений в избранное")
	for i := 0; i < testdata.FavoritesLimit; i++ {
		request.Fields["title"] = fixtures.AdvertisementTitle()
		ad := f.Create(request)
		favorites.AddFavorite(t, buyer.Token, ad.ID)
		expectedIDs = append(expectedIDs, ad.ID)
	}
	full := favorites.GetFavorites(t, buyer.Token, "")
	require.Equal(t, testdata.FavoritesLimit, full.Total, "Лимит из 50 записей должен быть доступен")
	require.ElementsMatch(t, expectedIDs, assertions.FavoriteIDs(full), "До достижения лимита потеряны записи избранного")
	base.Step("Попытаться добавить 51-е объявление")
	request.Fields["title"] = fixtures.AdvertisementTitle()
	extra := f.Create(request)
	err := favorites.AddFavoriteError(t, buyer.Token, extra.ID, http.StatusBadRequest)
	assertions.ErrorResponse(t, err, testdata.LimitReached, testdata.FavoritesLimitReached)
	after := favorites.GetFavorites(t, buyer.Token, "")
	require.Equal(t, testdata.FavoritesLimit, after.Total, "После отказа количество избранного изменилось")
	require.ElementsMatch(t, expectedIDs, assertions.FavoriteIDs(after), "После отказа состав избранного изменился")
	base.Check("50 записей доступны, 51-я отклонена, исходные записи сохранены")
}
