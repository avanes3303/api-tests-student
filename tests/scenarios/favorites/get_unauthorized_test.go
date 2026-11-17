package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/testdata"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

// TestGetUnauthorized проверяет обязательную авторизацию ручки избранного.
func TestGetUnauthorized(t *testing.T) {
	fixtures.Init(t)
	for _, tc := range fixtures.UnauthorizedCases() {
		t.Run(tc.Name, func(t *testing.T) {
			base.Step("Get избранного без авторизации")
			actual := favorites.GetFavoritesError(t, tc.Token, http.StatusUnauthorized)
			assertions.ErrorResponse(t, actual, testdata.Unauthorized, tc.Message)
		})
	}
}
