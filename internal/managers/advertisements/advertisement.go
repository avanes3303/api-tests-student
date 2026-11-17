package advertisements

import (
	"bytes"
	"encoding/json"
	"image"
	"image/png"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	client "api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/helpers/response"
	"api-tests-template/internal/managers/advertisements/models"
	apiModels "api-tests-template/internal/models"
)

// CreateAdvertisement создаёт объявление и возвращает все сохранённые параметры.
func CreateAdvertisement(t *testing.T, token string, data models.CreateRequest) models.Advertisement {
	t.Helper()
	return response.JSON[models.Advertisement](t, client.HttpPostAdvertisement(t, token, data), http.StatusCreated)
}

// CreateAdvertisementError проверяет отклонение невалидного запроса создания.
func CreateAdvertisementError(t *testing.T, token string, data models.CreateRequest, status int) apiModels.ErrorResponse {
	t.Helper()
	raw := client.HttpPostAdvertisement(t, token, data)
	require.NotNil(t, raw, "API создания объявления не вернул HTTP-ответ")
	// Даже при ошибочном принятии запроса удаляем неожиданно созданное объявление.
	if raw.StatusCode == http.StatusCreated {
		body := response.Read(t, raw, http.StatusCreated)
		var created models.Advertisement
		require.NoError(t, json.Unmarshal(body, &created), "Не удалось разобрать неожиданно созданное объявление")
		if created.ID != "" {
			t.Cleanup(func() { DeleteAdvertisement(t, token, created.ID) })
		}
		t.Fatalf("Ожидалось отклонение запроса с HTTP %d, но объявление создано с HTTP 201", status)
	}
	return response.JSON[apiModels.ErrorResponse](t, raw, status)
}

// GetAdvertisement возвращает сохранённое объявление по его идентификатору.
func GetAdvertisement(t *testing.T, id string) models.Advertisement {
	t.Helper()
	return response.JSON[models.Advertisement](t, client.HttpGetAdvertisement(t, id), http.StatusOK)
}

// GetAdvertisementError проверяет ошибку получения отсутствующего объявления.
func GetAdvertisementError(t *testing.T, id string, status int) apiModels.ErrorResponse {
	t.Helper()
	return response.JSON[apiModels.ErrorResponse](t, client.HttpGetAdvertisement(t, id), status)
}

// GetPhotos возвращает фотографии, доступные через API объявления.
func GetPhotos(t *testing.T, token, id string) []models.Photo {
	t.Helper()
	return response.JSON[[]models.Photo](t, client.HttpGetAdvertisementPhotos(t, token, id), http.StatusOK)
}

// GetPhotoSize скачивает PNG-фотографию и проверяет, что сервер возвращает изображение.
func GetPhotoSize(t *testing.T, photoURL string) image.Config {
	t.Helper()
	body := response.Read(t, client.HttpGetPhoto(t, photoURL), http.StatusOK)
	config, err := png.DecodeConfig(bytes.NewReader(body))
	require.NoError(t, err, "По URL фотографии получен некорректный PNG-файл")
	return config
}

// DeleteAdvertisement удаляет объявление и проверяет HTTP 204.
func DeleteAdvertisement(t *testing.T, token, id string) {
	t.Helper()
	response.Read(t, client.HttpDeleteAdvertisement(t, token, id), http.StatusNoContent)
}

// CleanupAdvertisement удаляет тестовые данные, допуская уже удалённое объявление.
func CleanupAdvertisement(t *testing.T, token, id string) {
	t.Helper()
	raw := client.HttpDeleteAdvertisement(t, token, id)
	require.NotNil(t, raw, "API удаления объявления не вернул HTTP-ответ")
	if raw.StatusCode == http.StatusNotFound {
		response.Read(t, raw, http.StatusNotFound)
		return
	}
	if raw.StatusCode == http.StatusForbidden {
		// Стенд возвращает 403 при повторном DELETE уже удалённого объявления.
		// Дополнительно подтверждаем отсутствие через публичный GET.
		response.Read(t, raw, http.StatusForbidden)
		GetAdvertisementError(t, id, http.StatusNotFound)
		return
	}
	response.Read(t, raw, http.StatusNoContent)
}
