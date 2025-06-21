package repos

import (
	"database/sql"

	"github.com/velosypedno/zlagoda/internal/models"
)

type IndividualsRepo struct {
	db *sql.DB
}

func NewIndividualsRepo(db *sql.DB) *IndividualsRepo {
	return &IndividualsRepo{
		db: db,
	}
}

func (r *IndividualsRepo) QueryVlad1(categoryID int, months int) ([]models.Vlad1Response, error) {
	query := `
	SELECT
	    c.category_id,
	    c.category_name,
	    p.product_id,
	    p.product_name,
	    COUNT(s.upc) AS total_sales,
	    SUM(s.product_number) AS total_units_sold,
	    SUM(s.product_number * s.selling_price) AS total_revenue
	FROM
	    category c
	JOIN
	    product p ON c.category_id = p.category_id
	JOIN
	    store_product sp ON p.product_id = sp.product_id
	JOIN
	    sale s ON sp.upc = s.upc
	JOIN
	    receipt r ON s.receipt_number = r.receipt_number
	WHERE
	    r.print_date BETWEEN CURRENT_DATE - ($2 * INTERVAL '1 month') AND CURRENT_DATE
	    AND c.category_id = $1
	GROUP BY
	    c.category_id, c.category_name, p.product_id, p.product_name
	HAVING
	    SUM(s.product_number) >= ALL (
	        SELECT SUM(s2.product_number)
	        FROM
	            product p2
	        JOIN
	            store_product sp2 ON p2.product_id = sp2.product_id
	        JOIN
	            sale s2 ON sp2.upc = s2.upc
	        JOIN
	            receipt r2 ON s2.receipt_number = r2.receipt_number
	        WHERE
	            p2.category_id = c.category_id
	            AND r2.print_date BETWEEN CURRENT_DATE - ($2 * INTERVAL '1 month') AND CURRENT_DATE
	        GROUP BY
	            p2.product_id
	    )
	ORDER BY
	    total_units_sold DESC, total_revenue DESC
	LIMIT 5;
	`
	rows, err := r.db.Query(query, categoryID, months)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Vlad1Response
	for rows.Next() {
		var result models.Vlad1Response
		err := rows.Scan(
			&result.CategoryID,
			&result.CategoryName,
			&result.ProductID,
			&result.ProductName,
			&result.TotalSales,
			&result.TotalUnitsSold,
			&result.TotalRevenue,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *IndividualsRepo) QueryVlad2() ([]models.Vlad2Response, error) {
	query := `
	SELECT
	    e.employee_id,
	    e.empl_surname,
	    e.empl_name
	FROM
	    employee AS e
	WHERE
	    NOT EXISTS (
	        SELECT 1
	        FROM
	            receipt AS r
	        WHERE
	            r.employee_id = e.employee_id
	            AND EXISTS (
	                SELECT 1
	                FROM
	                    sale AS s
	                    INNER JOIN store_product AS sp
	                        ON s.upc = sp.upc
	                WHERE
	                    s.receipt_number = r.receipt_number
	                    AND sp.promotional_product = TRUE
	            )
	    )
	LIMIT 5;
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Vlad2Response
	for rows.Next() {
		var result models.Vlad2Response
		err := rows.Scan(
			&result.EmployeeID,
			&result.Surname,
			&result.EmployeeName,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *IndividualsRepo) QueryArthur1(startDate, endDate string) ([]models.Arthur1Response, error) {
	query := `
	SELECT
	    c.category_name,
	    SUM(s.product_number) as units_sold,
	    SUM(s.selling_price * s.product_number) as revenue
	FROM sale s
	    JOIN receipt r ON s.receipt_number = r.receipt_number
	    JOIN store_product sp ON s.upc = sp.upc
	    JOIN product p ON sp.product_id = p.product_id
	    JOIN category c ON p.category_id = c.category_id
	WHERE r.print_date BETWEEN $1::date AND $2::date
	GROUP BY c.category_name
	ORDER BY revenue DESC, c.category_name ASC
	LIMIT 5;
	`
	rows, err := r.db.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Arthur1Response
	for rows.Next() {
		var result models.Arthur1Response
		err := rows.Scan(
			&result.CategoryName,
			&result.UnitsSold,
			&result.Revenue,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *IndividualsRepo) QueryArthur2() ([]models.Arthur2Response, error) {
	query := `
	SELECT
	    sp.upc,
	    p.product_name,
	    sp.products_number,
	    c.category_name
	FROM              store_product sp
	INNER JOIN        product       p  ON p.product_id  = sp.product_id
	INNER JOIN        category      c  ON c.category_id = p.category_id

	WHERE NOT EXISTS (
	          SELECT 1
	          FROM   sale s
	          WHERE  s.upc = sp.upc
	      )

	  AND NOT EXISTS (
	          SELECT 1
	          FROM   store_product sp2
	          WHERE  sp2.upc = sp.upc
	            AND (sp2.upc_prom IS NOT NULL
	                 OR sp2.promotional_product = TRUE)
	      )

	  AND sp.products_number > 0
	ORDER BY
	    c.category_name,
	    p.product_name
	LIMIT 5;
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Arthur2Response
	for rows.Next() {
		var result models.Arthur2Response
		err := rows.Scan(
			&result.UPC,
			&result.ProductName,
			&result.ProductsNumber,
			&result.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *IndividualsRepo) QueryOleksii1(discountThreshold int) ([]models.Oleksii1Response, error) {
	query := `
	SELECT
		e.employee_id,
		e.empl_surname,
		e.empl_name,
		COUNT(DISTINCT cc.card_number) as high_discount_customers,
		COUNT(DISTINCT r.receipt_number) as total_receipts_high_discount,
		SUM(r.sum_total) as total_revenue_high_discount,
		AVG(r.sum_total) as avg_receipt_amount,
		AVG(cc.percent) as avg_customer_discount
	FROM employee e
	JOIN receipt r ON e.employee_id = r.employee_id
	JOIN customer_card cc ON r.card_number = cc.card_number
	WHERE cc.percent > $1
		AND e.empl_role = 'Cashier'
	GROUP BY e.employee_id, e.empl_surname, e.empl_name
	HAVING COUNT(DISTINCT cc.card_number) > 0
	ORDER BY high_discount_customers DESC,
			total_revenue_high_discount DESC,
			e.empl_surname ASC
	LIMIT 5;
	`
	rows, err := r.db.Query(query, discountThreshold)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Oleksii1Response
	for rows.Next() {
		var result models.Oleksii1Response
		err := rows.Scan(
			&result.EmployeeID,
			&result.EmployeeSurname,
			&result.EmployeeName,
			&result.HighDiscountCustomers,
			&result.TotalReceiptsHighDisc,
			&result.TotalRevenueHighDisc,
			&result.AvgReceiptAmount,
			&result.AvgCustomerDiscount,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *IndividualsRepo) QueryOleksii2() ([]models.Oleksii2Response, error) {
	query := `
	SELECT
	    cc.card_number,
	    cc.cust_surname,
	    cc.cust_name,
	    cc.phone_number
	FROM customer_card cc
	WHERE NOT EXISTS (
	    SELECT 1
	    FROM category c
	    WHERE NOT EXISTS (
	        SELECT 1
	        FROM sale s
	        JOIN receipt       r  ON r.receipt_number = s.receipt_number
	        JOIN store_product sp ON sp.upc           = s.upc
	        JOIN product       p  ON p.product_id     = sp.product_id
	        WHERE r.card_number = cc.card_number
	          AND r.print_date  >= CURRENT_DATE - INTERVAL '1 month'
	          AND p.category_id = c.category_id
	    )
	)
	LIMIT 5;
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Oleksii2Response
	for rows.Next() {
		var result models.Oleksii2Response
		err := rows.Scan(
			&result.CardNumber,
			&result.Surname,
			&result.Name,
			&result.PhoneNumber,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
