package models

type Product struct {
	ID          string `json:"id,omitempty"`
	NAME        string `json:"name,omitempty" validate:"required"`
	DESCRIPTION string `json:"description,omitempty" validate:"required"`
	IMAGEURL    string `json:"imageUrl,omitempty" validate:"required"`
}
