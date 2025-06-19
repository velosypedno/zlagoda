package repos

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/velosypedno/zlagoda/internal/models"
)

type SaleRepo struct {
	db *sql.DB
}

func NewSaleRepo(db *sql.DB) *SaleRepo {
	return &SaleRepo{
		db: db,
	}
}

func (r *SaleRepo) CreateSale(s models.SaleCreate) error {
	query := `
		INSERT INTO sale (
			upc,
			receipt_number,
			product_number,
			selling_price
		) VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		query,
		s.UPC,
		s.ReceiptNumber,
		s.ProductNumber,
		s.SellingPrice,
	)

	return err
}

func (r *SaleRepo) RetrieveSaleByKey(upc, receiptNumber string) (models.SaleRetrieve, error) {
	query := `
		SELECT
			upc,
			receipt_number,
			product_number,
			selling_price
		FROM sale
		WHERE upc = $1 AND receipt_number = $2
	`

	var sale models.SaleRetrieve
	err := r.db.QueryRow(query, upc, receiptNumber).Scan(
		&sale.UPC,
		&sale.ReceiptNumber,
		&sale.ProductNumber,
		&sale.SellingPrice,
	)

	if err != nil {
		return models.SaleRetrieve{}, err
	}

	return sale, nil
}

func (r *SaleRepo) RetrieveSalesByReceipt(receiptNumber string) ([]models.SaleRetrieve, error) {
	query := `
		SELECT
			upc,
			receipt_number,
			product_number,
			selling_price
		FROM sale
		WHERE receipt_number = $1
		ORDER BY upc
	`

	rows, err := r.db.Query(query, receiptNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []models.SaleRetrieve
	for rows.Next() {
		var sale models.SaleRetrieve
		err := rows.Scan(
			&sale.UPC,
			&sale.ReceiptNumber,
			&sale.ProductNumber,
			&sale.SellingPrice,
		)
		if err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

func (r *SaleRepo) RetrieveSalesByUPC(upc string) ([]models.SaleRetrieve, error) {
	query := `
		SELECT
			upc,
			receipt_number,
			product_number,
			selling_price
		FROM sale
		WHERE upc = $1
		ORDER BY receipt_number
	`

	rows, err := r.db.Query(query, upc)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []models.SaleRetrieve
	for rows.Next() {
		var sale models.SaleRetrieve
		err := rows.Scan(
			&sale.UPC,
			&sale.ReceiptNumber,
			&sale.ProductNumber,
			&sale.SellingPrice,
		)
		if err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

func (r *SaleRepo) RetrieveAllSales() ([]models.SaleRetrieve, error) {
	query := `
		SELECT
			upc,
			receipt_number,
			product_number,
			selling_price
		FROM sale
		ORDER BY receipt_number, upc
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []models.SaleRetrieve
	for rows.Next() {
		var sale models.SaleRetrieve
		err := rows.Scan(
			&sale.UPC,
			&sale.ReceiptNumber,
			&sale.ProductNumber,
			&sale.SellingPrice,
		)
		if err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

func (r *SaleRepo) RetrieveSalesWithDetails() ([]models.SaleWithDetails, error) {
	query := `
		SELECT
			s.upc,
			s.receipt_number,
			s.product_number,
			s.selling_price,
			p.product_name,
			c.category_name,
			p.characteristics,
			(s.product_number * s.selling_price) as total_price
		FROM sale s
		JOIN store_product sp ON s.upc = sp.upc
		JOIN product p ON sp.product_id = p.product_id
		JOIN category c ON p.category_id = c.category_id
		ORDER BY s.receipt_number, s.upc
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []models.SaleWithDetails
	for rows.Next() {
		var sale models.SaleWithDetails
		err := rows.Scan(
			&sale.UPC,
			&sale.ReceiptNumber,
			&sale.ProductNumber,
			&sale.SellingPrice,
			&sale.ProductName,
			&sale.CategoryName,
			&sale.Characteristics,
			&sale.TotalPrice,
		)
		if err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

func (r *SaleRepo) RetrieveSalesWithDetailsByReceipt(receiptNumber string) ([]models.SaleWithDetails, error) {
	query := `
		SELECT
			s.upc,
			s.receipt_number,
			s.product_number,
			s.selling_price,
			p.product_name,
			c.category_name,
			p.characteristics,
			(s.product_number * s.selling_price) as total_price
		FROM sale s
		JOIN store_product sp ON s.upc = sp.upc
		JOIN product p ON sp.product_id = p.product_id
		JOIN category c ON p.category_id = c.category_id
		WHERE s.receipt_number = $1
		ORDER BY s.upc
	`

	rows, err := r.db.Query(query, receiptNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []models.SaleWithDetails
	for rows.Next() {
		var sale models.SaleWithDetails
		err := rows.Scan(
			&sale.UPC,
			&sale.ReceiptNumber,
			&sale.ProductNumber,
			&sale.SellingPrice,
			&sale.ProductName,
			&sale.CategoryName,
			&sale.Characteristics,
			&sale.TotalPrice,
		)
		if err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}

	return sales, nil
}

func (r *SaleRepo) UpdateSale(upc, receiptNumber string, s models.SaleUpdate) error {
	setParts := []string{}
	args := []interface{}{}
	argIndex := 1

	if s.ProductNumber != nil {
		setParts = append(setParts, fmt.Sprintf("product_number = $%d", argIndex))
		args = append(args, *s.ProductNumber)
		argIndex++
	}

	if s.SellingPrice != nil {
		setParts = append(setParts, fmt.Sprintf("selling_price = $%d", argIndex))
		args = append(args, *s.SellingPrice)
		argIndex++
	}

	if len(setParts) == 0 {
		return nil // Nothing to update
	}

	query := fmt.Sprintf(`
		UPDATE sale
		SET %s
		WHERE upc = $%d AND receipt_number = $%d
	`, strings.Join(setParts, ", "), argIndex, argIndex+1)

	args = append(args, upc, receiptNumber)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *SaleRepo) DeleteSale(upc, receiptNumber string) error {
	query := `DELETE FROM sale WHERE upc = $1 AND receipt_number = $2`
	_, err := r.db.Exec(query, upc, receiptNumber)
	return err
}

func (r *SaleRepo) DeleteSalesByReceipt(receiptNumber string) error {
	query := `DELETE FROM sale WHERE receipt_number = $1`
	_, err := r.db.Exec(query, receiptNumber)
	return err
}

func (r *SaleRepo) GetReceiptTotal(receiptNumber string) (float64, error) {
	query := `
		SELECT COALESCE(SUM(product_number * selling_price), 0)
		FROM sale
		WHERE receipt_number = $1
	`

	var total float64
	err := r.db.QueryRow(query, receiptNumber).Scan(&total)
	return total, err
}

func (r *SaleRepo) GetSalesStatsByProduct(productID int, startDate, endDate string) (int, float64, error) {
	query := `
		SELECT
			COALESCE(SUM(s.product_number), 0) as total_quantity,
			COALESCE(SUM(s.product_number * s.selling_price), 0) as total_revenue
		FROM sale s
		JOIN store_product sp ON s.upc = sp.upc
		JOIN receipt r ON s.receipt_number = r.receipt_number
		WHERE sp.product_id = $1
		AND r.print_date >= $2::date
		AND r.print_date <= $3::date
	`

	var totalQuantity int
	var totalRevenue float64
	err := r.db.QueryRow(query, productID, startDate, endDate).Scan(&totalQuantity, &totalRevenue)
	return totalQuantity, totalRevenue, err
}

func (r *SaleRepo) GetTopSellingProducts(limit int) ([]struct {
	ProductID    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	TotalSold    int     `json:"total_sold"`
	TotalRevenue float64 `json:"total_revenue"`
}, error) {
	query := `
		SELECT
			p.product_id,
			p.product_name,
			SUM(s.product_number) as total_sold,
			SUM(s.product_number * s.selling_price) as total_revenue
		FROM sale s
		JOIN store_product sp ON s.upc = sp.upc
		JOIN product p ON sp.product_id = p.product_id
		GROUP BY p.product_id, p.product_name
		ORDER BY total_sold DESC
		LIMIT $1
	`

	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []struct {
		ProductID    int     `json:"product_id"`
		ProductName  string  `json:"product_name"`
		TotalSold    int     `json:"total_sold"`
		TotalRevenue float64 `json:"total_revenue"`
	}

	for rows.Next() {
		var item struct {
			ProductID    int     `json:"product_id"`
			ProductName  string  `json:"product_name"`
			TotalSold    int     `json:"total_sold"`
			TotalRevenue float64 `json:"total_revenue"`
		}
		err := rows.Scan(
			&item.ProductID,
			&item.ProductName,
			&item.TotalSold,
			&item.TotalRevenue,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	return results, nil
}
