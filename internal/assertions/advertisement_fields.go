package assertions

import (
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisements/models"
)

// AdvertisementMatchesRequest сверяет сохранённые параметры с исходным multipart-запросом.
func AdvertisementMatchesRequest(t *testing.T, actual models.Advertisement, request models.CreateRequest, ownerID string) {
	t.Helper()
	price, err := strconv.ParseFloat(request.Fields["price"], 64)
	require.NoError(t, err, "Некорректная цена в позитивной фикстуре")
	quantity, err := strconv.Atoi(request.Fields["quantity"])
	require.NoError(t, err, "Некорректное количество в позитивной фикстуре")
	require.NotEmpty(t, actual.ID, "Объявление не содержит id")
	require.Equal(t, request.Fields["title"], actual.Title, "Название не соответствует запросу")
	require.Equal(t, request.Fields["description"], actual.Description, "Описание не соответствует запросу")
	require.Equal(t, price, actual.Price, "Цена не соответствует запросу")
	require.Equal(t, quantity, actual.Quantity, "Количество не соответствует запросу")
	require.Equal(t, ownerID, actual.UserID, "Некорректный владелец объявления")
	require.NotEmpty(t, actual.CreatedAt, "Дата создания отсутствует")
	require.NotEmpty(t, actual.UpdatedAt, "Дата обновления отсутствует")
	require.Len(t, actual.Photos, len(request.Photos), "Не все фотографии сохранены")
	for _, photo := range actual.Photos {
		require.NotEmpty(t, photo.ID, "Фотография не содержит id")
		parsed, parseErr := url.ParseRequestURI(photo.URL)
		require.NoError(t, parseErr, "Фотография содержит некорректный URL")
		require.NotEmpty(t, parsed.Host, "В URL фотографии отсутствует сервер")
	}
}
