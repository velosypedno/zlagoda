package repos

import (
	"database/sql"

	"github.com/velosypedno/zlagoda/internal/models"
)

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) CreateCategory(c models.CategoryCreate) error {
	query := `INSERT INTO category (category_name) VALUES($1)`
	_, err := r.db.Exec(query, c.Name)
	return err
}

func (r *CategoryRepo) RetrieveCategoryByID(id int) (models.CategoryRetrieve, error) {
	query := `SELECT category_id, category_name FROM category WHERE category_id = $1`
	row := r.db.QueryRow(query, id)

	var category models.CategoryRetrieve
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		return models.CategoryRetrieve{}, err
	}

	return category, nil
}
