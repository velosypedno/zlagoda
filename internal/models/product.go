// internal/models/product.go
package models

type ProductCreate struct {
	Name            string `json:"name" binding:"required"` // -> product_name
	Characteristics string `json:"characteristics" binding:"required"`
	CategoryID      int    `json:"category_id" binding:"required"`
}

type ProductUpdate struct {
	Name            *string `json:"name"`
	Characteristics *string `json:"characteristics"`
	CategoryID      *int    `json:"category_id"`
}

type ProductRetrieve struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Characteristics string `json:"characteristics"`
	CategoryID      int    `json:"category_id"`
}
