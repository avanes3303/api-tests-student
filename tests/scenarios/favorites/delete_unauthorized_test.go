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

// TestDeleteUnauthorized проверяет обязательную авторизацию ручки избранного.
func TestDeleteUnauthorized(t *testing.T) {
	f := fixtures.New(t)
	ad := f.Create(fixtures.Advertisement(t))
	for _, tc := range fixtures.UnauthorizedCases() {
		t.Run(tc.Name, func(t *testing.T) {
			base.Step("Delete избранного без авторизации")
			actual := favorites.DeleteFavoriteError(t, tc.Token, ad.ID, http.StatusUnauthorized)
			assertions.ErrorResponse(t, actual, testdata.Unauthorized, tc.Message)
		})
	}
}
