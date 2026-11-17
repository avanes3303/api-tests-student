package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/favorites/models"
)

// FavoriteIDs возвращает идентификаторы объявлений из ответа избранного.
func FavoriteIDs(list models.ListResponse) []string {
	ids := make([]string, 0, len(list.Items))
	for _, item := range list.Items {
		ids = append(ids, item.Advertisement.ID)
	}
	return ids
}

// FavoriteMatchesAdvertisement проверяет появление единственной записи по уникальному названию.
func FavoriteMatchesAdvertisement(t *testing.T, list models.ListResponse, advertisementID string) {
	t.Helper()
	require.Equal(t, 1, list.Total, "По уникальному названию ожидалась одна запись избранного")
	require.Equal(t, []string{advertisementID}, FavoriteIDs(list), "В избранном отсутствует добавленное объявление")
	require.NotEmpty(t, list.Items[0].ID, "Запись избранного не содержит id")
	require.NotEmpty(t, list.Items[0].AddedAt, "Запись избранного не содержит дату добавления")
}
