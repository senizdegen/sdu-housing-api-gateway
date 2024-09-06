package property_service

import "time"

type Property struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Price       float64   `json:"price"`
	Bedrooms    int       `json:"bedrooms"`
	Bathrooms   int       `json:"bathrooms"`
	Images      []string  `json:"images"` // Массив ссылок на изображения
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreatePropertyDTO struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Price       float64  `json:"price"`
	Bedrooms    int      `json:"bedrooms"`
	Bathrooms   int      `json:"bathrooms"`
	Images      []string `json:"images"` // Массив ссылок на изображения
}

type UpdatePropertyDTO struct {
}

type CreatePropertyResponse struct {
	UUID string `json:"uuid"`
}
