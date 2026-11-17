package advertisementCreate

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/constants/testdata"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
)

// TestCreateLongTitle проверяет отклонение объявления с названием длиннее 50 символов.
func TestCreateLongTitle(t *testing.T) {
	f := fixtures.New(t)
	request := fixtures.Advertisement(t)
	request.Fields["title"] = strings.Repeat("Я", testdata.TitleLimit+1)
	base.Step("Попытаться создать объявление с названием длиннее 50 символов")
	actual := advertisements.CreateAdvertisementError(t, f.Owner.Token, request, http.StatusBadRequest)
	assertions.ErrorResponse(t, actual, testdata.InvalidRequest, fmt.Sprintf("title must not exceed %d characters provided %d", testdata.TitleLimit, testdata.TitleLimit+1))
}
