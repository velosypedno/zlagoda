package models

import "time"

type CheckSaleItem struct {
	UPC           string  `json:"upc" binding:"required,len=12"`
	ProductNumber int     `json:"product_number" binding:"required,gte=1"`
	SellingPrice  float64 `json:"selling_price" binding:"required,gte=0"`
}

type CheckCreate struct {
	EmployeeId  string         `json:"employee_id" binding:"required,len=10"`
	CardNumber  *string        `json:"card_number" binding:"omitempty,len=13"`
	PrintDate   string         `json:"print_date" binding:"required"`
	Items       []CheckSaleItem `json:"items" binding:"required,dive,required"`
}

type CheckCreateResponse struct {
	ReceiptNumber string    `json:"receipt_number"`
	PrintDate     time.Time `json:"print_date"`
	TotalSum      float64   `json:"total_sum"`
	VAT           float64   `json:"vat"`
} 