package repos

import (
	"database/sql"
	"fmt"

	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/utils"
)

func getNewCardNumber(r *CustomerCardRepo) (string, error) {
	const maxRetries = 10

	for i := 0; i < maxRetries; i++ {
		cardNumber, err := utils.GenerateID(13)
		if err != nil {
			return "", fmt.Errorf("failed to generate card number: %w", err)
		}

		var exists bool
		err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM customer_card WHERE card_number = $1)", cardNumber).Scan(&exists)
		if err != nil {
			return "", fmt.Errorf("failed to check card number uniqueness: %w", err)
		}

		if !exists {
			return cardNumber, nil
		}
	}

	return "", fmt.Errorf("failed to generate unique card number after %d attempts", maxRetries)
}

type CustomerCardRepo struct {
	db *sql.DB
}

func NewCustomerCardRepo(db *sql.DB) *CustomerCardRepo {
	return &CustomerCardRepo{
		db: db,
	}
}

func (r *CustomerCardRepo) CreateCustomerCard(c models.CustomerCardCreate) (string, error) {
	query := `
		INSERT INTO customer_card (
			card_number,
			cust_surname,
			cust_name,
			cust_patronymic,
			phone_number,
			city,
			street,
			zip_code,
			percent
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING card_number
	`

	cardNumber, err := getNewCardNumber(r)
	if err != nil {
		return "", err
	}
	err = r.db.QueryRow(
		query,
		cardNumber,
		c.Surname,
		c.Name,
		c.Patronymic,
		c.PhoneNumber,
		c.City,
		c.Street,
		c.ZipCode,
		c.Percent,
	).Scan(&cardNumber)

	return cardNumber, err
}

func (r *CustomerCardRepo) RetrieveCustomerCardByCardNumber(cardNumber string) (models.CustomerCardRetrieve, error) {
	query := `
		SELECT
			card_number,
			cust_surname,
			cust_name,
			cust_patronymic,
			phone_number,
			city,
			street,
			zip_code,
			percent
		FROM customer_card
		WHERE card_number = $1
	`
	var customerCard models.CustomerCardRetrieve
	err := r.db.QueryRow(query, cardNumber).Scan(
		&customerCard.CardNumber,
		&customerCard.Surname,
		&customerCard.Name,
		&customerCard.Patronymic,
		&customerCard.PhoneNumber,
		&customerCard.City,
		&customerCard.Street,
		&customerCard.ZipCode,
		&customerCard.Percent,
	)
	if err != nil {
		return models.CustomerCardRetrieve{}, err
	}

	return customerCard, nil
}

func (r *CustomerCardRepo) RetrieveCustomerCards() ([]models.CustomerCardRetrieve, error) {
	query := `
		SELECT
			card_number,
			cust_surname,
			cust_name,
			cust_patronymic,
			phone_number,
			city,
			street,
			zip_code,
			percent
		FROM customer_card
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customerCards []models.CustomerCardRetrieve
	for rows.Next() {
		var customerCard models.CustomerCardRetrieve
		err := rows.Scan(
			&customerCard.CardNumber,
			&customerCard.Surname,
			&customerCard.Name,
			&customerCard.Patronymic,
			&customerCard.PhoneNumber,
			&customerCard.City,
			&customerCard.Street,
			&customerCard.ZipCode,
			&customerCard.Percent,
		)
		if err != nil {
			return nil, err
		}
		customerCards = append(customerCards, customerCard)
	}

	return customerCards, nil
}

func (r *CustomerCardRepo) DeleteCustomerCard(cardNumber string) error {
	query := `DELETE FROM customer_card WHERE card_number = $1`
	_, err := r.db.Exec(query, cardNumber)
	return err
}

func (r *CustomerCardRepo) UpdateCustomerCard(cardNumber string, c models.CustomerCardUpdate) error {
	query := `
		UPDATE customer_card
		SET
			cust_surname = $2,
			cust_name = $3,
			cust_patronymic = $4,
			phone_number = $5,
			city = $6,
			street = $7,
			zip_code = $8,
			percent = $9
		WHERE card_number = $1
	`
	_, err := r.db.Exec(
		query,
		cardNumber,
		c.Surname,
		c.Name,
		c.Patronymic,
		c.PhoneNumber,
		c.City,
		c.Street,
		c.ZipCode,
		c.Percent,
	)
	return err
}
