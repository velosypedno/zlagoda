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

func (r *CategoryRepo) CreateCategory(c models.CategoryCreate) (int, error) {
	var id int
	query := `INSERT INTO category (category_name) VALUES($1) RETURNING category_id`
	err := r.db.QueryRow(query, c.Name).Scan(&id)
	return id, err
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

func (r *CategoryRepo) RetrieveCategories() ([]models.CategoryRetrieve, error) {
	query := `SELECT category_id, category_name FROM category`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.CategoryRetrieve
	for rows.Next() {
		var category models.CategoryRetrieve
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
