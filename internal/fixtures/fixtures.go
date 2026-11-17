package fixtures

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/constants/testdata"
	"api-tests-template/internal/managers/advertisements"
	adModels "api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/managers/auth"
	authModels "api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

var ownerSession struct {
	sync.Mutex
	user authModels.LoginOkResponse
}

// Context содержит владельца данных и автоматически очищаемые ресурсы одного теста.
type Context struct {
	T     *testing.T
	Owner authModels.LoginOkResponse
}

// Init загружает настройки API без обязательной авторизации для гостевых сценариев.
func Init(t *testing.T) {
	t.Helper()
	base.SetupSuite()
	require.NotEmpty(t, os.Getenv("API_URL"), "Задайте API_URL в .env или .env.override")
}

// New авторизует владельца из локальных переменных окружения.
func New(t *testing.T) *Context {
	t.Helper()
	Init(t)
	ownerSession.Lock()
	defer ownerSession.Unlock()
	base.Precondition("Авторизация владельца тестовых объявлений")
	if ownerSession.user.Token == "" {
		ownerSession.user = auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
	}
	return WithOwner(t, ownerSession.user)
}

// WithOwner создаёт фикстуру на основе уже авторизованного владельца.
func WithOwner(t *testing.T, owner authModels.LoginOkResponse) *Context {
	t.Helper()
	return &Context{T: t, Owner: owner}
}

// Buyer создаёт или переиспользует отдельного покупателя, привязанного к учётной записи и сценарию.
func (f *Context) Buyer() authModels.LoginOkResponse {
	f.T.Helper()
	base.Precondition("Подготовка изолированного покупателя для сценария")
	seed := sha256.Sum256([]byte(os.Getenv("API_URL") + "|" + f.Owner.User.Id + "|" + os.Getenv("TEST_PASSWORD") + "|" + f.T.Name()))
	value := hex.EncodeToString(seed[:])
	user := auth.EnsureTestUser(f.T, authModels.RegisterRequest{
		Email: "qa-api-lab-" + value[:24] + "@example.com", Password: "Qa1" + value,
		FirstName: "Покупатель", LastName: "Тестовый",
	})
	require.NotEqual(f.T, f.Owner.User.Id, user.User.Id, "Покупатель и владелец должны быть разными пользователями")
	return user
}

// Create создаёт объявление и регистрирует его удаление даже после падения проверки.
func (f *Context) Create(data adModels.CreateRequest) adModels.Advertisement {
	f.T.Helper()
	base.Precondition("Создание объявления с уникальным названием")
	ad := advertisements.CreateAdvertisement(f.T, f.Owner.Token, data)
	require.NotEmpty(f.T, ad.ID, "Создание объявления не вернуло id")
	f.T.Cleanup(func() {
		base.Step("Очистка тестового объявления " + ad.ID)
		advertisements.CleanupAdvertisement(f.T, f.Owner.Token, ad.ID)
	})
	return ad
}

// Advertisement возвращает новый валидный набор multipart-полей и одну PNG-фотографию.
func Advertisement(t *testing.T) adModels.CreateRequest {
	t.Helper()
	return adModels.CreateRequest{
		Fields: map[string]string{
			"title": AdvertisementTitle(), "description": testdata.AdvertisementText,
			"price": testdata.AdvertisementPrice, "quantity": testdata.AdvertisementQuantity,
		},
		Photos: Photos(t, 1),
	}
}

// AdvertisementTitle возвращает уникальное название для изоляции тестовых объявлений.
func AdvertisementTitle() string {
	return "API lab " + utils.RandomString(16)
}

// Photos создаёт маленькие валидные PNG-файлы в автоматически удаляемой временной папке.
func Photos(t *testing.T, count int) []string {
	t.Helper()
	directory := t.TempDir()
	files := make([]string, 0, count)
	colors := []color.RGBA{{R: 200, G: 100, A: 255}, {G: 200, B: 100, A: 255}, {R: 100, B: 200, A: 255}}
	for i := 0; i < count; i++ {
		photo := image.NewRGBA(image.Rect(0, 0, testdata.PhotoWidth, testdata.PhotoHeight))
		draw.Draw(photo, photo.Bounds(), &image.Uniform{C: colors[i%len(colors)]}, image.Point{}, draw.Src)
		var buffer bytes.Buffer
		require.NoError(t, png.Encode(&buffer, photo), "Не удалось подготовить тестовую фотографию")
		file := filepath.Join(directory, fmt.Sprintf("photo-%d.png", i))
		require.NoError(t, os.WriteFile(file, buffer.Bytes(), 0o600), "Не удалось сохранить тестовую фотографию")
		files = append(files, file)
	}
	return files
}

// AuthorizationCase описывает отсутствие токена или неверный токен и ожидаемое сообщение.
type AuthorizationCase struct {
	Name    string
	Token   string
	Message string
}

// UnauthorizedCases возвращает общие варианты проверки обязательной авторизации.
func UnauthorizedCases() []AuthorizationCase {
	return []AuthorizationCase{
		{Name: "без токена", Message: testdata.AuthorizationRequired},
		{Name: "неверный токен", Token: testdata.InvalidToken, Message: testdata.InvalidAuthorization},
	}
}
