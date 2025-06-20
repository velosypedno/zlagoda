package models

type StoreProductCreate struct {
	UPCProm            *string `json:"upc_prom" binding:"omitempty,len=12"`
	ProductID          int     `json:"product_id" binding:"required"`
	SellingPrice       float64 `json:"selling_price" binding:"required,gte=0"`
	ProductsNumber     int     `json:"products_number" binding:"required,gte=0"`
	PromotionalProduct bool    `json:"promotional_product" binding:"required"`
}

type StoreProductRetrieve struct {
	UPC                string  `json:"upc"`
	UPCProm            *string `json:"upc_prom"`
	ProductID          int     `json:"product_id"`
	SellingPrice       float64 `json:"selling_price"`
	ProductsNumber     int     `json:"products_number"`
	PromotionalProduct bool    `json:"promotional_product"`
}

type StoreProductUpdate struct {
	UPCProm            *string  `json:"upc_prom" binding:"omitempty,len=12"`
	ProductID          *int     `json:"product_id"`
	SellingPrice       *float64 `json:"selling_price" binding:"omitempty,gte=0"`
	ProductsNumber     *int     `json:"products_number" binding:"omitempty,gte=0"`
	PromotionalProduct *bool    `json:"promotional_product"`
}

type StoreProductWithDetails struct {
	UPC                string  `json:"upc"`
	UPCProm            *string `json:"upc_prom"`
	ProductID          int     `json:"product_id"`
	ProductName        string  `json:"product_name"`
	CategoryName       string  `json:"category_name"`
	Characteristics    string  `json:"characteristics"`
	SellingPrice       float64 `json:"selling_price"`
	ProductsNumber     int     `json:"products_number"`
	PromotionalProduct bool    `json:"promotional_product"`
}
