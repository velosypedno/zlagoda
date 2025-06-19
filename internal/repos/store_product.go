package repos

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/velosypedno/zlagoda/internal/models"
)

type StoreProductRepo struct {
	db *sql.DB
}

func NewStoreProductRepo(db *sql.DB) *StoreProductRepo {
	return &StoreProductRepo{
		db: db,
	}
}

func (r *StoreProductRepo) CreateStoreProduct(sp models.StoreProductCreate) (string, error) {
	query := `
		INSERT INTO store_product (
			upc,
			upc_prom,
			product_id,
			selling_price,
			products_number,
			promotional_product
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING upc
	`

	var upc string
	err := r.db.QueryRow(
		query,
		sp.UPC,
		sp.UPCProm,
		sp.ProductID,
		sp.SellingPrice,
		sp.ProductsNumber,
		sp.PromotionalProduct,
	).Scan(&upc)

	return upc, err
}

func (r *StoreProductRepo) RetrieveStoreProductByUPC(upc string) (models.StoreProductRetrieve, error) {
	query := `
		SELECT
			upc,
			upc_prom,
			product_id,
			selling_price,
			products_number,
			promotional_product
		FROM store_product
		WHERE upc = $1
	`

	var storeProduct models.StoreProductRetrieve
	err := r.db.QueryRow(query, upc).Scan(
		&storeProduct.UPC,
		&storeProduct.UPCProm,
		&storeProduct.ProductID,
		&storeProduct.SellingPrice,
		&storeProduct.ProductsNumber,
		&storeProduct.PromotionalProduct,
	)

	if err != nil {
		return models.StoreProductRetrieve{}, err
	}

	return storeProduct, nil
}

func (r *StoreProductRepo) RetrieveStoreProducts() ([]models.StoreProductRetrieve, error) {
	query := `
		SELECT
			upc,
			upc_prom,
			product_id,
			selling_price,
			products_number,
			promotional_product
		FROM store_product
		ORDER BY upc
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storeProducts []models.StoreProductRetrieve
	for rows.Next() {
		var storeProduct models.StoreProductRetrieve
		err := rows.Scan(
			&storeProduct.UPC,
			&storeProduct.UPCProm,
			&storeProduct.ProductID,
			&storeProduct.SellingPrice,
			&storeProduct.ProductsNumber,
			&storeProduct.PromotionalProduct,
		)
		if err != nil {
			return nil, err
		}
		storeProducts = append(storeProducts, storeProduct)
	}

	return storeProducts, nil
}

func (r *StoreProductRepo) RetrieveStoreProductsWithDetails() ([]models.StoreProductWithDetails, error) {
	query := `
		SELECT
			sp.upc,
			sp.upc_prom,
			sp.product_id,
			p.product_name,
			c.category_name,
			p.characteristics,
			sp.selling_price,
			sp.products_number,
			sp.promotional_product
		FROM store_product sp
		JOIN product p ON sp.product_id = p.product_id
		JOIN category c ON p.category_id = c.category_id
		ORDER BY sp.upc
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storeProducts []models.StoreProductWithDetails
	for rows.Next() {
		var storeProduct models.StoreProductWithDetails
		err := rows.Scan(
			&storeProduct.UPC,
			&storeProduct.UPCProm,
			&storeProduct.ProductID,
			&storeProduct.ProductName,
			&storeProduct.CategoryName,
			&storeProduct.Characteristics,
			&storeProduct.SellingPrice,
			&storeProduct.ProductsNumber,
			&storeProduct.PromotionalProduct,
		)
		if err != nil {
			return nil, err
		}
		storeProducts = append(storeProducts, storeProduct)
	}

	return storeProducts, nil
}

func (r *StoreProductRepo) RetrieveStoreProductsByProductID(productID int) ([]models.StoreProductRetrieve, error) {
	query := `
		SELECT
			upc,
			upc_prom,
			product_id,
			selling_price,
			products_number,
			promotional_product
		FROM store_product
		WHERE product_id = $1
		ORDER BY upc
	`

	rows, err := r.db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storeProducts []models.StoreProductRetrieve
	for rows.Next() {
		var storeProduct models.StoreProductRetrieve
		err := rows.Scan(
			&storeProduct.UPC,
			&storeProduct.UPCProm,
			&storeProduct.ProductID,
			&storeProduct.SellingPrice,
			&storeProduct.ProductsNumber,
			&storeProduct.PromotionalProduct,
		)
		if err != nil {
			return nil, err
		}
		storeProducts = append(storeProducts, storeProduct)
	}

	return storeProducts, nil
}

func (r *StoreProductRepo) RetrievePromotionalProducts() ([]models.StoreProductRetrieve, error) {
	query := `
		SELECT
			upc,
			upc_prom,
			product_id,
			selling_price,
			products_number,
			promotional_product
		FROM store_product
		WHERE promotional_product = true
		ORDER BY upc
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storeProducts []models.StoreProductRetrieve
	for rows.Next() {
		var storeProduct models.StoreProductRetrieve
		err := rows.Scan(
			&storeProduct.UPC,
			&storeProduct.UPCProm,
			&storeProduct.ProductID,
			&storeProduct.SellingPrice,
			&storeProduct.ProductsNumber,
			&storeProduct.PromotionalProduct,
		)
		if err != nil {
			return nil, err
		}
		storeProducts = append(storeProducts, storeProduct)
	}

	return storeProducts, nil
}

func (r *StoreProductRepo) UpdateStoreProduct(upc string, sp models.StoreProductUpdate) error {
	setParts := []string{}
	args := []interface{}{}
	argIndex := 1

	if sp.UPCProm != nil {
		setParts = append(setParts, fmt.Sprintf("upc_prom = $%d", argIndex))
		args = append(args, sp.UPCProm)
		argIndex++
	}

	if sp.ProductID != nil {
		setParts = append(setParts, fmt.Sprintf("product_id = $%d", argIndex))
		args = append(args, *sp.ProductID)
		argIndex++
	}

	if sp.SellingPrice != nil {
		setParts = append(setParts, fmt.Sprintf("selling_price = $%d", argIndex))
		args = append(args, *sp.SellingPrice)
		argIndex++
	}

	if sp.ProductsNumber != nil {
		setParts = append(setParts, fmt.Sprintf("products_number = $%d", argIndex))
		args = append(args, *sp.ProductsNumber)
		argIndex++
	}

	if sp.PromotionalProduct != nil {
		setParts = append(setParts, fmt.Sprintf("promotional_product = $%d", argIndex))
		args = append(args, *sp.PromotionalProduct)
		argIndex++
	}

	if len(setParts) == 0 {
		return nil // Nothing to update
	}

	query := fmt.Sprintf(`
		UPDATE store_product
		SET %s
		WHERE upc = $%d
	`, strings.Join(setParts, ", "), argIndex)

	args = append(args, upc)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *StoreProductRepo) DeleteStoreProduct(upc string) error {
	query := `DELETE FROM store_product WHERE upc = $1`
	_, err := r.db.Exec(query, upc)
	return err
}

func (r *StoreProductRepo) UpdateProductQuantity(upc string, quantityChange int) error {
	query := `
		UPDATE store_product
		SET products_number = products_number + $2
		WHERE upc = $1 AND products_number + $2 >= 0
	`
	result, err := r.db.Exec(query, upc, quantityChange)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("insufficient stock or product not found")
	}

	return nil
}

func (r *StoreProductRepo) CheckStockAvailability(upc string, requiredQuantity int) (bool, error) {
	query := `SELECT products_number FROM store_product WHERE upc = $1`

	var currentStock int
	err := r.db.QueryRow(query, upc).Scan(&currentStock)
	if err != nil {
		return false, err
	}

	return currentStock >= requiredQuantity, nil
}
