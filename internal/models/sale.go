package models

type SaleCreate struct {
	UPC           string  `json:"upc" binding:"required,len=12"`
	ReceiptNumber string  `json:"receipt_number" binding:"required,len=10"`
	ProductNumber int     `json:"product_number" binding:"required,gte=1"`
	SellingPrice  float64 `json:"selling_price" binding:"required,gte=0"`
}

type SaleRetrieve struct {
	UPC           string  `json:"upc"`
	ReceiptNumber string  `json:"receipt_number"`
	ProductNumber int     `json:"product_number"`
	SellingPrice  float64 `json:"selling_price"`
}

type SaleUpdate struct {
	ProductNumber *int     `json:"product_number" binding:"omitempty,gte=1"`
	SellingPrice  *float64 `json:"selling_price" binding:"omitempty,gte=0"`
}

// Extended model with product details for API responses
type SaleWithDetails struct {
	UPC             string  `json:"upc"`
	ReceiptNumber   string  `json:"receipt_number"`
	ProductNumber   int     `json:"product_number"`
	SellingPrice    float64 `json:"selling_price"`
	ProductName     string  `json:"product_name"`
	CategoryName    string  `json:"category_name"`
	Characteristics string  `json:"characteristics"`
	TotalPrice      float64 `json:"total_price"` // ProductNumber * SellingPrice
}

// Composite key for sale operations
type SaleKey struct {
	UPC           string `json:"upc"`
	ReceiptNumber string `json:"receipt_number"`
}
