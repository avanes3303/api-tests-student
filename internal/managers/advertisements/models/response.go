package models

// AdvertisementsList описывает ответ API со списком объявлений.
type AdvertisementsList struct {
	Items      []Advertisement `json:"items"`
	Total      int             `json:"total"`
	HasNext    bool            `json:"has_next"`
	NextCursor string          `json:"next_cursor"`
}

// Advertisement содержит данные объявления, используемые в проверках поиска.
type Advertisement struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
