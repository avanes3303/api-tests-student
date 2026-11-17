package models

import advertisements "api-tests-template/internal/managers/advertisements/models"

// AddRequest задаёт объявление, добавляемое в избранное.
type AddRequest struct {
	AdvertisementID string `json:"advertisement_id,omitempty"`
}

// Favorite содержит идентификатор записи и полные данные объявления.
type Favorite struct {
	ID            string                       `json:"favorite_id"`
	AddedAt       string                       `json:"added_at"`
	Advertisement advertisements.Advertisement `json:"advertisement"`
}

// ListResponse описывает страницу избранного текущего пользователя.
type ListResponse struct {
	Items      []Favorite `json:"items"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
}
