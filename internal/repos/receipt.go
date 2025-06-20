package repos

import (
	"database/sql"
	"fmt"

	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/utils"
)

func getNewReceiptNumber(r *ReceiptRepo) (string, error) {
	const maxRetries = 10

	for i := 0; i < maxRetries; i++ {
		receiptNumber, err := utils.GenerateID(10)
		if err != nil {
			return "", fmt.Errorf("failed to generate receipt number: %w", err)
		}

		var exists bool
		err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM receipt WHERE receipt_number = $1)", receiptNumber).Scan(&exists)
		if err != nil {
			return "", fmt.Errorf("failed to check receipt number uniqueness: %w", err)
		}

		if !exists {
			return receiptNumber, nil
		}
	}

	return "", fmt.Errorf("failed to generate unique receipt number after %d attempts", maxRetries)
}

type ReceiptRepo struct {
	db *sql.DB
}

func NewReceiptRepo(db *sql.DB) *ReceiptRepo {
	return &ReceiptRepo{
		db: db,
	}
}

func (r *ReceiptRepo) CreateReceipt(c models.ReceiptCreate) (string, error) {
	query := `
		INSERT INTO receipt (
			receipt_number,
			employee_id,
			card_number,
			print_date,
			sum_total,
			vat
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING receipt_number
	`

	receiptNumber, err := getNewReceiptNumber(r)
	if err != nil {
		return "", err
	}
	err = r.db.QueryRow(
		query,
		receiptNumber,
		c.EmployeeId,
		c.CardNumber,
		c.PrintDate,
		c.TotalSum,
		c.VAT,
	).Scan(&receiptNumber)

	return receiptNumber, err
}

func (r *ReceiptRepo) RetrieveReceiptByReceiptNumber(receiptNumber string) (models.ReceiptRetrieve, error) {
	query := `
		SELECT
			receipt_number,
			employee_id,
			card_number,
			print_date,
			sum_total,
			vat
		FROM receipt
		WHERE receipt_number = $1
	`
	var receipt models.ReceiptRetrieve
	err := r.db.QueryRow(query, receiptNumber).Scan(
		&receipt.ReceiptNumber,
		&receipt.EmployeeId,
		&receipt.CardNumber,
		&receipt.PrintDate,
		&receipt.TotalSum,
		&receipt.VAT,
	)
	if err != nil {
		return models.ReceiptRetrieve{}, err
	}

	return receipt, nil
}

func (r *ReceiptRepo) RetrieveReceipts() ([]models.ReceiptRetrieve, error) {
	query := `
		SELECT
			receipt_number,
			employee_id,
			card_number,
			print_date,
			sum_total,
			vat
		FROM receipt
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var receipts []models.ReceiptRetrieve
	for rows.Next() {
		var receipt models.ReceiptRetrieve
		err := rows.Scan(
			&receipt.ReceiptNumber,
			&receipt.EmployeeId,
			&receipt.CardNumber,
			&receipt.PrintDate,
			&receipt.TotalSum,
			&receipt.VAT,
		)
		if err != nil {
			return nil, err
		}
		receipts = append(receipts, receipt)
	}

	return receipts, nil
}

func (r *ReceiptRepo) DeleteReceipt(receiptNumber string) error {
	query := `DELETE FROM receipt WHERE receipt_number = $1`
	_, err := r.db.Exec(query, receiptNumber)
	return err
}

func (r *ReceiptRepo) UpdateReceipt(receiptNumber string, c models.ReceiptUpdate) error {
	query := `
		UPDATE receipt
		SET
			employee_id = $2,
			card_number = $3,
			print_date = $4,
			sum_total = $5,
			vat = $6
		WHERE receipt_number = $1
	`
	_, err := r.db.Exec(
		query,
		receiptNumber,
		c.EmployeeId,
		c.CardNumber,
		c.PrintDate,
		c.TotalSum,
		c.VAT,
	)
	return err
}
