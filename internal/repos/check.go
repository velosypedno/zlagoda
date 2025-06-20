package repos

import (
	"database/sql"
	"fmt"
	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/utils"
)

type CheckRepoImpl struct {
	db *sql.DB
}

func NewCheckRepo(db *sql.DB) *CheckRepoImpl {
	return &CheckRepoImpl{db: db}
}

func (r *CheckRepoImpl) BeginTx() (*sql.Tx, error) {
	return r.db.Begin()
}

// Generate a unique receipt_number within the current transaction
func getNewCheckReceiptNumber(tx *sql.Tx) (string, error) {
	const maxRetries = 10
	for i := 0; i < maxRetries; i++ {
		receiptNumber, err := utils.GenerateSecureID(10)
		if err != nil {
			return "", err
		}
		var exists bool
		err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM receipt WHERE receipt_number = $1)", receiptNumber).Scan(&exists)
		if err != nil {
			return "", err
		}
		if !exists {
			return receiptNumber, nil
		}
	}
	return "", fmt.Errorf("failed to generate unique receipt number")
}

func (r *CheckRepoImpl) CreateReceiptTx(tx *sql.Tx, c models.ReceiptCreate) (string, error) {
	receiptNumber, err := getNewCheckReceiptNumber(tx)
	if err != nil {
		return "", err
	}
	query := `INSERT INTO receipt (receipt_number, employee_id, card_number, print_date, sum_total, vat) VALUES ($1, $2, $3, $4, $5, $6) RETURNING receipt_number`
	err = tx.QueryRow(query, receiptNumber, c.EmployeeId, c.CardNumber, c.PrintDate, c.TotalSum, c.VAT).Scan(&receiptNumber)
	return receiptNumber, err
}

func (r *CheckRepoImpl) CreateSaleTx(tx *sql.Tx, s models.SaleCreate) error {
	query := `INSERT INTO sale (upc, receipt_number, product_number, selling_price) VALUES ($1, $2, $3, $4)`
	_, err := tx.Exec(query, s.UPC, s.ReceiptNumber, s.ProductNumber, s.SellingPrice)
	return err
}

func (r *CheckRepoImpl) UpdateProductQuantityTx(tx *sql.Tx, upc string, quantityChange int) error {
	query := `UPDATE store_product SET products_number = products_number + $2 WHERE upc = $1 AND products_number + $2 >= 0`
	result, err := tx.Exec(query, upc, quantityChange)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *CheckRepoImpl) GetStoreProductStockTx(tx *sql.Tx, upc string) (int, error) {
	query := `SELECT products_number FROM store_product WHERE upc = $1`
	var stock int
	err := tx.QueryRow(query, upc).Scan(&stock)
	return stock, err
}

func (r *CheckRepoImpl) RetrieveChecks() ([]models.ReceiptRetrieve, error) {
	query := `
		SELECT
			receipt_number,
			employee_id,
			card_number,
			print_date,
			sum_total,
			vat
		FROM receipt
		ORDER BY print_date DESC, receipt_number DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var checks []models.ReceiptRetrieve
	for rows.Next() {
		var check models.ReceiptRetrieve
		err := rows.Scan(
			&check.ReceiptNumber,
			&check.EmployeeId,
			&check.CardNumber,
			&check.PrintDate,
			&check.TotalSum,
			&check.VAT,
		)
		if err != nil {
			return nil, err
		}
		checks = append(checks, check)
	}

	return checks, nil
} 