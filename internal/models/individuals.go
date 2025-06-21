package models

// Vlad1 - Most sold product in a category within a time period
type Vlad1Response struct {
	CategoryID     int     `json:"category_id"`
	CategoryName   string  `json:"category_name"`
	ProductID      int     `json:"product_id"`
	ProductName    string  `json:"product_name"`
	TotalSales     int     `json:"total_sales"`
	TotalUnitsSold int     `json:"total_units_sold"`
	TotalRevenue   float64 `json:"total_revenue"`
}

// Vlad2 - Employees who never sold promotional products
type Vlad2Response struct {
	EmployeeID   string `json:"employee_id"`
	EmployeeName string `json:"employee_name"`
	Surname      string `json:"surname"`
}

// Arthur1 - Category sales statistics within date range
type Arthur1Response struct {
	CategoryName string  `json:"category_name"`
	UnitsSold    int     `json:"units_sold"`
	Revenue      float64 `json:"revenue"`
}

// Arthur2 - Products in store that have never been sold and are not promotional
type Arthur2Response struct {
	UPC            string `json:"upc"`
	ProductName    string `json:"product_name"`
	ProductsNumber int    `json:"products_number"`
	CategoryName   string `json:"category_name"`
}

// Oleksii1 - Cashiers who served customers with high discount
type Oleksii1Response struct {
	EmployeeID            string  `json:"employee_id"`
	EmployeeSurname       string  `json:"employee_surname"`
	EmployeeName          string  `json:"employee_name"`
	HighDiscountCustomers int     `json:"high_discount_customers"`
	TotalReceiptsHighDisc int     `json:"total_receipts_high_discount"`
	TotalRevenueHighDisc  float64 `json:"total_revenue_high_discount"`
	AvgReceiptAmount      float64 `json:"avg_receipt_amount"`
	AvgCustomerDiscount   float64 `json:"avg_customer_discount"`
}

// Oleksii2 - Customers who bought from all categories in the last month
type Oleksii2Response struct {
	CardNumber  string `json:"card_number"`
	Surname     string `json:"cust_surname"`
	Name        string `json:"cust_name"`
	PhoneNumber string `json:"phone_number"`
}
