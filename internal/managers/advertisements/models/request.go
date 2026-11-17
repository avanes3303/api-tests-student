package models

// CreateRequest хранит multipart-поля и файлы, включая невалидные значения для негативных тестов.
type CreateRequest struct {
	Fields map[string]string
	Photos []string
}
