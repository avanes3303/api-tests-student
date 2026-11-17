package advertisementSearch

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/fixtures"
	advList "api-tests-template/internal/managers/advertisements"
	authModels "api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

type TestSuite struct {
	suite.Suite
	loginData       authModels.LoginOkResponse
	existingAdTitle string
	searchBeginning string
	searchMiddle    string
	searchEnd       string
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	s.loginData = fixtures.New(s.T()).Owner
}

// SetupTest создаёт собственное объявление с уникальными подстроками для каждого теста.
func (s *TestSuite) SetupTest() {
	s.searchBeginning = "Учебное" + utils.RandomString(6)
	s.searchMiddle = "объявление" + utils.RandomString(6)
	s.searchEnd = "API" + utils.RandomString(6)
	s.existingAdTitle = s.searchBeginning + " " + s.searchMiddle + " " + s.searchEnd
	request := fixtures.Advertisement(s.T())
	request.Fields["title"] = s.existingAdTitle
	fixtures.WithOwner(s.T(), s.loginData).Create(request)
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

func (s *TestSuite) TestSearchByTitle() {
	base.Step("Поиск по уникальной подстроке названия")
	list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, s.searchMiddle)

	s.Run("В ответе есть хотя бы одно объявление", func() {
		require.NotEmpty(s.T(), list.Items, "Поиск вернул пустой список объявлений")
	})
	s.Run("Среди найденных есть наше объявление", func() {
		assertions.ContainsAdvertisementByTitle(s.T(), list.Items, s.existingAdTitle)
	})
}

func (s *TestSuite) TestSearchByTitle_Substring() {
	cases := []struct {
		name   string
		search string
	}{
		{"начало", s.searchBeginning},
		{"середина", s.searchMiddle},
		{"конец", s.searchEnd},
		{"нижний регистр", strings.ToLower(s.searchMiddle)},
		{"верхний регистр", strings.ToUpper(s.searchMiddle)},
		{"смешанный регистр", utils.MixCase(s.searchMiddle)},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			base.Step("Поиск названия: " + tc.name)
			list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, tc.search)
			require.NotEmpty(s.T(), list.Items, "Поиск по подстроке %q вернул пустой список", tc.search)
			assertions.ContainsAdvertisementByTitle(s.T(), list.Items, s.existingAdTitle)
		})
	}
}
