package assertions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/models"
)

// ErrorResponse проверяет точный код и сообщение ожидаемой ошибки API.
func ErrorResponse(t *testing.T, actual models.ErrorResponse, code, message string) {
	t.Helper()
	require.Equal(t, code, actual.Error, "Некорректный код ошибки API")
	require.Equal(t, message, actual.Message, "Некорректное сообщение ошибки API")
	t.Logf("[Проверка] Получены ожидаемые ошибка %q и сообщение %q", code, message)
}
