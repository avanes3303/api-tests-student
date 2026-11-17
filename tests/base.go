package tests

import (
	"log"

	"api-tests-template/internal/utils"
)

// SetupSuite действия перед запуском сьюта с тестами
func SetupSuite() {
	log.Println("Init environment variables")
	utils.LoadEnv()
}

// TearDownSuite действия после запуска сьюта с тестами
func TearDownSuite() {
	log.Println("Tear down suite")
}

// Precondition выводит текст в консоль с меткой [Precondition]
func Precondition(text string) {
	utils.LogWithLabelAndTimeStamp("Precondition", text)
}

// Step записывает действие сценария в журнал теста.
func Step(text string) {
	utils.LogWithLabelAndTimeStamp("Шаг", text)
}

// Check записывает успешный результат проверок сценария.
func Check(text string) {
	utils.LogWithLabelAndTimeStamp("Проверка", text)
}
