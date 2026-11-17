package advertisements

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"

	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/managers/advertisements/models"
)

// HttpPostAdvertisement отправляет multipart-запрос создания объявления.
func HttpPostAdvertisement(t *testing.T, token string, data models.CreateRequest) *http.Response {
	t.Helper()
	request := apiRunner.GetRunner().Auth(token).Create().Post(path.AdvertisementPath)
	for field, value := range data.Fields {
		request.MultipartFormData(field, value)
	}
	if len(data.Photos) > 0 {
		request.MultipartFile("photos", data.Photos...)
	}
	return request.Expect(t).End().Response
}

// HttpGetAdvertisement получает объявление по query-параметру id из Swagger.
func HttpGetAdvertisement(t *testing.T, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Create().Get(path.AdvertisementPath).
		Query("id", id).Expect(t).End().Response
}

// HttpDeleteAdvertisement удаляет принадлежащее пользователю объявление.
func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Delete(path.AdvertisementPath).
		Query("id", id).Expect(t).End().Response
}

// HttpGetAdvertisementPhotos получает список загруженных фотографий объявления.
func HttpGetAdvertisementPhotos(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().
		Get(fmt.Sprintf("%s/%s/photos", path.AdvertisementsPath, id)).Expect(t).End().Response
}

// HttpGetPhoto получает файл фотографии по URL без передачи пользовательского токена на CDN.
func HttpGetPhoto(t *testing.T, photoURL string) *http.Response {
	t.Helper()
	return apitest.New().EnableNetworking().HttpClient(apiRunner.NewHTTPClient()).
		Get(photoURL).Expect(t).End().Response
}
