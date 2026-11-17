package favorites

import (
	"net/http"
	"testing"

	client "api-tests-template/internal/client/http/favorites"
	"api-tests-template/internal/helpers/response"
	"api-tests-template/internal/managers/favorites/models"
	apiModels "api-tests-template/internal/models"
)

// AddFavorite добавляет чужое объявление в избранное пользователя.
func AddFavorite(t *testing.T, token, advertisementID string) models.Favorite {
	t.Helper()
	return response.JSON[models.Favorite](t, client.HttpPostFavorite(t, token, models.AddRequest{AdvertisementID: advertisementID}), http.StatusCreated)
}

// AddFavoriteError проверяет ошибку добавления объявления в избранное.
func AddFavoriteError(t *testing.T, token, advertisementID string, status int) apiModels.ErrorResponse {
	t.Helper()
	return response.JSON[apiModels.ErrorResponse](t, client.HttpPostFavorite(t, token, models.AddRequest{AdvertisementID: advertisementID}), status)
}

// GetFavorites возвращает страницу избранного с поиском по названию.
func GetFavorites(t *testing.T, token, search string) models.ListResponse {
	t.Helper()
	return response.JSON[models.ListResponse](t, client.HttpGetFavorites(t, token, search), http.StatusOK)
}

// GetFavoritesError проверяет отказ в доступе к избранному.
func GetFavoritesError(t *testing.T, token string, status int) apiModels.ErrorResponse {
	t.Helper()
	return response.JSON[apiModels.ErrorResponse](t, client.HttpGetFavorites(t, token, ""), status)
}

// DeleteFavorite удаляет объявление из избранного и проверяет HTTP 204.
func DeleteFavorite(t *testing.T, token, advertisementID string) {
	t.Helper()
	response.Read(t, client.HttpDeleteFavorite(t, token, advertisementID), http.StatusNoContent)
}

// DeleteFavoriteError проверяет ошибку удаления записи избранного.
func DeleteFavoriteError(t *testing.T, token, advertisementID string, status int) apiModels.ErrorResponse {
	t.Helper()
	return response.JSON[apiModels.ErrorResponse](t, client.HttpDeleteFavorite(t, token, advertisementID), status)
}
