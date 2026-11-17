package response

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// Read проверяет HTTP-статус, читает ответ и закрывает его тело.
func Read(t *testing.T, raw *http.Response, status int) []byte {
	t.Helper()
	require.NotNil(t, raw, "API не вернул HTTP-ответ")
	defer func() {
		require.NoError(t, raw.Body.Close(), "Не удалось закрыть тело HTTP-ответа")
	}()
	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err, "Не удалось прочитать HTTP-ответ")
	t.Logf("[Ответ] HTTP %d", raw.StatusCode)
	require.Equalf(t, status, raw.StatusCode, "Неожиданный HTTP-статус; тело ответа: %s", body)
	return body
}

// JSON разбирает проверенный HTTP-ответ в типизированную модель.
func JSON[T any](t *testing.T, raw *http.Response, status int) T {
	t.Helper()
	body := Read(t, raw, status)
	var value T
	require.NoError(t, json.Unmarshal(body, &value), "Некорректный JSON в ответе API")
	return value
}
