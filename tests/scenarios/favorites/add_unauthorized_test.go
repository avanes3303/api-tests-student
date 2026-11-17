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

// TestAddUnauthorized проверяет обязательную авторизацию ручки избранного.
func TestAddUnauthorized(t *testing.T) {
	f := fixtures.New(t)
	ad := f.Create(fixtures.Advertisement(t))
	for _, tc := range fixtures.UnauthorizedCases() {
		t.Run(tc.Name, func(t *testing.T) {
			base.Step("Add избранного без авторизации")
			actual := favorites.AddFavoriteError(t, tc.Token, ad.ID, http.StatusUnauthorized)
			assertions.ErrorResponse(t, actual, testdata.Unauthorized, tc.Message)
		})
	}
}
