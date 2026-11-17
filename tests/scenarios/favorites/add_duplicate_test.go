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

// TestAddDuplicate проверяет запрет повторного добавления одного объявления.
func TestAddDuplicate(t *testing.T) {
	f := fixtures.New(t)
	buyer := f.Buyer()
	ad := f.Create(fixtures.Advertisement(t))
	favorites.AddFavorite(t, buyer.Token, ad.ID)
	base.Step("Повторно добавить то же объявление в избранное")
	actual := favorites.AddFavoriteError(t, buyer.Token, ad.ID, http.StatusConflict)
	assertions.ErrorResponse(t, actual, testdata.AlreadyExists, testdata.DuplicateFavorite)
	list := favorites.GetFavorites(t, buyer.Token, ad.Title)
	assertions.FavoriteMatchesAdvertisement(t, list, ad.ID)
	base.Check("После повторного запроса в избранном осталась одна запись")
}
