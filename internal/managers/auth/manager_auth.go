package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/auth"
	"api-tests-template/internal/helpers/response"
	"api-tests-template/internal/managers/auth/models"
)

// Login выполняет авторизацию пользователя и возвращает токен с данными пользователя, ожидает 200
func Login(t *testing.T, username, password string) models.LoginOkResponse {
	t.Helper()
	require.NotEmpty(t, username, "Отсутствует логин пользователя")
	require.NotEmpty(t, password, "Отсутствует пароль пользователя")

	request := models.LoginRequest{
		Email:    username,
		Password: password,
	}

	loginResponse := response.JSON[models.LoginOkResponse](t, auth.HttpPostAuthLogin(t, request), http.StatusOK)
	require.NotEmpty(t, loginResponse.Token, "Авторизация не вернула токен")
	require.NotEmpty(t, loginResponse.User.Id, "Авторизация не вернула идентификатор пользователя")

	return loginResponse
}

// EnsureTestUser создаёт изолированного пользователя или входит в существующую тестовую учётную запись.
func EnsureTestUser(t *testing.T, data models.RegisterRequest) models.LoginOkResponse {
	t.Helper()
	raw := auth.HttpPostAuthRegister(t, data)
	require.NotNil(t, raw, "API регистрации не вернул HTTP-ответ")
	if raw.StatusCode == http.StatusConflict {
		response.Read(t, raw, http.StatusConflict)
		return Login(t, data.Email, data.Password)
	}
	user := response.JSON[models.LoginOkResponse](t, raw, http.StatusCreated)
	require.NotEmpty(t, user.Token, "Регистрация не вернула токен")
	require.NotEmpty(t, user.User.Id, "Регистрация не вернула идентификатор пользователя")
	return user
}
