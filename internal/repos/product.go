package repos

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/velosypedno/zlagoda/internal/models"
)

type ProductRepo interface {
	CreateProduct(p models.ProductCreate) (int, error)
	RetrieveProductByID(id int) (models.ProductRetrieve, error)
	RetrieveProducts() ([]models.ProductRetrieve, error)
	RetrieveProductsByCategory(categoryID int) ([]models.ProductRetrieve, error)
	RetrieveProductsByName(name string) ([]models.ProductRetrieve, error)
	UpdateProduct(id int, p models.ProductUpdate) error
	DeleteProduct(id int) error
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) CreateProduct(p models.ProductCreate) (int, error) {
	var id int
	err := r.db.QueryRow(
		`INSERT INTO product (product_name, characteristics, category_id)
		 VALUES ($1, $2, $3)
		 RETURNING product_id`,
		p.Name, p.Characteristics, p.CategoryID,
	).Scan(&id)
	return id, err
}

func (r *productRepo) RetrieveProductByID(id int) (models.ProductRetrieve, error) {
	var pr models.ProductRetrieve
	err := r.db.QueryRow(
		`SELECT product_id, product_name, characteristics, category_id
		 FROM product
		 WHERE product_id = $1`,
		id,
	).Scan(&pr.ID, &pr.Name, &pr.Characteristics, &pr.CategoryID)
	return pr, err
}

func (r *productRepo) RetrieveProducts() ([]models.ProductRetrieve, error) {
	rows, err := r.db.Query(
		`SELECT product_id, product_name, characteristics, category_id
		 FROM product`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.ProductRetrieve
	for rows.Next() {
		var pr models.ProductRetrieve
		if err := rows.Scan(&pr.ID, &pr.Name, &pr.Characteristics, &pr.CategoryID); err != nil {
			return nil, err
		}
		list = append(list, pr)
	}
	return list, rows.Err()
}

func (r *productRepo) RetrieveProductsByCategory(categoryID int) ([]models.ProductRetrieve, error) {
	rows, err := r.db.Query(
		`SELECT product_id, product_name, characteristics, category_id
		 FROM product
		 WHERE category_id = $1`,
		categoryID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.ProductRetrieve
	for rows.Next() {
		var pr models.ProductRetrieve
		if err := rows.Scan(&pr.ID, &pr.Name, &pr.Characteristics, &pr.CategoryID); err != nil {
			return nil, err
		}
		list = append(list, pr)
	}
	return list, rows.Err()
}

func (r *productRepo) RetrieveProductsByName(name string) ([]models.ProductRetrieve, error) {
	rows, err := r.db.Query(
		`SELECT product_id, product_name, characteristics, category_id
		 FROM product
		 WHERE product_name LIKE $1`,
		"%"+name+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.ProductRetrieve
	for rows.Next() {
		var pr models.ProductRetrieve
		if err := rows.Scan(&pr.ID, &pr.Name, &pr.Characteristics, &pr.CategoryID); err != nil {
			return nil, err
		}
		list = append(list, pr)
	}
	return list, rows.Err()
}

func (r *productRepo) UpdateProduct(id int, p models.ProductUpdate) error {
	set, args, idx := []string{}, []interface{}{}, 1

	if p.Name != nil {
		set = append(set, fmt.Sprintf("product_name = $%d", idx))
		args = append(args, *p.Name)
		idx++
	}
	if p.Characteristics != nil {
		set = append(set, fmt.Sprintf("characteristics = $%d", idx))
		args = append(args, *p.Characteristics)
		idx++
	}
	if p.CategoryID != nil {
		set = append(set, fmt.Sprintf("category_id = $%d", idx))
		args = append(args, *p.CategoryID)
		idx++
	}
	if len(set) == 0 {
		return nil // нічого оновлювати
	}

	query := `UPDATE product SET ` + strings.Join(set, ", ") +
		fmt.Sprintf(" WHERE product_id = $%d", idx)
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *productRepo) DeleteProduct(id int) error {
	_, err := r.db.Exec(`DELETE FROM product WHERE product_id = $1`, id)
	return err
}
