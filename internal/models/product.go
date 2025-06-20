package models

type ProductCreate struct {
	CategoryID      int    `json:"category_id" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Characteristics string `json:"characteristics" binding:"required"`
}

type ProductRetrieve struct {
	ID              int    `json:"id"`
	CategoryID      int    `json:"category_id"`
	Name            string `json:"name"`
	Characteristics string `json:"characteristics"`
}

type ProductUpdate struct {
	CategoryID      *int    `json:"category_id"`
	Name            *string `json:"name"`
	Characteristics *string `json:"characteristics"`
}
