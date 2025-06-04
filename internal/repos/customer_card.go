package repos

import (
	"database/sql"

	"math/rand"
	"slices"

	"github.com/velosypedno/zlagoda/internal/models"
)

type CustomerCardRepo struct {
	db *sql.DB
}

func NewCustomerCardRepo(db *sql.DB) *CustomerCardRepo {
	return &CustomerCardRepo{
		db: db,
	}
}

func generateCardNumber() string {
	var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 13)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}

func getNewCardNumber(r *CustomerCardRepo) (string, error) {
	query := `SELECT card_number FROM customer_card`

	rows, err := r.db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var cardNumbers []string
	for rows.Next() {
		var currCardNumber string
		err := rows.Scan(&currCardNumber)
		if err != nil {
			return "", err
		}
		cardNumbers = append(cardNumbers, currCardNumber)
	}

	var newCardNumber string
	for {
		newCardNumber = generateCardNumber()
		if !slices.Contains(cardNumbers, newCardNumber) {
			break
		}
	}
	return newCardNumber, err
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
	row := r.db.QueryRow(query, cardNumber)

	var customerCard models.CustomerCardRetrieve
	err := row.Scan(
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
			percent = $9,
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
