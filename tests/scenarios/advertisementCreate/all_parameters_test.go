package advertisementCreate

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/testdata"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
)

// TestCreateWithAllParameters проверяет создание, чтение, поиск и доступность всех фотографий.
func TestCreateWithAllParameters(t *testing.T) {
	f := fixtures.New(t)
	request := fixtures.Advertisement(t)
	request.Photos = fixtures.Photos(t, testdata.PhotoLimit)
	base.Step("Создать объявление со всеми полями и тремя фотографиями")
	created := f.Create(request)
	assertions.AdvertisementMatchesRequest(t, created, request, f.Owner.User.Id)

	base.Step("Получить объявление по id и сверить с исходным запросом")
	actual := advertisements.GetAdvertisement(t, created.ID)
	require.Equal(t, created.ID, actual.ID, "GET вернул другое объявление")
	assertions.AdvertisementMatchesRequest(t, actual, request, f.Owner.User.Id)
	base.Check("Все параметры сохранены без изменений")

	base.Step("Найти объявление по полному уникальному названию")
	list := advertisements.GetAdvertisementsList(t, f.Owner.Token, request.Fields["title"])
	require.Len(t, list.Items, 1, "Поиск по полному уникальному названию должен вернуть одно объявление")
	require.Equal(t, created.ID, list.Items[0].ID, "Поиск вернул другое объявление")
	base.Check("Созданное объявление найдено по полному названию")

	base.Step("Получить список фотографий и скачать каждую фотографию")
	photos := advertisements.GetPhotos(t, f.Owner.Token, created.ID)
	require.ElementsMatch(t, actual.Photos, photos, "Список фотографий не соответствует карточке объявления")
	for _, photo := range photos {
		config := advertisements.GetPhotoSize(t, photo.URL)
		require.Equal(t, testdata.PhotoWidth, config.Width, "Ширина загруженной фотографии изменилась")
		require.Equal(t, testdata.PhotoHeight, config.Height, "Высота загруженной фотографии изменилась")
	}
	base.Check("Все три фотографии загружены и доступны по HTTP")
}
