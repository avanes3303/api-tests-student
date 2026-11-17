package advertisementCreate

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/testdata"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
)

// TestCreateUnauthorized проверяет отсутствие и невалидность токена при создании.
func TestCreateUnauthorized(t *testing.T) {
	fixtures.Init(t)
	for _, tc := range fixtures.UnauthorizedCases() {
		t.Run(tc.Name, func(t *testing.T) {
			base.Step("Попытаться создать валидное объявление без авторизации")
			err := advertisements.CreateAdvertisementError(t, tc.Token, fixtures.Advertisement(t), http.StatusUnauthorized)
			assertions.ErrorResponse(t, err, testdata.Unauthorized, tc.Message)
		})
	}
}
