package models

import "time"

type ReceiptCreate struct {
	EmployeeId *string
	CardNumber *string
	PrintDate  *time.Time
	TotalSum   *float64
	VAT        *float64
}

type ReceiptRetrieve struct {
	ReceiptNumber *string
	EmployeeId    *string
	CardNumber    *string
	PrintDate     *time.Time
	TotalSum      *float64
	VAT           *float64
}

type ReceiptUpdate struct {
	EmployeeId *string
	CardNumber *string
	PrintDate  *time.Time
	TotalSum   *float64
	VAT        *float64
}

type ReceiptCreateComplete struct {
	EmployeeId *string
	CardNumber *string
	PrintDate  *time.Time
	Items      []ReceiptItem
}

type ReceiptItem struct {
	UPC           *string
	ProductNumber *int
	SellingPrice  *float64
}
