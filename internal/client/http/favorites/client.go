package favorites

import (
	"net/http"
	"strconv"
	"testing"

	"api-tests-template/internal/constants/path"
	"api-tests-template/internal/constants/testdata"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/managers/favorites/models"
)

// HttpPostFavorite добавляет объявление в избранное текущего пользователя.
func HttpPostFavorite(t *testing.T, token string, data models.AddRequest) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Post(path.FavoritesPath).
		JSON(data).Expect(t).End().Response
}

// HttpGetFavorites получает избранное с необязательным поиском по названию.
func HttpGetFavorites(t *testing.T, token, search string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Get(path.FavoritesPath).
		Query("page", "1").Query("page_size", strconv.Itoa(testdata.FavoritesLimit)).
		Query("search", search).Expect(t).End().Response
}

// HttpDeleteFavorite удаляет запись избранного по идентификатору объявления.
func HttpDeleteFavorite(t *testing.T, token, advertisementID string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Delete(path.FavoritesPath + "/" + advertisementID).
		Expect(t).End().Response
}
